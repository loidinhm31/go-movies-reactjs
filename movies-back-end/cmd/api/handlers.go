package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/golang-jwt/jwt/v4"
	"io"
	"log"
	"movies/backend/internal/graph"
	"movies/backend/internal/models"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

func (app *Application) Home(w http.ResponseWriter, r *http.Request) {
	var payload = struct {
		Status  string `json:"status"`
		Message string `json:"message"`
		Version string `json:"version"`
	}{
		Status:  "active",
		Message: "Go Movies up and running",
		Version: "1.0.0",
	}
	_ = app.writeJSON(w, http.StatusOK, payload)
}

func (app *Application) authenticate(w http.ResponseWriter, r *http.Request) {
	// Read json payload
	var requestPayload struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	// Validate user against database
	user, err := app.DB.GetUserByEmail(requestPayload.Email)
	if err != nil {
		app.errorJSON(w, errors.New("invalid credentials"), http.StatusBadRequest)
		return
	}

	// Check password
	valid, err := user.PasswordMatches(requestPayload.Password)
	if err != nil || !valid {
		app.errorJSON(w, errors.New("invalid credentials"), http.StatusBadRequest)
		return
	}

	// Create a jwt user
	u := jwtUser{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
	}

	// Generate tokens
	tokens, err := app.auth.GenerateTokenPair(&u)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	refreshCookie := app.auth.GetRefreshCookie(tokens.RefreshToken)
	http.SetCookie(w, refreshCookie)

	app.writeJSON(w, http.StatusAccepted, tokens)
}

func (app *Application) refreshToken(w http.ResponseWriter, r *http.Request) {
	for _, cookie := range r.Cookies() {
		if cookie.Name == app.auth.CookieName {
			claims := &Claims{}
			refreshToken := cookie.Value

			// Parse the token to get the claims
			_, err := jwt.ParseWithClaims(refreshToken, claims, func(token *jwt.Token) (interface{}, error) {
				return []byte(app.JWTSecret), nil
			})
			if err != nil {
				app.errorJSON(w, errors.New("unauthorized"), http.StatusUnauthorized)
				return
			}

			// Get the user id from the token claims
			userID, err := strconv.Atoi(claims.Subject)
			if err != nil {
				app.errorJSON(w, errors.New("unknown user"), http.StatusUnauthorized)
				return
			}

			user, err := app.DB.GetUserByID(userID)
			if err != nil {
				app.errorJSON(w, errors.New("unknown user"), http.StatusUnauthorized)
				return
			}

			u := jwtUser{
				ID:        user.ID,
				FirstName: user.FirstName,
				LastName:  user.LastName,
			}

			tokenPairs, err := app.auth.GenerateTokenPair(&u)
			if err != nil {
				app.errorJSON(w, errors.New("error generating tokens"), http.StatusUnauthorized)
				return
			}

			http.SetCookie(w, app.auth.GetRefreshCookie(tokenPairs.RefreshToken))

			app.writeJSON(w, http.StatusOK, tokenPairs)

		}
	}
}

func (app *Application) logout(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, app.auth.GetExpiredRefreshCookie())
	w.WriteHeader(http.StatusAccepted)
}

func (app *Application) AllMovies(w http.ResponseWriter, r *http.Request) {
	movies, err := app.DB.AllMovies()
	if err != nil {
		app.errorJSON(w, err)
		return
	}
	_ = app.writeJSON(w, http.StatusOK, movies)
}

func (app *Application) GetMovie(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	movieID, err := strconv.Atoi(id)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	movie, err := app.DB.OneMovie(movieID)
	if err != nil {
		app.errorJSON(w, err)
		return
	}
	_ = app.writeJSON(w, http.StatusOK, movie)
}

func (app *Application) MovieForEdit(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	movieID, err := strconv.Atoi(id)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	movie, genres, err := app.DB.OneMovieForEdit(movieID)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	var payload = struct {
		Movie  *models.Movie   `json:"movie"`
		Genres []*models.Genre `json:"genres"`
	}{
		movie,
		genres,
	}
	_ = app.writeJSON(w, http.StatusOK, payload)
}

func (app *Application) AllGenres(w http.ResponseWriter, r *http.Request) {
	genres, err := app.DB.AllGenres()
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	_ = app.writeJSON(w, http.StatusOK, genres)
}

func (app *Application) InsertMovie(w http.ResponseWriter, r *http.Request) {
	var movie models.Movie

	err := app.readJSON(w, r, &movie)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	// Try to get an image
	movie = app.getPoster(movie)

	movie.CreatedAt = time.Now()
	movie.UpdatedAt = time.Now()

	newID, err := app.DB.InsertMovie(movie)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	// Handle genres
	err = app.DB.UpdateMovieGenres(newID, movie.GenresArray)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	resp := JSONResponse{
		Error:   false,
		Message: "movie updated",
	}

	app.writeJSON(w, http.StatusAccepted, resp)
}

func (app *Application) getPoster(movie models.Movie) models.Movie {
	type TheMovieDB struct {
		Page    int `json:"page"`
		Results []struct {
			PosterPath string `json:"poster_path"`
		} `json:"results"`
		TotalPages int `json:"total_pages"`
	}

	client := &http.Client{}
	theUrl := fmt.Sprintf("https://api.themoviedb.org/3/search/movie?api_key=%s", app.APIKey)

	req, err := http.NewRequest("GET", theUrl+"&query="+url.QueryEscape(movie.Title), nil)
	if err != nil {
		log.Println(err)
		return movie
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return movie
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return movie
	}

	var responseObject TheMovieDB

	json.Unmarshal(bodyBytes, &responseObject)

	if len(responseObject.Results) > 0 {
		movie.Image = responseObject.Results[0].PosterPath
	}

	return movie
}

func (app *Application) UpdateMovie(w http.ResponseWriter, r *http.Request) {
	var payload models.Movie

	err := app.readJSON(w, r, &payload)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	movie, err := app.DB.OneMovie(payload.ID)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	movie.Title = payload.Title
	movie.ReleaseDate = payload.ReleaseDate
	movie.Description = payload.Description
	movie.MPAARating = payload.MPAARating
	movie.RunTime = payload.RunTime
	movie.UpdatedAt = time.Now()

	err = app.DB.UpdateMovie(*movie)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	err = app.DB.UpdateMovieGenres(movie.ID, payload.GenresArray)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	resp := JSONResponse{
		Error:   false,
		Message: "movie updated",
	}
	app.writeJSON(w, http.StatusAccepted, resp)
}

func (app *Application) DeleteMovie(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	err = app.DB.DeleteMovie(id)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	resp := JSONResponse{
		Error:   false,
		Message: "movie deleted",
	}
	app.writeJSON(w, http.StatusAccepted, resp)
}

func (app *Application) AllMoviesByGenre(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	movies, err := app.DB.AllMovies(id)
	if err != nil {
		app.errorJSON(w, err)
		return
	}
	app.writeJSON(w, http.StatusOK, movies)
}

func (app *Application) moviesGraphQL(w http.ResponseWriter, r *http.Request) {
	// Need to populate our Graph type with the movies
	movies, _ := app.DB.AllMovies()

	// Get the query from the request
	q, _ := io.ReadAll(r.Body)
	query := string(q)

	// Create a new variable of type *graph.Graph
	g := graph.New(movies)

	// Set the query string on the variable
	g.QueryString = query

	// Perform the query
	resp, err := g.Query()
	if err != nil {
		app.errorJSON(w, err)
		return
	}
	app.writeJSON(w, http.StatusOK, resp)
	//j, _ := json.MarshalIndent(resp, "", "\t")
	//w.Header().Set("Content-Type", "application/json")
	//w.WriteHeader(http.StatusOK)
	//w.Write(j)
}

package http

import (
	"github.com/gin-gonic/gin"
	"movies-service/internal/genres"
	"net/http"
)

type genreHandler struct {
	genreService genres.Service
}

func NewGenreHandler(genreService genres.Service) genres.GenreHandler {
	return &genreHandler{
		genreService: genreService,
	}
}

func (mh *genreHandler) FetchGenres() gin.HandlerFunc {
	return func(c *gin.Context) {
		genres, err := mh.genreService.GetAllGenres(c)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			c.Abort()
			return
		}
		c.JSON(http.StatusOK, genres)
	}
}

//
//func (app *Application) getPoster(movie models.Movie) models.Movie {
//	type TheMovieDB struct {
//		PageNumber    int `json:"page"`
//		Results []struct {
//			PosterPath string `json:"poster_path"`
//		} `json:"results"`
//		TotalPages int `json:"total_pages"`
//	}
//
//	client := &http.Client{}
//	theUrl := fmt.Sprintf("https://api.themoviedb.org/3/search/movie?api_key=%s", app.APIKey)
//
//	req, err := http.NewRequest("GET", theUrl+"&query="+url.QueryEscape(movie.Title), nil)
//	if err != nil {
//		log.Println(err)
//		return movie
//	}
//
//	req.Header.Add("Accept", "application/json")
//	req.Header.Add("Content-Type", "application/json")
//
//	resp, err := client.Do(req)
//	if err != nil {
//		log.Println(err)
//		return movie
//	}
//	defer resp.Body.Close()
//
//	bodyBytes, err := io.ReadAll(resp.Body)
//	if err != nil {
//		log.Println(err)
//		return movie
//	}
//
//	var responseObject TheMovieDB
//
//	json.Unmarshal(bodyBytes, &responseObject)
//
//	if len(responseObject.Results) > 0 {
//		movie.Image = responseObject.Results[0].PosterPath
//	}
//
//	return movie
//}

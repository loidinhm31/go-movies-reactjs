package service

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"movies-service/config"
	"movies-service/internal/control"
	"movies-service/internal/dto"
	"movies-service/internal/errors"
	"movies-service/internal/middlewares"
	"movies-service/internal/reference"
	"net/http"
	"net/url"
	"time"
)

type referenceService struct {
	cfg  *config.Config
	ctrl control.Service
}

func NewReferenceService(cfg *config.Config, ctrl control.Service) reference.Service {
	return &referenceService{
		cfg,
		ctrl,
	}
}

func (is *referenceService) GetMoviesByType(ctx context.Context, movie *dto.MovieDto) ([]*dto.MovieDto, error) {
	log.Println("checking privilege...")
	username := fmt.Sprintf("%s", ctx.Value(middlewares.CtxUserKey))
	if !is.ctrl.CheckPrivilege(username) {
		return nil, errors.ErrUnAuthorized
	}

	client := &http.Client{}

	var theUrl string
	if movie.TypeCode == "TV" {
		theUrl = fmt.Sprintf("%s/search/tv?api_key=%s", is.cfg.Tmdb.Url, is.cfg.Tmdb.ApiKey)
	} else {
		theUrl = fmt.Sprintf("%s/search/movie?api_key=%s", is.cfg.Tmdb.Url, is.cfg.Tmdb.ApiKey)
	}

	req, err := http.NewRequest("GET", theUrl+"&query="+url.QueryEscape(movie.Title), nil)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")

	log.Println("executing query...")
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		err = Body.Close()
		if err != nil {
			log.Println(err)
		}
	}(resp.Body)

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	var responseObject *dto.TheMovieDBPage

	err = json.Unmarshal(bodyBytes, &responseObject)
	if err != nil {
		return nil, err
	}
	log.Println("got results from TMDB")

	var movieDtos []*dto.MovieDto
	var releaseDate time.Time

	for _, m := range responseObject.Results {
		releaseDate, err = time.Parse("2006-01-02", m.ReleaseDate)
		var title string
		if movie.TypeCode == "TV" {
			title = m.Name
			releaseDate, err = time.Parse("2006-01-02", m.FirstAirDate)
		} else if movie.TypeCode == "MOVIE" {
			title = m.Title
			releaseDate, err = time.Parse("2006-01-02", m.ReleaseDate)
		}

		movieDtos = append(movieDtos, &dto.MovieDto{
			ID:          int(m.ID),
			Title:       title,
			ReleaseDate: releaseDate,
			Description: m.Overview,
			ImagePath:   m.PosterPath,
			VoteCount:   m.VoteCount,
			VoteAverage: m.VoteAverage,
		})
	}
	return movieDtos, nil
}

func (is *referenceService) GetMovieById(ctx context.Context, movieId int64, movieType string) (*dto.MovieDto, error) {
	log.Println("checking privilege...")
	username := fmt.Sprintf("%s", ctx.Value(middlewares.CtxUserKey))
	if !is.ctrl.CheckPrivilege(username) {
		return nil, errors.ErrUnAuthorized
	}

	client := &http.Client{}

	var theUrl string
	if movieType == "MOVIE" {
		theUrl = fmt.Sprintf("%s/movie/%d?api_key=%s", is.cfg.Tmdb.Url, movieId, is.cfg.Tmdb.ApiKey)
	} else if movieType == "TV" {
		theUrl = fmt.Sprintf("%s/tv/%d?api_key=%s", is.cfg.Tmdb.Url, movieId, is.cfg.Tmdb.ApiKey)
	}

	req, err := http.NewRequest("GET", theUrl, nil)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")

	log.Println("executing query...")
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		err = Body.Close()
		if err != nil {
			log.Println(err)
		}
	}(resp.Body)

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	var responseObject *dto.TheMovieDB

	err = json.Unmarshal(bodyBytes, &responseObject)
	if err != nil {
		return nil, err
	}
	log.Println("got the result from TMDB")

	var genres []*dto.GenreDto
	for _, g := range responseObject.Genres {
		genres = append(genres, &dto.GenreDto{
			ID:   g.ID,
			Name: g.Name,
		})
	}

	var title string
	var releaseDate time.Time
	var runtime int
	if movieType == "TV" {
		title = responseObject.Name
		releaseDate, err = time.Parse("2006-01-02", responseObject.FirstAirDate)
		if len(responseObject.EpisodeRuntime) > 0 {
			runtime = responseObject.EpisodeRuntime[0]
		}
	} else if movieType == "MOVIE" {
		title = responseObject.Title
		releaseDate, err = time.Parse("2006-01-02", responseObject.ReleaseDate)
		runtime = responseObject.Runtime
	}

	return &dto.MovieDto{
		ID:          int(responseObject.ID),
		TypeCode:    movieType,
		Title:       title,
		ReleaseDate: releaseDate,
		Description: responseObject.Overview,
		ImagePath:   responseObject.PosterPath,
		Runtime:     runtime,
		Genres:      genres,
	}, nil
}

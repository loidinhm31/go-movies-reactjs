package service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"movies-service/config"
	"movies-service/internal/control"
	"movies-service/internal/dto"
	"movies-service/internal/integration"
	"movies-service/internal/middlewares"
	"movies-service/internal/movies"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

type integrationService struct {
	cfg             *config.Config
	ctrl            control.Service
	movieRepository movies.MovieRepository
}

func NewIntegrationService(cfg *config.Config, ctrl control.Service, movieRepository movies.MovieRepository) integration.Service {
	return &integrationService{
		cfg,
		ctrl,
		movieRepository,
	}
}

func (is *integrationService) credentials() *cloudinary.Cloudinary {
	// Add your Cloudinary credentials, set configuration parameter
	cld, _ := cloudinary.NewFromParams(
		is.cfg.Cloudinary.CloudName,
		is.cfg.Cloudinary.ApiKey,
		is.cfg.Cloudinary.ApiSecret,
	)
	cld.Config.URL.Secure = true
	return cld
}

func (is *integrationService) UploadVideo(ctx context.Context, file multipart.File) (string, error) {
	log.Println("checking role...")
	username := fmt.Sprintf("%s", ctx.Value(middlewares.CtxUserKey))
	if !is.ctrl.CheckUser(username) {
		return "", errors.New("user not found")
	}

	cld := is.credentials()
	log.Println("uploading file...")
	fileKey := is.cfg.Cloudinary.FolderPath + strconv.FormatInt(time.Now().Unix(), 10)
	res, err := cld.Upload.Upload(ctx, file, uploader.UploadParams{
		PublicID:   fileKey,
		EagerAsync: api.Bool(true),
	})
	if err != nil {
		return "", err
	}
	log.Printf("complete upload file %s\n", fileKey)
	return fileKey + "." + res.Format, nil
}

func (is *integrationService) DeleteVideo(ctx context.Context, fileId string) (string, error) {
	log.Println("checking role...")
	username := fmt.Sprintf("%s", ctx.Value(middlewares.CtxUserKey))
	if !is.ctrl.CheckUser(username) {
		return "", errors.New("user not found")
	}

	cld := is.credentials()

	log.Println("deleting file...")
	filePath := is.cfg.Cloudinary.FolderPath + fileId
	res, err := cld.Upload.Destroy(ctx, uploader.DestroyParams{
		PublicID:     filePath,
		ResourceType: "video",
	})
	if err != nil {
		return "", err
	}
	log.Printf("deleted file %s\n", filePath)
	return res.Result, nil
}

func (is *integrationService) GetMoviesByType(ctx context.Context, movie *dto.MovieDto) ([]*dto.MovieDto, error) {
	log.Println("checking privilege...")
	username := fmt.Sprintf("%s", ctx.Value(middlewares.CtxUserKey))
	if !is.ctrl.CheckPrivilege(username) {
		return nil, errors.New("unauthorized")
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
		releaseDate, err = time.Parse("2006-01-02", m.ReleaseDae)
		movieDtos = append(movieDtos, &dto.MovieDto{
			ID:          int(m.ID),
			Title:       m.Title,
			ReleaseDate: releaseDate,
			Description: m.Overview,
			ImagePath:   m.PosterPath,
			VoteCount:   m.VoteCount,
			VoteAverage: m.VoteAverage,
		})
	}
	return movieDtos, nil
}

func (is *integrationService) GetMovieById(ctx context.Context, movieId int64, movieType string) (*dto.MovieDto, error) {
	log.Println("checking privilege...")
	username := fmt.Sprintf("%s", ctx.Value(middlewares.CtxUserKey))
	if !is.ctrl.CheckPrivilege(username) {
		return nil, errors.New("unauthorized")
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

	releaseDate, err := time.Parse("2006-01-02", responseObject.ReleaseDae)
	return &dto.MovieDto{
		ID:          int(responseObject.ID),
		TypeCode:    movieType,
		Title:       responseObject.Title,
		ReleaseDate: releaseDate,
		Description: responseObject.Overview,
		ImagePath:   responseObject.PosterPath,
		Runtime:     responseObject.Runtime,
		Genres:      genres,
	}, nil
}

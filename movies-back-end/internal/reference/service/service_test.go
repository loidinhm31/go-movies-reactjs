package service

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"movies-service/config"
	"movies-service/internal/dto"
	"movies-service/internal/errors"
	"movies-service/internal/test/helper"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetMoviesByType(t *testing.T) {
	t.Run("Unauthorized user", func(t *testing.T) {
		ctrl := new(helper.MockManagementCtrl)
		ctrl.On("CheckPrivilege", mock.Anything).Return(false)

		service := NewReferenceService(nil, ctrl)

		movie := &dto.MovieDto{
			Title:    "Test Movie",
			TypeCode: "MOVIE",
		}

		movies, err := service.GetMoviesByType(context.Background(), movie)

		assert.Error(t, err)
		assert.Equal(t, errors.ErrUnAuthorized, err)
		assert.Nil(t, movies)
	})

	t.Run("API request error", func(t *testing.T) {
		mockCtrl := new(helper.MockManagementCtrl)
		mockCtrl.On("CheckPrivilege", mock.Anything).Return(true)

		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusInternalServerError)
		}))
		defer server.Close()

		cfg := &config.Config{
			Tmdb: config.TmdbConfig{
				Url:    server.URL,
				ApiKey: "your-api-key",
			},
		}

		service := NewReferenceService(cfg, mockCtrl)

		movie := &dto.MovieDto{
			Title:    "Test Movie",
			TypeCode: "MOVIE",
		}

		movies, err := service.GetMoviesByType(context.Background(), movie)

		assert.Error(t, err)
		assert.Nil(t, movies)
	})

	t.Run("Parsing error for API response", func(t *testing.T) {
		mockCtrl := new(helper.MockManagementCtrl)
		mockCtrl.On("CheckPrivilege", mock.Anything).Return(true)

		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			response := "[]"
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(response))
		}))
		defer server.Close()

		cfg := &config.Config{
			Tmdb: config.TmdbConfig{
				Url:    server.URL,
				ApiKey: "your-api-key",
			},
		}

		service := NewReferenceService(cfg, mockCtrl)

		movie := &dto.MovieDto{
			Title:    "Test Movie",
			TypeCode: "MOVIE",
		}

		movies, err := service.GetMoviesByType(context.Background(), movie)

		assert.Error(t, err)
		assert.Nil(t, movies)
	})

	t.Run("Valid response from API (type MOVIE)", func(t *testing.T) {
		mockCtrl := new(helper.MockManagementCtrl)
		mockCtrl.On("CheckPrivilege", mock.Anything).Return(true)

		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			response := `{"results": [{"id": 1, "title": "Movie 1"}, {"id": 2, "title": "Movie 2"}]}`
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(response))
		}))
		defer server.Close()

		cfg := &config.Config{
			Tmdb: config.TmdbConfig{
				Url:    server.URL,
				ApiKey: "your-api-key",
			},
		}

		service := NewReferenceService(cfg, mockCtrl)

		movies, err := service.GetMoviesByType(context.Background(), &dto.MovieDto{
			Title: "Ref", TypeCode: "MOVIE",
		})

		assert.NoError(t, err)
		assert.NotNil(t, movies)
		assert.Len(t, movies, 2)
		assert.Equal(t, "Movie 1", movies[0].Title)
		assert.Equal(t, "Movie 2", movies[1].Title)
	})

	t.Run("Valid response from API (type TV)", func(t *testing.T) {
		mockCtrl := new(helper.MockManagementCtrl)
		mockCtrl.On("CheckPrivilege", mock.Anything).Return(true)

		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			response := `{"results": [{"id": 1, "name": "TV 1"}, {"id": 2, "name": "TV 2"}]}`
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(response))
		}))
		defer server.Close()

		cfg := &config.Config{
			Tmdb: config.TmdbConfig{
				Url:    server.URL,
				ApiKey: "your-api-key",
			},
		}

		service := NewReferenceService(cfg, mockCtrl)

		movies, err := service.GetMoviesByType(context.Background(), &dto.MovieDto{
			Title: "Ref", TypeCode: "TV",
		})

		assert.NoError(t, err)
		assert.NotNil(t, movies)
		assert.Len(t, movies, 2)
		assert.Equal(t, "TV 1", movies[0].Title)
		assert.Equal(t, "TV 2", movies[1].Title)
	})
}

func TestGetMovieById(t *testing.T) {
	t.Run("Unauthorized user", func(t *testing.T) {
		mockCtrl := new(helper.MockManagementCtrl)
		mockCtrl.On("CheckPrivilege", mock.Anything).Return(false)

		service := NewReferenceService(nil, mockCtrl)

		movieId := int64(1)
		movieType := "MOVIE"

		movie, err := service.GetMovieById(context.Background(), movieId, movieType)

		assert.Error(t, err)
		assert.Equal(t, errors.ErrUnAuthorized, err)
		assert.Nil(t, movie)

	})

	t.Run("API request error", func(t *testing.T) {
		mockCtrl := new(helper.MockManagementCtrl)
		mockCtrl.On("CheckPrivilege", mock.Anything).Return(true)

		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusInternalServerError)
		}))
		defer server.Close()

		cfg := &config.Config{
			Tmdb: config.TmdbConfig{
				Url:    server.URL,
				ApiKey: "your-api-key",
			},
		}

		service := NewReferenceService(cfg, mockCtrl)

		movieId := int64(1)
		movieType := "MOVIE"

		movie, err := service.GetMovieById(context.Background(), movieId, movieType)

		assert.Error(t, err)
		assert.Nil(t, movie)
	})

	t.Run("Parsing error for API response", func(t *testing.T) {
		mockCtrl := new(helper.MockManagementCtrl)
		mockCtrl.On("CheckPrivilege", mock.Anything).Return(true)

		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			response := `{["}`
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(response))
		}))
		defer server.Close()

		cfg := &config.Config{
			Tmdb: config.TmdbConfig{
				Url:    server.URL,
				ApiKey: "your-api-key",
			},
		}

		service := NewReferenceService(cfg, mockCtrl)

		movieId := int64(1)
		movieType := "MOVIE"

		movie, err := service.GetMovieById(context.Background(), movieId, movieType)

		assert.Error(t, err)
		assert.Nil(t, movie)
	})

	t.Run("Valid response from API (type MOVIE)", func(t *testing.T) {
		mockCtrl := new(helper.MockManagementCtrl)
		mockCtrl.On("CheckPrivilege", mock.Anything).Return(true)

		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			response := `{"id": 1, "title": "Test Movie", "release_date": "1974-05-21", "overview": "Test Overview", "poster_path": "/path/to/image.jpg", "runtime": 120, "genres": [{"id": 1, "name": "Action"}, {"id": 2, "name": "Adventure"}]}`
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(response))
		}))
		defer server.Close()

		cfg := &config.Config{
			Tmdb: config.TmdbConfig{
				Url:    server.URL,
				ApiKey: "your-api-key",
			},
		}

		service := NewReferenceService(cfg, mockCtrl)

		movieId := int64(1)
		movieType := "MOVIE"

		movie, err := service.GetMovieById(context.Background(), movieId, movieType)

		assert.NoError(t, err)
		assert.NotNil(t, movie)
		assert.Equal(t, int(movieId), movie.ID)
		assert.Equal(t, movieType, movie.TypeCode)
		assert.Equal(t, "Test Movie", movie.Title)
		assert.Equal(t, "1974-05-21", movie.ReleaseDate.Format("2006-01-02"))
		assert.Equal(t, "Test Overview", movie.Description)
		assert.Equal(t, "/path/to/image.jpg", movie.ImageUrl)
		assert.Equal(t, 120, movie.Runtime)
		assert.Len(t, movie.Genres, 2)
		assert.Equal(t, 1, movie.Genres[0].ID)
		assert.Equal(t, "Action", movie.Genres[0].Name)
		assert.Equal(t, 2, movie.Genres[1].ID)
		assert.Equal(t, "Adventure", movie.Genres[1].Name)
	})

	t.Run("Valid response from API (type TV)", func(t *testing.T) {
		mockCtrl := new(helper.MockManagementCtrl)
		mockCtrl.On("CheckPrivilege", mock.Anything).Return(true)

		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			response := `{"id": 1, "name": "Test Movie", "first_air_date": "1974-05-21", "overview": "Test Overview", "poster_path": "/path/to/image.jpg", "episode_run_time": [120], "genres": [{"id": 1, "name": "Action"}, {"id": 2, "name": "Adventure"}]}`
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(response))
		}))
		defer server.Close()

		cfg := &config.Config{
			Tmdb: config.TmdbConfig{
				Url:    server.URL,
				ApiKey: "your-api-key",
			},
		}

		service := NewReferenceService(cfg, mockCtrl)

		movieId := int64(1)
		movieType := "TV"

		movie, err := service.GetMovieById(context.Background(), movieId, movieType)

		assert.NoError(t, err)
		assert.NotNil(t, movie)
		assert.Equal(t, int(movieId), movie.ID)
		assert.Equal(t, movieType, movie.TypeCode)
		assert.Equal(t, "Test Movie", movie.Title)
		assert.Equal(t, "1974-05-21", movie.ReleaseDate.Format("2006-01-02"))
		assert.Equal(t, "Test Overview", movie.Description)
		assert.Equal(t, "/path/to/image.jpg", movie.ImageUrl)
		assert.Equal(t, 120, movie.Runtime)
		assert.Len(t, movie.Genres, 2)
		assert.Equal(t, 1, movie.Genres[0].ID)
		assert.Equal(t, "Action", movie.Genres[0].Name)
		assert.Equal(t, 2, movie.Genres[1].ID)
		assert.Equal(t, "Adventure", movie.Genres[1].Name)
	})

}

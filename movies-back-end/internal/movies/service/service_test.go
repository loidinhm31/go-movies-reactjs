package service

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"movies-service/internal/dto"
	"movies-service/internal/errors"
	"movies-service/internal/models"
	"movies-service/internal/movies"
	"movies-service/pkg/pagination"
	"movies-service/pkg/test_helper"
	"testing"
	"time"
)

func initMock() (*test_helper.MockManagementCtrl, *test_helper.MockMovieRepository, movies.Service) {
	// Create a mock movie repository
	mockRepo := new(test_helper.MockMovieRepository)

	// Create a mock management controller
	mockCtrl := new(test_helper.MockManagementCtrl)

	// Create a genre service instance with the mock repository and controller
	movieService := NewMovieService(mockCtrl, mockRepo)

	return mockCtrl, mockRepo, movieService
}

func TestGetAllMoviesByType(t *testing.T) {

	t.Run("Valid movie type (MOVIE)", func(t *testing.T) {
		_, mockRepo, movieService := initMock()

		// Set up mock expectations and return values
		mockRepo.On("FindAllMoviesByType", context.Background(), "MOVIE", &pagination.PageRequest{}, &pagination.Page[*models.Movie]{}).
			Return(&pagination.Page[*models.Movie]{
				PageSize:      1,
				PageNumber:    0,
				Sort:          pagination.Sort{},
				TotalElements: 2,
				TotalPages:    1,
				Content: []*models.Movie{
					{Title: "Movie1", TypeCode: "MOVIE", Genres: []*models.Genre{{Name: "Action", TypeCode: "MOVIE"}}},
					{Title: "Movie2", TypeCode: "MOVIE", Genres: []*models.Genre{{Name: "Action", TypeCode: "MOVIE"}}},
				},
			}, nil)

		pageMovie, err := movieService.GetAllMoviesByType(context.Background(), "MOVIE", &pagination.PageRequest{})
		assert.NoError(t, err)
		assert.Equal(t, 2, len(pageMovie.Content))
		assert.Equal(t, "Movie1", pageMovie.Content[0].Title)
		assert.Equal(t, "Movie2", pageMovie.Content[1].Title)
	})

	t.Run("Empty movie type", func(t *testing.T) {
		_, mockRepo, movieService := initMock()

		// Set up mock expectations and return values
		mockRepo.On("FindAllMovies", context.Background(), &pagination.PageRequest{}, &pagination.Page[*models.Movie]{}).
			Return(&pagination.Page[*models.Movie]{
				PageSize:      1,
				PageNumber:    0,
				Sort:          pagination.Sort{},
				TotalElements: 2,
				TotalPages:    1,
				Content: []*models.Movie{
					{Title: "Movie1", TypeCode: "MOVIE", Genres: []*models.Genre{{Name: "Action", TypeCode: "MOVIE"}}},
					{Title: "Movie2", TypeCode: "TV", Genres: []*models.Genre{{Name: "Action", TypeCode: "TV"}}},
				},
			}, nil)

		pageMovie, err := movieService.GetAllMoviesByType(context.Background(), "", &pagination.PageRequest{})

		assert.NoError(t, err)
		assert.Equal(t, 2, len(pageMovie.Content))
		assert.Equal(t, "Movie1", pageMovie.Content[0].Title)
		assert.Equal(t, "MOVIE", pageMovie.Content[0].TypeCode)
		assert.Equal(t, "Movie2", pageMovie.Content[1].Title)
		assert.Equal(t, "TV", pageMovie.Content[1].TypeCode)
	})
}

func TestGetMovieById(t *testing.T) {
	t.Run("Valid", func(t *testing.T) {
		_, mockRepo, movieService := initMock()

		// Set up mock expectations and return values
		mockRepo.On("FindMovieById", context.Background(), 1).
			Return(&models.Movie{
				ID:       1,
				Title:    "Movie1",
				TypeCode: "TV",
			}, nil)

		result, err := movieService.GetMovieById(context.Background(), 1)
		assert.NoError(t, err)
		assert.Equal(t, 1, result.ID)
		assert.Equal(t, "Movie1", result.Title)
		assert.Equal(t, "TV", result.TypeCode)
	})
}

func TestGetMoviesByGenre(t *testing.T) {
	t.Run("Valid data", func(t *testing.T) {
		_, mockRepo, movieService := initMock()

		// Set up mock expectations and return values
		genreId := 5

		mockRepo.On("FindMoviesByGenre", context.Background(), &pagination.PageRequest{}, &pagination.Page[*models.Movie]{}, genreId).
			Return(&pagination.Page[*models.Movie]{
				PageSize:      1,
				PageNumber:    0,
				Sort:          pagination.Sort{},
				TotalElements: 2,
				TotalPages:    1,
				Content: []*models.Movie{
					{Title: "Movie1", TypeCode: "MOVIE", Genres: []*models.Genre{{ID: genreId, Name: "Action", TypeCode: "MOVIE"}}},
					{Title: "Movie2", TypeCode: "MOVIE", Genres: []*models.Genre{{ID: genreId, Name: "Action", TypeCode: "MOVIE"}}},
				},
			}, nil)

		pageMovie, err := movieService.GetMoviesByGenre(context.Background(), &pagination.PageRequest{}, genreId)

		assert.NoError(t, err)
		assert.Equal(t, 2, len(pageMovie.Content))
		assert.Equal(t, "Movie1", pageMovie.Content[0].Title)
		assert.Equal(t, "MOVIE", pageMovie.Content[0].TypeCode)
		assert.Equal(t, genreId, pageMovie.Content[0].Genres[0].ID)
		assert.Equal(t, "Movie2", pageMovie.Content[1].Title)
		assert.Equal(t, "MOVIE", pageMovie.Content[1].TypeCode)
		assert.Equal(t, genreId, pageMovie.Content[1].Genres[0].ID)
	})
}

func TestAddMovie(t *testing.T) {
	t.Run("Invalid Input", func(t *testing.T) {
		_, _, movieService := initMock()

		err := movieService.AddMovie(context.Background(), &dto.MovieDto{
			ID: 1,
		})
		assert.Error(t, err)
		assert.Equal(t, errors.ErrInvalidInput.Error(), err.Error())
	})

	t.Run("Unauthorized", func(t *testing.T) {
		mockCtrl, _, movieService := initMock()

		// Set up mock expectations and return values
		mockCtrl.On("CheckPrivilege", mock.Anything).Return(false)

		err := movieService.AddMovie(context.Background(), &dto.MovieDto{
			Title:       "Movie1",
			TypeCode:    "TV",
			Runtime:     10,
			Description: "Hello",
			ReleaseDate: time.Now(),
			MpaaRating:  "R",
			Genres: []*dto.GenreDto{
				{Name: "Genre1", TypeCode: "TV"},
			},
		})
		assert.Error(t, err)
		assert.Equal(t, errors.ErrUnAuthorized.Error(), err.Error())
	})

	t.Run("Mismatch Type code", func(t *testing.T) {
		mockCtrl, _, movieService := initMock()

		// Set up mock expectations and return values
		mockCtrl.On("CheckPrivilege", mock.Anything).Return(true)

		err := movieService.AddMovie(context.Background(), &dto.MovieDto{
			Title:       "Movie1",
			TypeCode:    "TV",
			Runtime:     10,
			Description: "Hello",
			ReleaseDate: time.Now(),
			MpaaRating:  "R",
			Genres: []*dto.GenreDto{
				{Name: "Genre2", TypeCode: "TV", Checked: true},
				{Name: "Genre1", TypeCode: "MOVIE", Checked: true},
			},
		})
		assert.Error(t, err)
		assert.Equal(t, errors.ErrInvalidInput.Error(), err.Error())
	})

	t.Run("Genres empty with check", func(t *testing.T) {
		mockCtrl, _, movieService := initMock()

		// Set up mock expectations and return values
		mockCtrl.On("CheckPrivilege", mock.Anything).Return(true)

		err := movieService.AddMovie(context.Background(), &dto.MovieDto{
			Title:       "Movie1",
			TypeCode:    "TV",
			Runtime:     10,
			Description: "Hello",
			ReleaseDate: time.Now(),
			MpaaRating:  "R",
			Genres: []*dto.GenreDto{
				{Name: "Genre2", TypeCode: "TV", Checked: false},
				{Name: "Genre1", TypeCode: "MOVIE", Checked: false},
			},
		})
		assert.Error(t, err)
		assert.Equal(t, errors.ErrInvalidInputDetail("genres cannot empty").Error(), err.Error())
	})

	t.Run("Valid data", func(t *testing.T) {
		mockCtrl, mockRepo, movieService := initMock()

		// Set up mock expectations and return values
		mockCtrl.On("CheckPrivilege", mock.Anything).Return(true)

		mockRepo.On("InsertMovie", context.Background(), mock.Anything).
			Return(nil).Once()

		err := movieService.AddMovie(context.Background(), &dto.MovieDto{
			Title:       "Movie1",
			TypeCode:    "TV",
			Runtime:     10,
			Description: "Hello",
			ReleaseDate: time.Now(),
			MpaaRating:  "R",
			Genres: []*dto.GenreDto{
				{Name: "Genre2", TypeCode: "TV", Checked: true},
				{Name: "Genre1", TypeCode: "TV", Checked: true},
			},
		})
		assert.NoError(t, err)
	})
}

func TestUpdateMovie(t *testing.T) {
	t.Run("Invalid Input", func(t *testing.T) {
		_, _, movieService := initMock()

		err := movieService.UpdateMovie(context.Background(), &dto.MovieDto{
			ID: 0,
		})
		assert.Error(t, err)
		assert.Equal(t, errors.ErrInvalidInput.Error(), err.Error())
	})

	t.Run("Existed movie", func(t *testing.T) {
		_, mockRepo, movieService := initMock()

		// Set up mock expectations and return values
		mockRepo.On("FindMovieById", context.Background(), 1).
			Return(&models.Movie{}, errors.ErrResourceNotFound)

		err := movieService.UpdateMovie(context.Background(), &dto.MovieDto{
			ID:          1,
			Title:       "Movie1",
			TypeCode:    "TV",
			Runtime:     10,
			Description: "Hello",
			ReleaseDate: time.Now(),
			MpaaRating:  "R",
			Genres: []*dto.GenreDto{
				{Name: "Genre1", TypeCode: "TV"},
			},
		})
		assert.Error(t, err)
		assert.Equal(t, errors.ErrResourceNotFound.Error(), err.Error())
	})

	t.Run("Unauthorized", func(t *testing.T) {
		mockCtrl, mockRepo, movieService := initMock()

		// Set up mock expectations and return values
		mockRepo.On("FindMovieById", context.Background(), 1).
			Return(&models.Movie{}, nil)

		mockCtrl.On("CheckPrivilege", mock.Anything).Return(false)

		err := movieService.UpdateMovie(context.Background(), &dto.MovieDto{
			ID:          1,
			Title:       "Movie1",
			TypeCode:    "TV",
			Runtime:     10,
			Description: "Hello",
			ReleaseDate: time.Now(),
			MpaaRating:  "R",
			Genres: []*dto.GenreDto{
				{Name: "Genre1", TypeCode: "TV"},
			},
		})
		assert.Error(t, err)
		assert.Equal(t, errors.ErrUnAuthorized.Error(), err.Error())
	})

	t.Run("Mismatch Type code", func(t *testing.T) {
		mockCtrl, mockRepo, movieService := initMock()

		// Set up mock expectations and return values
		mockRepo.On("FindMovieById", context.Background(), 1).
			Return(&models.Movie{}, nil)

		mockCtrl.On("CheckPrivilege", mock.Anything).Return(true)

		mockRepo.On("UpdateMovie", context.Background(), mock.Anything).
			Return(nil)

		err := movieService.UpdateMovie(context.Background(), &dto.MovieDto{
			ID:          1,
			Title:       "Movie1",
			TypeCode:    "TV",
			Runtime:     10,
			Description: "Hello",
			ReleaseDate: time.Now(),
			MpaaRating:  "R",
			Genres: []*dto.GenreDto{
				{Name: "Genre2", TypeCode: "TV", Checked: true},
				{Name: "Genre1", TypeCode: "MOVIE", Checked: true},
			},
		})
		assert.Error(t, err)
		assert.Equal(t, errors.ErrInvalidInput.Error(), err.Error())
	})

	t.Run("Genres empty with check", func(t *testing.T) {
		mockCtrl, mockRepo, movieService := initMock()

		// Set up mock expectations and return values
		mockRepo.On("FindMovieById", context.Background(), 1).
			Return(&models.Movie{}, nil)

		mockCtrl.On("CheckPrivilege", mock.Anything).Return(true)

		mockRepo.On("UpdateMovie", context.Background(), mock.Anything).
			Return(nil)

		err := movieService.UpdateMovie(context.Background(), &dto.MovieDto{
			ID:          1,
			Title:       "Movie1",
			TypeCode:    "TV",
			Runtime:     10,
			Description: "Hello",
			ReleaseDate: time.Now(),
			MpaaRating:  "R",
			Genres: []*dto.GenreDto{
				{Name: "Genre2", TypeCode: "TV", Checked: false},
				{Name: "Genre1", TypeCode: "MOVIE", Checked: false},
			},
		})
		assert.Error(t, err)
		assert.Equal(t, errors.ErrInvalidInputDetail("genres cannot empty").Error(), err.Error())
	})

	t.Run("Valid data", func(t *testing.T) {
		mockCtrl, mockRepo, movieService := initMock()

		// Set up mock expectations and return values
		mockRepo.On("FindMovieById", context.Background(), 1).
			Return(&models.Movie{}, nil)

		mockCtrl.On("CheckPrivilege", mock.Anything).Return(true)

		mockRepo.On("UpdateMovie", context.Background(), mock.Anything).
			Return(nil)

		mockRepo.On("UpdateMovieGenres", context.Background(), mock.Anything, mock.Anything).
			Return(nil)

		err := movieService.UpdateMovie(context.Background(), &dto.MovieDto{
			ID:          1,
			Title:       "Movie1",
			TypeCode:    "MOVIE",
			Runtime:     10,
			Description: "Hello",
			ReleaseDate: time.Now(),
			MpaaRating:  "R",
			Genres: []*dto.GenreDto{
				{Name: "Genre2", TypeCode: "MOVIE", Checked: true},
				{Name: "Genre1", TypeCode: "MOVIE", Checked: false},
			},
		})
		assert.NoError(t, err)
	})
}

func TestDeleteMovieById(t *testing.T) {
	t.Run("Invalid ID", func(t *testing.T) {
		_, _, movieService := initMock()

		err := movieService.DeleteMovieById(context.Background(), 0)
		assert.Error(t, err)
		assert.Equal(t, errors.ErrInvalidInput.Error(), err.Error())
	})

	t.Run("Unauthorized", func(t *testing.T) {
		mockCtrl, _, movieService := initMock()

		// Set up mock expectations and return values
		mockCtrl.On("CheckPrivilege", mock.Anything).Return(false)

		err := movieService.DeleteMovieById(context.Background(), 1)
		assert.Error(t, err)
		assert.Equal(t, errors.ErrUnAuthorized.Error(), err.Error())
	})

	t.Run("Valid data", func(t *testing.T) {
		mockCtrl, mockRepo, movieService := initMock()

		// Set up mock expectations and return values
		mockCtrl.On("CheckPrivilege", mock.Anything).Return(true)

		mockRepo.On("DeleteMovieById", context.Background(), 1).
			Return(nil).Once()

		err := movieService.DeleteMovieById(context.Background(), 1)
		assert.NoError(t, err)
	})
}

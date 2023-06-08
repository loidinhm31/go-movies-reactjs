package service

import (
	"context"
	"database/sql"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"movies-service/internal/common/dto"
	"movies-service/internal/common/entity"
	"movies-service/internal/errors"
	"movies-service/internal/movie"
	"movies-service/internal/test/helper"
	"movies-service/pkg/pagination"
	"testing"
	"time"
)

func initMock() (*helper.MockManagementCtrl, *helper.MockUserRepository, *helper.MockMovieRepository, *helper.MockCollectionRepository, *helper.MockPaymentRepository, *helper.MockBlobService, movie.Service) {
	// Create a mock movie repository
	mockMovieRepo := new(helper.MockMovieRepository)

	// Create a mock management controller
	mockCtrl := new(helper.MockManagementCtrl)

	mockUserRepo := new(helper.MockUserRepository)

	mockBlobSvc := new(helper.MockBlobService)

	mockCollectionRepo := new(helper.MockCollectionRepository)

	mockPaymentRepo := new(helper.MockPaymentRepository)

	// Create a genre service instance with the mock repository and controller
	movieService := NewMovieService(mockCtrl, mockUserRepo, mockMovieRepo, mockCollectionRepo, mockPaymentRepo, mockBlobSvc)

	return mockCtrl, mockUserRepo, mockMovieRepo, mockCollectionRepo, mockPaymentRepo, mockBlobSvc, movieService
}

func TestMovieService_GetAllMoviesByType(t *testing.T) {

	t.Run("Valid movie type (MOVIE)", func(t *testing.T) {
		_, _, mockMovieRepo, _, _, _, movieService := initMock()

		// Set up mock expectations and return values
		mockMovieRepo.On("FindAllMoviesByType", context.Background(), "title", "MOVIE", &pagination.PageRequest{}, &pagination.Page[*entity.Movie]{}).
			Return(&pagination.Page[*entity.Movie]{
				PageSize:      1,
				PageNumber:    0,
				Sort:          pagination.Sort{},
				TotalElements: 2,
				TotalPages:    1,
				Content: []*entity.Movie{
					{Title: "Movie1", TypeCode: "MOVIE", Genres: []*entity.Genre{{Name: "Action", TypeCode: "MOVIE"}}},
					{Title: "Movie2", TypeCode: "MOVIE", Genres: []*entity.Genre{{Name: "Action", TypeCode: "MOVIE"}}},
				},
			}, nil)

		pageMovie, err := movieService.GetAllMoviesByType(context.Background(), "title", "MOVIE", &pagination.PageRequest{})
		assert.NoError(t, err)
		assert.Equal(t, 2, len(pageMovie.Content))
		assert.Equal(t, "Movie1", pageMovie.Content[0].Title)
		assert.Equal(t, "Movie2", pageMovie.Content[1].Title)
	})

	t.Run("Empty movie type", func(t *testing.T) {
		_, _, mockMovieRepo, _, _, _, movieService := initMock()

		// Set up mock expectations and return values
		mockMovieRepo.On("FindAllMovies", context.Background(), "title", &pagination.PageRequest{}, &pagination.Page[*entity.Movie]{}).
			Return(&pagination.Page[*entity.Movie]{
				PageSize:      1,
				PageNumber:    0,
				Sort:          pagination.Sort{},
				TotalElements: 2,
				TotalPages:    1,
				Content: []*entity.Movie{
					{Title: "Movie1", TypeCode: "MOVIE", Genres: []*entity.Genre{{Name: "Action", TypeCode: "MOVIE"}}},
					{Title: "Movie2", TypeCode: "TV", Genres: []*entity.Genre{{Name: "Action", TypeCode: "TV"}}},
				},
			}, nil)

		pageMovie, err := movieService.GetAllMoviesByType(context.Background(), "title", "", &pagination.PageRequest{})

		assert.NoError(t, err)
		assert.Equal(t, 2, len(pageMovie.Content))
		assert.Equal(t, "Movie1", pageMovie.Content[0].Title)
		assert.Equal(t, "MOVIE", pageMovie.Content[0].TypeCode)
		assert.Equal(t, "Movie2", pageMovie.Content[1].Title)
		assert.Equal(t, "TV", pageMovie.Content[1].TypeCode)
	})
}

func TestMovieService_GetMovieByID(t *testing.T) {
	t.Run("Invalid Client", func(t *testing.T) {
		_, mockUserRepo, _, _, _, _, movieService := initMock()

		// Set up mock expectations and return values
		mockUserRepo.On("FindUserByUsernameAndIsNew", mock.Anything, mock.Anything, false).
			Return(&entity.User{
				Role: &entity.Role{
					RoleCode: "BANNED",
				},
			}, nil)

		_, err := movieService.GetMovieByID(context.Background(), uint(1))
		assert.Error(t, err)
		assert.Equal(t, errors.ErrInvalidClient, err)
	})

	t.Run("Valid", func(t *testing.T) {
		_, mockUserRepo, mockMovieRepo, _, mockPaymentRepo, _, movieService := initMock()

		// Set up mock expectations and return values
		mockUserRepo.On("FindUserByUsernameAndIsNew", mock.Anything, mock.Anything, false).
			Return(&entity.User{
				Role: &entity.Role{
					RoleCode: "GENERAL",
				},
			}, nil)

		mockMovieRepo.On("FindMovieByID", context.Background(), mock.Anything).
			Return(&entity.Movie{
				ID:       uint(1),
				Title:    "Movie1",
				TypeCode: "TV",
				Price: sql.NullFloat64{
					Float64: 10.5,
					Valid:   true,
				},
			}, nil)

		mockPaymentRepo.On("FindPaymentByUserIDAndTypeCodeAndRefID", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
			Return(&entity.Payment{}, nil)

		result, err := movieService.GetMovieByID(context.Background(), uint(1))
		assert.NoError(t, err)
		assert.Equal(t, uint(1), result.ID)
		assert.Equal(t, "Movie1", result.Title)
		assert.Equal(t, "TV", result.TypeCode)
	})
}

func TestMovieService_GetMoviesByGenre(t *testing.T) {
	t.Run("Valid data", func(t *testing.T) {
		_, _, mockMovieRepo, _, _, _, movieService := initMock()

		// Set up mock expectations and return values
		genreId := uint(5)

		mockMovieRepo.On("FindMoviesByGenre", context.Background(), &pagination.PageRequest{}, &pagination.Page[*entity.Movie]{}, genreId).
			Return(&pagination.Page[*entity.Movie]{
				PageSize:      1,
				PageNumber:    0,
				Sort:          pagination.Sort{},
				TotalElements: 2,
				TotalPages:    1,
				Content: []*entity.Movie{
					{Title: "Movie1", TypeCode: "MOVIE", Genres: []*entity.Genre{{ID: genreId, Name: "Action", TypeCode: "MOVIE"}}},
					{Title: "Movie2", TypeCode: "MOVIE", Genres: []*entity.Genre{{ID: genreId, Name: "Action", TypeCode: "MOVIE"}}},
				},
			}, nil)

		pageMovie, err := movieService.GetMoviesByGenre(context.Background(), &pagination.PageRequest{}, uint(genreId))

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

func TestMovieService_AddMovie(t *testing.T) {
	t.Run("Invalid Input", func(t *testing.T) {
		_, _, _, _, _, _, movieService := initMock()

		err := movieService.AddMovie(context.Background(), &dto.MovieDto{
			ID: 1,
		})
		assert.Error(t, err)
		assert.Equal(t, errors.ErrInvalidInput.Error(), err.Error())
	})

	t.Run("Unauthorized", func(t *testing.T) {
		mockCtrl, _, _, _, _, _, movieService := initMock()

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
		mockCtrl, _, _, _, _, _, movieService := initMock()

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
		mockCtrl, _, _, _, _, _, movieService := initMock()

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
		mockCtrl, _, mockMovieRepo, _, _, _, movieService := initMock()

		// Set up mock expectations and return values
		mockCtrl.On("CheckPrivilege", mock.Anything).Return(true)

		mockMovieRepo.On("InsertMovie", context.Background(), mock.Anything).
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

func TestMovieService_UpdateMovie(t *testing.T) {
	t.Run("Invalid Input", func(t *testing.T) {
		_, _, _, _, _, _, movieService := initMock()

		err := movieService.UpdateMovie(context.Background(), &dto.MovieDto{
			ID: uint(0),
		})
		assert.Error(t, err)
		assert.Equal(t, errors.ErrInvalidInput.Error(), err.Error())
	})

	t.Run("Existed movie", func(t *testing.T) {
		_, _, mockMovieRepo, _, _, _, movieService := initMock()

		// Set up mock expectations and return values
		mockMovieRepo.On("FindMovieByID", context.Background(), uint(1)).
			Return(&entity.Movie{}, errors.ErrResourceNotFound)

		err := movieService.UpdateMovie(context.Background(), &dto.MovieDto{
			ID:          uint(1),
			Title:       "Movie1",
			TypeCode:    "TV",
			Runtime:     uint(10),
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
		mockCtrl, _, mockMovieRepo, _, _, _, movieService := initMock()

		// Set up mock expectations and return values
		mockMovieRepo.On("FindMovieByID", context.Background(), uint(1)).
			Return(&entity.Movie{}, nil)

		mockCtrl.On("CheckPrivilege", mock.Anything).Return(false)

		err := movieService.UpdateMovie(context.Background(), &dto.MovieDto{
			ID:          uint(1),
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
		mockCtrl, _, mockMovieRepo, _, _, _, movieService := initMock()

		// Set up mock expectations and return values
		mockMovieRepo.On("FindMovieByID", context.Background(), uint(1)).
			Return(&entity.Movie{}, nil)

		mockCtrl.On("CheckPrivilege", mock.Anything).Return(true)

		mockMovieRepo.On("UpdateMovie", context.Background(), mock.Anything).
			Return(nil)

		err := movieService.UpdateMovie(context.Background(), &dto.MovieDto{
			ID:          uint(1),
			Title:       "Movie1",
			TypeCode:    "TV",
			Runtime:     uint(10),
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
		mockCtrl, _, mockMovieRepo, _, _, _, movieService := initMock()

		// Set up mock expectations and return values
		mockMovieRepo.On("FindMovieByID", context.Background(), uint(1)).
			Return(&entity.Movie{}, nil)

		mockCtrl.On("CheckPrivilege", mock.Anything).Return(true)

		mockMovieRepo.On("UpdateMovie", context.Background(), mock.Anything).
			Return(nil)

		err := movieService.UpdateMovie(context.Background(), &dto.MovieDto{
			ID:          uint(1),
			Title:       "Movie1",
			TypeCode:    "TV",
			Runtime:     uint(10),
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
		mockCtrl, _, mockMovieRepo, _, _, _, movieService := initMock()

		// Set up mock expectations and return values
		mockMovieRepo.On("FindMovieByID", context.Background(), uint(1)).
			Return(&entity.Movie{}, nil)

		mockCtrl.On("CheckPrivilege", mock.Anything).Return(true)

		mockMovieRepo.On("UpdateMovie", context.Background(), mock.Anything).
			Return(nil)

		mockMovieRepo.On("UpdateMovieGenres", context.Background(), mock.Anything, mock.Anything).
			Return(nil)

		err := movieService.UpdateMovie(context.Background(), &dto.MovieDto{
			ID:          uint(1),
			Title:       "Movie1",
			TypeCode:    "MOVIE",
			Runtime:     uint(10),
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

func TestMovieService_RemoveMovieByID(t *testing.T) {
	t.Run("Invalid ID", func(t *testing.T) {
		_, _, _, _, _, _, movieService := initMock()

		err := movieService.RemoveMovieByID(context.Background(), 0)
		assert.Error(t, err)
		assert.Equal(t, errors.ErrInvalidInput.Error(), err.Error())
	})

	t.Run("Added to Payment ", func(t *testing.T) {
		_, _, _, _, mockPaymentRepo, _, movieService := initMock()

		movieID := uint(1)

		mockPaymentRepo.On("FindPaymentsByTypeCodeAndRefID", mock.Anything, "MOVIE", mock.Anything).
			Return([]*entity.Payment{
				{ID: uint(1)},
			}, nil)

		err := movieService.RemoveMovieByID(context.Background(), movieID)
		assert.Error(t, err)
		assert.Equal(t, errors.ErrCannotExecuteAction.Error(), err.Error())
	})

	t.Run("Added to Collection ", func(t *testing.T) {
		_, _, _, mockCollectionRepo, mockPaymentRepo, _, movieService := initMock()

		movieID := uint(1)

		mockPaymentRepo.On("FindPaymentsByTypeCodeAndRefID", mock.Anything, "MOVIE", mock.Anything).
			Return([]*entity.Payment{}, nil)

		mockCollectionRepo.On("FindCollectionsByMovieID", mock.Anything, mock.Anything).
			Return([]*entity.Collection{
				{ID: uint(1)},
			}, nil)

		err := movieService.RemoveMovieByID(context.Background(), movieID)
		assert.Error(t, err)
		assert.Equal(t, errors.ErrCannotExecuteAction.Error(), err.Error())
	})

	t.Run("Invalid Client", func(t *testing.T) {
		mockCtrl, _, _, mockCollectionRepo, mockPaymentRepo, _, movieService := initMock()

		// Set up mock expectations and return values
		mockPaymentRepo.On("FindPaymentsByTypeCodeAndRefID", mock.Anything, "MOVIE", mock.Anything).
			Return([]*entity.Payment{}, nil)

		mockCollectionRepo.On("FindCollectionsByMovieID", mock.Anything, uint(1)).
			Return([]*entity.Collection{}, nil)

		mockCtrl.On("CheckPrivilege", mock.Anything).Return(false)

		err := movieService.RemoveMovieByID(context.Background(), uint(1))
		assert.Error(t, err)
		assert.Equal(t, errors.ErrInvalidClient.Error(), err.Error())
	})

	t.Run("Valid data", func(t *testing.T) {
		mockCtrl, _, mockMovieRepo, mockCollectionRepo, mockPaymentRepo, mockBlobSvc, movieService := initMock()

		// Set up mock expectations and return values
		mockPaymentRepo.On("FindPaymentsByTypeCodeAndRefID", mock.Anything, "MOVIE", mock.Anything).
			Return([]*entity.Payment{}, nil)

		mockCollectionRepo.On("FindCollectionsByMovieID", mock.Anything, uint(1)).
			Return([]*entity.Collection{}, nil)

		mockCtrl.On("CheckPrivilege", mock.Anything).Return(true)

		mockMovieRepo.On("FindMovieByID", context.Background(), uint(1)).
			Return(&entity.Movie{
				VideoPath: sql.NullString{String: "path", Valid: true},
				ImageUrl:  sql.NullString{String: "path", Valid: true},
			}, nil)

		mockBlobSvc.On("DeleteFile", context.Background(), "path", "video").
			Return("ok", nil)

		mockBlobSvc.On("DeleteFile", context.Background(), "path", "image").
			Return("ok", nil)

		mockMovieRepo.On("DeleteMovieByID", context.Background(), uint(1)).
			Return(nil).Once()

		err := movieService.RemoveMovieByID(context.Background(), uint(1))
		assert.NoError(t, err)
	})
}

func TestMovieService_GetMovieByEpisodeID(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		mockCtrl, _, mockMovieRepo, _, _, _, movieService := initMock()

		mockCtrl.On("CheckUser", mock.Anything).Return(true, true)

		mockMovieRepo.On("FindMovieByEpisodeID", mock.Anything, mock.Anything).
			Return(&entity.Movie{
				ID:    uint(1),
				Title: "movie1",
			}, nil)

		result, err := movieService.GetMovieByEpisodeID(context.Background(), uint(10))

		assert.NoError(t, err)
		assert.Equal(t, uint(1), result.ID)
		assert.Equal(t, "movie1", result.Title)
	})
}

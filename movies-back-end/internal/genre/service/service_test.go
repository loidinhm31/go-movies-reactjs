package service

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"movies-service/internal/dto"
	"movies-service/internal/errors"
	"movies-service/internal/genre"
	"movies-service/internal/model"
	"movies-service/internal/test/helper"
	"testing"
)

func initMock() (*helper.MockManagementCtrl, *helper.MockGenreRepository, genre.Service) {
	// Create a mock genre repository
	mockRepo := new(helper.MockGenreRepository)

	// Create a mock management controller
	mockCtrl := new(helper.MockManagementCtrl)

	// Create a genre service instance with the mock repository and controller
	genreService := NewGenreService(mockCtrl, mockRepo)
	return mockCtrl, mockRepo, genreService
}

func TestGetAllGenresByTypeCode(t *testing.T) {

	t.Run("Valid movie type (MOVIE)", func(t *testing.T) {
		_, mockRepo, genreService := initMock()

		// Set up mock expectations and return values
		mockRepo.On("FindAllGenresByTypeCode", context.Background(), "MOVIE").
			Return([]*model.Genre{
				{Name: "Action1", TypeCode: "MOVIE"},
				{Name: "Action2", TypeCode: "MOVIE"},
			}, nil)

		genres, err := genreService.GetAllGenresByTypeCode(context.Background(), "MOVIE")
		assert.NoError(t, err)
		assert.Equal(t, 2, len(genres))
		assert.Equal(t, "Action1", genres[0].Name)
		assert.Equal(t, "Action2", genres[1].Name)
	})

	t.Run("Empty movie type", func(t *testing.T) {
		_, mockRepo, genreService := initMock()

		// Set up mock expectations and return values
		mockRepo.On("FindAllGenres", context.Background()).
			Return([]*model.Genre{
				{Name: "Comedy1", TypeCode: "MOVIE"},
				{Name: "Comedy2", TypeCode: "TV"},
			}, nil)

		genres, err := genreService.GetAllGenresByTypeCode(context.Background(), "")
		assert.NoError(t, err)
		assert.Equal(t, 2, len(genres))
		assert.Equal(t, "Comedy1", genres[0].Name)
		assert.Equal(t, "MOVIE", genres[0].TypeCode)
		assert.Equal(t, "Comedy2", genres[1].Name)
		assert.Equal(t, "TV", genres[1].TypeCode)
	})
}

func TestAddGenres(t *testing.T) {

	t.Run("Valid genre and privileged user", func(t *testing.T) {
		mockCtrl, mockRepo, genreService := initMock()

		// Set up mock expectations and return values
		mockCtrl.On("CheckPrivilege", mock.Anything).Return(true)

		genreDtos := []dto.GenreDto{
			{Name: "Genre3", TypeCode: "MOVIE"},
			{Name: "Genre4", TypeCode: "TV"},
		}

		mockRepo.On("FindGenreByNameAndTypeCode", context.Background(), &model.Genre{
			Name:     genreDtos[0].Name,
			TypeCode: genreDtos[0].TypeCode,
		}).
			Return(&model.Genre{}, nil)

		mockRepo.On("FindGenreByNameAndTypeCode", context.Background(), &model.Genre{
			Name:     genreDtos[1].Name,
			TypeCode: genreDtos[1].TypeCode,
		}).
			Return(&model.Genre{}, nil)

		mockRepo.On("InsertGenres", context.Background(), mock.Anything).
			Return(nil)
		err := genreService.AddGenres(context.Background(), genreDtos)
		assert.NoError(t, err)
	})

	t.Run("Genre with existing name and type code", func(t *testing.T) {
		mockCtrl, mockRepo, genreService := initMock()

		// Set up mock expectations and return values
		mockCtrl.On("CheckPrivilege", mock.Anything).Return(true)

		genreDtos := []dto.GenreDto{
			{Name: "Action1", TypeCode: "MOVIE"},
			{Name: "Genre2", TypeCode: "TV"},
		}

		mockRepo.On("FindGenreByNameAndTypeCode", context.Background(), &model.Genre{
			Name:     genreDtos[0].Name,
			TypeCode: genreDtos[0].TypeCode,
		}).
			Return(&model.Genre{}, nil)

		mockRepo.On("FindGenreByNameAndTypeCode", context.Background(), &model.Genre{
			Name:     genreDtos[1].Name,
			TypeCode: genreDtos[1].TypeCode,
		}).
			Return(&model.Genre{Name: genreDtos[1].Name, TypeCode: genreDtos[1].TypeCode}, nil)

		err := genreService.AddGenres(context.Background(), genreDtos)
		assert.Error(t, err)
		assert.Equal(t, errors.ErrCannotExecuteAction.Error(), err.Error())
	})

	t.Run("Invalid genre type code", func(t *testing.T) {
		mockCtrl, mockRepo, genreService := initMock()

		// Set up mock expectations and return values
		mockCtrl.On("CheckPrivilege", mock.Anything).Return(true)

		mockRepo.On("FindGenreByNameAndTypeCode", context.Background(), mock.Anything).
			Return(&model.Genre{}, nil)

		genreDtos := []dto.GenreDto{
			{Name: "Genre1", TypeCode: "INVALID"},
		}
		err := genreService.AddGenres(context.Background(), genreDtos)
		assert.Error(t, err)
		assert.Equal(t, errors.ErrInvalidInput.Error(), err.Error())
	})

	t.Run("Unauthorized user", func(t *testing.T) {
		mockCtrl, _, genreService := initMock()

		// Set up mock expectations and return values
		mockCtrl.On("CheckPrivilege", mock.Anything).Return(false)

		genreDtos := []dto.GenreDto{
			{Name: "Genre1", TypeCode: "MOVIE"},
		}
		err := genreService.AddGenres(context.Background(), genreDtos)
		assert.Error(t, err)
		assert.Equal(t, errors.ErrUnAuthorized.Error(), err.Error())
	})
}

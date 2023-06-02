package service

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"movies-service/internal/dto"
	"movies-service/internal/errors"
	"movies-service/internal/model"
	"movies-service/internal/season"
	"movies-service/internal/test/helper"
	"testing"
	"time"
)

func initMock() (*helper.MockManagementCtrl, *helper.MockMovieRepository, *helper.MockSeasonRepository, *helper.MockEpisodeRepository, season.Service) {
	mockCtrl := new(helper.MockManagementCtrl)
	mockMovieRepo := new(helper.MockMovieRepository)
	mockSeasonRepo := new(helper.MockSeasonRepository)
	mockEpisodeRepo := new(helper.MockEpisodeRepository)
	seasonService := NewSeasonService(mockCtrl, mockMovieRepo, mockSeasonRepo, mockEpisodeRepo)

	return mockCtrl, mockMovieRepo, mockSeasonRepo, mockEpisodeRepo, seasonService
}

func TestSeasonService_GetSeasonsByID(t *testing.T) {
	t.Run("", func(t *testing.T) {
		_, _, mockSeasonRepo, _, seasonService := initMock()

		mockSeason := &model.Season{ID: uint(1), Name: "Season 1"}
		expectedSeasonDto := &dto.SeasonDto{ID: uint(1), Name: "Season 1"}

		// Set up expectations
		mockSeasonRepo.On("FindSeasonByID", mock.Anything, uint(1)).Return(mockSeason, nil)

		result, err := seasonService.GetSeasonsByID(nil, uint(1))

		// Assertions
		assert.NoError(t, err)
		assert.Equal(t, expectedSeasonDto, result)
	})
}

func TestSeasonService_GetSeasonsByMovieID(t *testing.T) {
	_, _, mockSeasonRepo, _, seasonService := initMock()

	mockSeasons := []*model.Season{
		{ID: uint(1), Name: "Season 1"},
		{ID: uint(2), Name: "Season 2"},
	}
	expectedSeasonDtos := []*dto.SeasonDto{
		{ID: uint(1), Name: "Season 1"},
		{ID: uint(2), Name: "Season 2"},
	}

	// Set up expectations
	mockSeasonRepo.On("FindSeasonsByMovieID", mock.Anything, uint(1)).Return(mockSeasons, nil)

	// Call the service method
	result, err := seasonService.GetSeasonsByMovieID(context.Background(), uint(1))

	// Assertions
	assert.NoError(t, err)
	assert.Equal(t, expectedSeasonDtos, result)
}

func TestSeasonService_AddSeason(t *testing.T) {
	t.Run("Invalid Input", func(t *testing.T) {
		_, _, _, _, seasonService := initMock()

		seasonDto := &dto.SeasonDto{
			ID:          uint(1),
			Name:        "",
			Description: "",
			AirDate:     time.Now(),
			MovieID:     uint(0),
		}

		// Call the service method
		err := seasonService.AddSeason(context.Background(), seasonDto)

		// Assertions
		assert.Equal(t, errors.ErrInvalidInput, err)
	})

	t.Run("Unauthorized", func(t *testing.T) {
		mockCtrl, _, _, _, seasonService := initMock()

		seasonDto := &dto.SeasonDto{
			ID:          uint(0),
			Name:        "Season 1",
			Description: "Description",
			AirDate:     time.Now(),
			MovieID:     uint(1),
		}

		// Set up expectations
		mockCtrl.On("CheckPrivilege", mock.Anything).Return(false)

		// Call the service method
		err := seasonService.AddSeason(context.Background(), seasonDto)

		// Assertions
		assert.Equal(t, errors.ErrUnAuthorized, err)
	})

	t.Run("Success", func(t *testing.T) {
		mockCtrl, mockMovieRepo, mockSeasonRepo, _, seasonService := initMock()

		seasonDto := &dto.SeasonDto{
			ID:          uint(0),
			Name:        "Season 1",
			Description: "Description",
			AirDate:     time.Now(),
			MovieID:     uint(1),
		}
		movieObj := &model.Movie{ID: uint(1)}

		// Set up expectations
		mockCtrl.On("CheckPrivilege", mock.Anything).Return(true)
		mockMovieRepo.On("FindMovieById", mock.Anything, uint(1)).Return(movieObj, nil)
		mockSeasonRepo.On("InsertSeason", mock.Anything, mock.AnythingOfType("*model.Season")).Return(nil)

		// Call the service method
		err := seasonService.AddSeason(context.Background(), seasonDto)

		// Assertions
		assert.NoError(t, err)
	})
}

func TestSeasonService_UpdateSeason(t *testing.T) {
	t.Run("Invalid Input", func(t *testing.T) {
		_, _, _, _, seasonService := initMock()

		seasonDto := &dto.SeasonDto{
			ID:          uint(0),
			Name:        "",
			Description: "",
			AirDate:     time.Now(),
			MovieID:     uint(0),
		}

		// Call the service method
		err := seasonService.UpdateSeason(context.Background(), seasonDto)

		// Assertions
		assert.Equal(t, errors.ErrInvalidInput, err)
	})

	t.Run("Season Not Found", func(t *testing.T) {
		mockCtrl, _, mockSeasonRepo, _, seasonService := initMock()

		seasonDto := &dto.SeasonDto{
			ID:          uint(1),
			Name:        "Season 2",
			Description: "Updated description",
			AirDate:     time.Now(),
			MovieID:     uint(1),
		}

		// Set up expectations
		mockCtrl.On("CheckPrivilege", mock.Anything).Return(true)
		mockSeasonRepo.On("FindSeasonByID", mock.Anything, uint(1)).
			Return(nil, errors.ErrResourceNotFound)

		// Call the service method
		err := seasonService.UpdateSeason(context.Background(), seasonDto)

		// Assertions
		assert.Error(t, err)
		assert.Equal(t, errors.ErrResourceNotFound, err)
	})

	t.Run("Unauthorized", func(t *testing.T) {
		mockCtrl, _, mockSeasonRepo, _, seasonService := initMock()

		seasonDto := &dto.SeasonDto{
			ID:          uint(1),
			Name:        "Season 2",
			Description: "Updated description",
			AirDate:     time.Now(),
			MovieID:     uint(1),
		}

		// Set up expectations
		mockCtrl.On("CheckPrivilege", mock.Anything).Return(false)
		mockSeasonRepo.On("FindSeasonByID", mock.Anything, uint(1)).Return(&model.Season{}, nil)

		// Call the service method
		err := seasonService.UpdateSeason(context.Background(), seasonDto)

		// Assertions
		assert.Equal(t, errors.ErrUnAuthorized, err)

	})

	t.Run("Success", func(t *testing.T) {
		mockCtrl, mockMovieRepo, mockSeasonRepo, _, seasonService := initMock()

		seasonDto := &dto.SeasonDto{
			ID:          uint(1),
			Name:        "Season 2",
			Description: "Updated description",
			AirDate:     time.Now(),
			MovieID:     uint(1),
		}
		mockSeason := &model.Season{
			ID:          uint(1),
			Name:        "Season 1",
			Description: "Description",
			AirDate:     time.Now(),
			Movie:       &model.Movie{ID: uint(1)},
		}

		// Set up expectations
		mockCtrl.On("CheckPrivilege", mock.Anything).Return(true)
		mockSeasonRepo.On("FindSeasonByID", mock.Anything, uint(1)).Return(mockSeason, nil)
		mockMovieRepo.On("FindMovieById", mock.Anything, uint(1)).Return(mockSeason.Movie, nil)
		mockSeasonRepo.On("UpdateSeason", mock.Anything, mock.AnythingOfType("*model.Season")).Return(nil)

		// Call the service method
		err := seasonService.UpdateSeason(context.Background(), seasonDto)

		// Assertions
		assert.NoError(t, err)
	})
}

func TestSeasonService_RemoveSeasonByID(t *testing.T) {
	t.Run("Invalid Input", func(t *testing.T) {
		_, _, _, _, seasonService := initMock()

		// Call the service method
		err := seasonService.RemoveSeasonByID(context.Background(), uint(0))

		// Assertions
		assert.Equal(t, errors.ErrInvalidInput, err)
	})

	t.Run("Unauthorized", func(t *testing.T) {
		mockCtrl, _, _, _, seasonService := initMock()

		// Set up expectations
		mockCtrl.On("CheckPrivilege", mock.Anything).Return(false)

		// Call the service method
		err := seasonService.RemoveSeasonByID(context.Background(), uint(1))

		// Assertions
		assert.Equal(t, errors.ErrUnAuthorized, err)

	})

	t.Run("Error Deleting Episode", func(t *testing.T) {
		mockCtrl, _, _, mockEpisodeRepo, seasonService := initMock()

		// Set up expectations
		mockCtrl.On("CheckPrivilege", mock.Anything).Return(true)
		mockEpisodeRepo.On("DeleteEpisodeBySeasonID", mock.Anything, uint(1)).
			Return(fmt.Errorf("failed to delete episodes"))

		// Call the service method
		err := seasonService.RemoveSeasonByID(context.Background(), uint(1))

		// Assertions
		assert.EqualError(t, err, "failed to delete episodes")

	})

	t.Run("Error Deleting Season", func(t *testing.T) {
		mockCtrl, _, mockSeasonRepo, mockEpisodeRepo, seasonService := initMock()

		// Set up expectations
		mockCtrl.On("CheckPrivilege", mock.Anything).Return(true)
		mockEpisodeRepo.On("DeleteEpisodeBySeasonID", mock.Anything, uint(1)).Return(nil)
		mockSeasonRepo.On("DeleteSeasonByID", mock.Anything, uint(1)).Return(fmt.Errorf("failed to delete season"))

		// Call the service method
		err := seasonService.RemoveSeasonByID(context.Background(), uint(1))

		// Assertions
		assert.EqualError(t, err, "failed to delete season")

	})

	t.Run("Success", func(t *testing.T) {
		mockCtrl, _, mockSeasonRepo, mockEpisodeRepo, seasonService := initMock()

		// Set up expectations
		mockCtrl.On("CheckPrivilege", mock.Anything).Return(true)
		mockEpisodeRepo.On("DeleteEpisodeBySeasonID", mock.Anything, uint(1)).Return(nil)
		mockSeasonRepo.On("DeleteSeasonByID", mock.Anything, uint(1)).Return(nil)

		// Call the service method
		err := seasonService.RemoveSeasonByID(context.Background(), uint(1))

		// Assertions
		assert.NoError(t, err)

	})
}

package service

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"movies-service/internal/dto"
	"movies-service/internal/episode"
	"movies-service/internal/errors"
	"movies-service/internal/model"
	"movies-service/internal/test/helper"
	"testing"
	"time"
)

func initMock() (*helper.MockManagementCtrl, *helper.MockSeasonRepository, *helper.MockEpisodeRepository, episode.Service) {
	mockCtrl := new(helper.MockManagementCtrl)
	mockSeasonRepo := new(helper.MockSeasonRepository)
	mockEpisodeRepo := new(helper.MockEpisodeRepository)

	episodeService := NewEpisodeService(mockCtrl, mockSeasonRepo, mockEpisodeRepo)

	return mockCtrl, mockSeasonRepo, mockEpisodeRepo, episodeService
}

func TestEpisodeService_GetEpisodesByID(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		mockCrtl, _, mockEpisodeRepo, episodeService := initMock()

		// Set up expectations
		episodeID := uint(1)
		mockEpisode := &model.Episode{
			ID:        episodeID,
			Name:      "Episode 1",
			AirDate:   time.Now(),
			Runtime:   uint(60),
			VideoPath: sql.NullString{String: "/path/to/episode1", Valid: true},
			SeasonID:  uint(1),
			CreatedAt: time.Now(),
			CreatedBy: "John",
			UpdatedAt: time.Now(),
			UpdatedBy: "John",
		}

		mockCrtl.On("CheckUser", mock.Anything).Return(false, false)

		mockEpisodeRepo.On("FindEpisodeByID", context.Background(), episodeID).
			Return(mockEpisode, nil)

		// Call the service method
		result, err := episodeService.GetEpisodesByID(context.Background(), episodeID)

		// Assertions
		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, episodeID, result.ID)
		assert.Equal(t, mockEpisode.Name, result.Name)
		assert.Equal(t, mockEpisode.AirDate, result.AirDate)
		assert.Equal(t, mockEpisode.Runtime, result.Runtime)
		assert.Equal(t, "", result.VideoPath)
		assert.Equal(t, mockEpisode.SeasonID, result.SeasonID)

	})
}

func TestEpisodeService_GetEpisodesBySeasonID(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		mockCrtl, _, mockEpisodeRepo, episodeService := initMock()

		// Set up expectations
		seasonID := uint(1)
		mockEpisodes := []*model.Episode{
			{
				ID:        uint(1),
				Name:      "Episode 1",
				AirDate:   time.Now(),
				Runtime:   uint(60),
				VideoPath: sql.NullString{String: "/path/to/episode1", Valid: true},
				SeasonID:  seasonID,
			},
			{
				ID:        uint(2),
				Name:      "Episode 2",
				AirDate:   time.Now(),
				Runtime:   uint(45),
				VideoPath: sql.NullString{String: "/path/to/episode2", Valid: true},
				SeasonID:  seasonID,
			},
		}
		mockCrtl.On("CheckUser", mock.Anything).Return(true, false)

		mockEpisodeRepo.On("FindEpisodesBySeasonID", context.Background(), seasonID).Return(mockEpisodes, nil)

		// Call the service method
		result, err := episodeService.GetEpisodesBySeasonID(context.Background(), seasonID)

		// Assertions
		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, len(mockEpisodes), len(result))

	})
}

func TestEpisodeService_AddEpisode(t *testing.T) {
	t.Run("Invalid Input", func(t *testing.T) {
		_, _, _, episodeService := initMock()

		// Set up input parameters
		episodeDto := &dto.EpisodeDto{
			ID:        1,
			Name:      "",
			AirDate:   time.Now(),
			Runtime:   uint(60),
			VideoPath: "/path/to/episode1",
			SeasonID:  uint(1),
		}

		// Call the service method
		err := episodeService.AddEpisode(context.Background(), episodeDto)

		// Assertions
		assert.Error(t, err)
		assert.Equal(t, errors.ErrInvalidInput, err)
	})

	t.Run("Unauthorized", func(t *testing.T) {
		mockCtrl, _, _, episodeService := initMock()

		// Set up input parameters
		episodeDto := &dto.EpisodeDto{
			Name:      "Episode 1",
			AirDate:   time.Now(),
			Runtime:   uint(60),
			VideoPath: "/path/to/episode1",
			SeasonID:  uint(1),
		}

		// Set up expectations
		mockCtrl.On("CheckPrivilege", mock.Anything).Return(false)

		// Call the service method
		err := episodeService.AddEpisode(context.Background(), episodeDto)

		// Assertions
		assert.Error(t, err)
		assert.Equal(t, errors.ErrUnAuthorized, err)
	})

	t.Run("Not Found Season", func(t *testing.T) {
		mockCtrl, mockSeasonRepo, _, episodeService := initMock()

		// Set up input parameters
		episodeDto := &dto.EpisodeDto{
			Name:      "Episode 1",
			AirDate:   time.Now(),
			Runtime:   uint(60),
			VideoPath: "/path/to/episode1",
			SeasonID:  uint(1),
		}

		// Set up expectations
		mockCtrl.On("CheckPrivilege", mock.Anything).Return(true)
		mockSeasonRepo.On("FindSeasonByID", context.Background(), episodeDto.SeasonID).Return(nil, fmt.Errorf("failed to find season"))

		// Call the service method
		err := episodeService.AddEpisode(context.Background(), episodeDto)

		// Assertions
		assert.Error(t, err)
		assert.EqualError(t, err, "failed to find season")
	})

	t.Run("Success", func(t *testing.T) {
		mockCtrl, mockSeasonRepo, mockEpisodeRepo, episodeService := initMock()

		// Set up input parameters
		episodeDto := &dto.EpisodeDto{
			Name:      "Episode 1",
			AirDate:   time.Now(),
			Runtime:   uint(60),
			VideoPath: "/path/to/episode1",
			SeasonID:  uint(1),
		}

		// Set up expectations
		mockCtrl.On("CheckPrivilege", mock.Anything).Return(true)
		mockSeasonRepo.On("FindSeasonByID", context.Background(), episodeDto.SeasonID).Return(&model.Season{}, nil)
		mockEpisodeRepo.On("InsertEpisode", context.Background(), mock.AnythingOfType("*model.Episode")).Return(nil)

		// Call the service method
		err := episodeService.AddEpisode(context.Background(), episodeDto)

		// Assertions
		assert.NoError(t, err)

	})
}

func TestEpisodeService_UpdateEpisode(t *testing.T) {
	t.Run("Invalid Input", func(t *testing.T) {
		_, _, _, episodeService := initMock()

		// Set up input parameters
		episodeDto := &dto.EpisodeDto{
			ID:        uint(0),
			Name:      "",
			AirDate:   time.Now(),
			Runtime:   uint(60),
			VideoPath: "/path/to/episode1",
			SeasonID:  uint(1),
		}

		// Call the service method
		err := episodeService.UpdateEpisode(context.Background(), episodeDto)

		// Assertions
		assert.Error(t, err)
		assert.Equal(t, errors.ErrInvalidInput, err)
	})

	t.Run("Episode Not Found", func(t *testing.T) {
		mockCtrl, _, mockEpisodeRepo, episodeService := initMock()

		// Set up input parameters
		episodeDto := &dto.EpisodeDto{
			ID:        uint(1),
			Name:      "Episode 1",
			AirDate:   time.Now(),
			Runtime:   uint(60),
			VideoPath: "/path/to/episode1",
			SeasonID:  uint(1),
		}

		// Set up expectations
		mockCtrl.On("CheckPrivilege", mock.Anything).Return(true)
		mockEpisodeRepo.On("FindEpisodeByID", mock.Anything, uint(1)).
			Return(nil, errors.ErrResourceNotFound)

		// Call the service method
		err := episodeService.UpdateEpisode(context.Background(), episodeDto)

		// Assertions
		assert.Error(t, err)
		assert.Equal(t, errors.ErrResourceNotFound, err)

	})

	t.Run("Unauthorized", func(t *testing.T) {
		mockCtrl, _, mockEpisodeRepo, episodeService := initMock()

		// Set up input parameters
		episodeDto := &dto.EpisodeDto{
			ID:        uint(1),
			Name:      "Episode 1",
			AirDate:   time.Now(),
			Runtime:   uint(60),
			VideoPath: "/path/to/episode1",
			SeasonID:  uint(1),
		}

		// Set up expectations
		mockEpisodeRepo.On("FindEpisodeByID", context.Background(), episodeDto.ID).
			Return(&model.Episode{}, nil)
		mockCtrl.On("CheckPrivilege", mock.Anything).Return(false)

		// Call the service method
		err := episodeService.UpdateEpisode(context.Background(), episodeDto)

		// Assertions
		assert.Error(t, err)
		assert.Equal(t, errors.ErrUnAuthorized, err)

	})

	t.Run("Error Getting Season", func(t *testing.T) {
		mockCtrl, mockSeasonRepo, mockEpisodeRepo, episodeService := initMock()

		// Set up input parameters
		episodeDto := &dto.EpisodeDto{
			ID:        uint(1),
			Name:      "Episode 1",
			AirDate:   time.Now(),
			Runtime:   uint(60),
			VideoPath: "/path/to/episode1",
			SeasonID:  uint(1),
		}

		// Set up expectations
		mockCtrl.On("CheckPrivilege", mock.Anything).Return(true)
		mockEpisodeRepo.On("FindEpisodeByID", context.Background(), episodeDto.ID).Return(&model.Episode{}, nil)
		mockSeasonRepo.On("FindSeasonByID", context.Background(), episodeDto.SeasonID).
			Return(nil, fmt.Errorf("failed to find season"))

		// Call the service method
		err := episodeService.UpdateEpisode(context.Background(), episodeDto)

		// Assertions
		assert.Error(t, err)
		assert.EqualError(t, err, "failed to find season")
	})

	t.Run("Success", func(t *testing.T) {
		mockCtrl, mockSeasonRepo, mockEpisodeRepo, episodeService := initMock()

		// Set up input parameters
		episodeDto := &dto.EpisodeDto{
			ID:        uint(1),
			Name:      "Episode 1",
			AirDate:   time.Now(),
			Runtime:   uint(60),
			VideoPath: "/path/to/episode1",
			SeasonID:  uint(1),
		}

		// Set up expectations
		mockCtrl.On("CheckPrivilege", mock.Anything).Return(true)
		mockEpisodeRepo.On("FindEpisodeByID", context.Background(), episodeDto.ID).Return(&model.Episode{}, nil)
		mockSeasonRepo.On("FindSeasonByID", context.Background(), episodeDto.SeasonID).Return(&model.Season{}, nil)
		mockEpisodeRepo.On("UpdateEpisode", context.Background(), mock.AnythingOfType("*model.Episode")).Return(nil)

		// Call the service method
		err := episodeService.UpdateEpisode(context.Background(), episodeDto)

		// Assertions
		assert.NoError(t, err)

	})
}

func TestEpisodeService_RemoveEpisodeByID(t *testing.T) {
	t.Run("Invalid Input", func(t *testing.T) {
		_, _, _, episodeService := initMock()

		// Set up input parameters
		episodeID := uint(0)

		// Call the service method
		err := episodeService.RemoveEpisodeByID(context.Background(), episodeID)

		// Assertions
		assert.Error(t, err)
		assert.Equal(t, errors.ErrInvalidInput, err)
	})

	t.Run("Unauthorized", func(t *testing.T) {
		mockCtrl, _, _, episodeService := initMock()

		// Set up input parameters
		episodeID := uint(1)

		// Set up expectations
		mockCtrl.On("CheckPrivilege", mock.Anything).Return(false)

		// Call the service method
		err := episodeService.RemoveEpisodeByID(context.Background(), episodeID)

		// Assertions
		assert.Error(t, err)
		assert.Equal(t, errors.ErrUnAuthorized, err)
	})

	t.Run("Error Deleting", func(t *testing.T) {
		mockCtrl, _, mockEpisodeRepo, episodeService := initMock()

		// Set up input parameters
		episodeID := uint(1)

		// Set up expectations
		mockCtrl.On("CheckPrivilege", mock.Anything).Return(true)
		mockEpisodeRepo.On("DeleteEpisodeByID", context.Background(), episodeID).Return(fmt.Errorf("failed to delete episode"))

		// Call the service method
		err := episodeService.RemoveEpisodeByID(context.Background(), episodeID)

		// Assertions
		assert.Error(t, err)
		assert.EqualError(t, err, "failed to delete episode")
	})

	t.Run("Success", func(t *testing.T) {
		mockCtrl, _, mockEpisodeRepo, episodeService := initMock()

		// Set up input parameters
		episodeID := uint(1)

		// Set up expectations
		mockCtrl.On("CheckPrivilege", mock.Anything).Return(true)
		mockEpisodeRepo.On("DeleteEpisodeByID", context.Background(), episodeID).Return(nil)

		// Call the service method
		err := episodeService.RemoveEpisodeByID(context.Background(), episodeID)

		// Assertions
		assert.NoError(t, err)
	})
}

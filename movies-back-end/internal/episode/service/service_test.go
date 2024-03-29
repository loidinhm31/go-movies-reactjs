package service

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"movies-service/internal/common/dto"
	"movies-service/internal/common/entity"
	"movies-service/internal/episode"
	"movies-service/internal/errors"
	"movies-service/internal/middlewares"
	"movies-service/internal/test/helper"
	"testing"
	"time"
)

func initMock() (*helper.MockManagementCtrl, *helper.MockUserRepository, *helper.MockSeasonRepository, *helper.MockCollectionRepository, *helper.MockPaymentRepository, *helper.MockEpisodeRepository, *helper.MockBlobService, episode.Service) {
	mockCtrl := new(helper.MockManagementCtrl)
	mockUserRepo := new(helper.MockUserRepository)
	mockSeasonRepo := new(helper.MockSeasonRepository)
	mockCollectionRepo := new(helper.MockCollectionRepository)
	mockPaymentRepo := new(helper.MockPaymentRepository)
	mockEpisodeRepo := new(helper.MockEpisodeRepository)
	mockBlobSvc := new(helper.MockBlobService)

	episodeSvc := NewEpisodeService(mockCtrl, mockUserRepo, mockSeasonRepo, mockCollectionRepo, mockPaymentRepo, mockEpisodeRepo, mockBlobSvc)

	return mockCtrl, mockUserRepo, mockSeasonRepo, mockCollectionRepo, mockPaymentRepo, mockEpisodeRepo, mockBlobSvc, episodeSvc
}

func TestEpisodeService_GetEpisodesByID(t *testing.T) {
	t.Run("Invalid Client", func(t *testing.T) {
		_, mockUserRepo, _, _, _, mockEpisodeRepo, _, episodeSvc := initMock()

		episodeID := uint(1)
		mockEpisode := &entity.Episode{
			ID:   episodeID,
			Name: "Episode 1",
			Price: sql.NullFloat64{
				Float64: 10.5,
				Valid:   true,
			},
			AirDate:   time.Now(),
			Runtime:   uint(60),
			VideoPath: sql.NullString{String: "/path/to/episode1", Valid: true},
			SeasonID:  uint(1),
			CreatedAt: time.Now(),
			CreatedBy: "John",
			UpdatedAt: time.Now(),
			UpdatedBy: "John",
		}

		mockEpisodeRepo.On("FindEpisodeByID", mock.Anything, mock.Anything).
			Return(mockEpisode, nil)

		mockUserRepo.On("FindUserByUsernameAndIsNew", mock.Anything, mock.Anything, mock.Anything).
			Return(&entity.User{
				Role: &entity.Role{RoleCode: "BANNED"},
			}, nil)

		// Call the service method
		result, err := episodeSvc.GetEpisodeByID(context.Background(), uint(1))

		// Assertions
		assert.NoError(t, err)
		assert.Equal(t, result.Name, mockEpisode.Name)
	})

	t.Run("Episode Has Price, Not Paid", func(t *testing.T) {
		_, mockUserRepo, _, _, mockPaymentRepo, mockEpisodeRepo, _, episodeSvc := initMock()

		// Set up expectations
		author := "test"
		ctxWithAuthor := context.WithValue(context.Background(), middlewares.CtxUserKey, author)

		movieType := "TV"
		episodeID := uint(1)
		mockEpisode := &entity.Episode{
			ID:   episodeID,
			Name: "Episode 1",
			Price: sql.NullFloat64{
				Float64: 10.5,
				Valid:   true,
			},
			AirDate:   time.Now(),
			Runtime:   uint(60),
			VideoPath: sql.NullString{String: "/path/to/episode1", Valid: true},
			SeasonID:  uint(1),
			CreatedAt: time.Now(),
			CreatedBy: "John",
			UpdatedAt: time.Now(),
			UpdatedBy: "John",
		}

		mockEpisodeRepo.On("FindEpisodeByID", ctxWithAuthor, mock.Anything).
			Return(mockEpisode, nil)

		mockUserRepo.On("FindUserByUsernameAndIsNew", ctxWithAuthor, author, mock.Anything).
			Return(&entity.User{
				Role: &entity.Role{RoleCode: "GENERAL"},
			}, nil)

		mockPaymentRepo.On("FindPaymentByUserIDAndTypeCodeAndRefID",
			ctxWithAuthor, mock.Anything, movieType, mock.Anything).
			Return(&entity.Payment{}, nil)

		// Call the service method
		result, err := episodeSvc.GetEpisodeByID(ctxWithAuthor, episodeID)

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

	t.Run("Success", func(t *testing.T) {
		_, mockUserRepo, _, _, _, mockEpisodeRepo, _, episodeSvc := initMock()

		// Set up expectations
		author := "test"
		ctxWithAuthor := context.WithValue(context.Background(), middlewares.CtxUserKey, author)

		episodeID := uint(1)
		mockEpisode := &entity.Episode{
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

		mockEpisodeRepo.On("FindEpisodeByID", ctxWithAuthor, episodeID).
			Return(mockEpisode, nil)

		mockUserRepo.On("FindUserByUsernameAndIsNew", ctxWithAuthor, author, mock.Anything).
			Return(&entity.User{
				Role: &entity.Role{RoleCode: "GENERAL"},
			}, nil)

		// Call the service method
		result, err := episodeSvc.GetEpisodeByID(ctxWithAuthor, episodeID)

		// Assertions
		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, episodeID, result.ID)
		assert.Equal(t, mockEpisode.Name, result.Name)
		assert.Equal(t, mockEpisode.AirDate, result.AirDate)
		assert.Equal(t, mockEpisode.Runtime, result.Runtime)
		assert.Equal(t, "/path/to/episode1", result.VideoPath)
		assert.Equal(t, mockEpisode.SeasonID, result.SeasonID)

	})
}

func TestEpisodeService_GetEpisodesBySeasonID(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		mockCrtl, _, _, _, _, mockEpisodeRepo, _, episodeSvc := initMock()

		// Set up expectations
		seasonID := uint(1)
		mockEpisodes := []*entity.Episode{
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
		result, err := episodeSvc.GetEpisodesBySeasonID(context.Background(), seasonID)

		// Assertions
		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, len(mockEpisodes), len(result))

	})
}

func TestEpisodeService_AddEpisode(t *testing.T) {
	t.Run("Invalid Input", func(t *testing.T) {
		_, _, _, _, _, _, _, episodeSvc := initMock()

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
		err := episodeSvc.AddEpisode(context.Background(), episodeDto)

		// Assertions
		assert.Error(t, err)
		assert.Equal(t, errors.ErrInvalidInput, err)
	})

	t.Run("Unauthorized", func(t *testing.T) {
		mockCtrl, _, _, _, _, _, _, episodeSvc := initMock()

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
		err := episodeSvc.AddEpisode(context.Background(), episodeDto)

		// Assertions
		assert.Error(t, err)
		assert.Equal(t, errors.ErrUnAuthorized, err)
	})

	t.Run("Not Found Season", func(t *testing.T) {
		mockCtrl, _, mockSeasonRepo, _, _, _, _, episodeSvc := initMock()

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
		err := episodeSvc.AddEpisode(context.Background(), episodeDto)

		// Assertions
		assert.Error(t, err)
		assert.EqualError(t, err, "failed to find season")
	})

	t.Run("Success", func(t *testing.T) {
		mockCtrl, _, mockSeasonRepo, _, _, mockEpisodeRepo, _, episodeSvc := initMock()

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
		mockSeasonRepo.On("FindSeasonByID", context.Background(), mock.Anything).
			Return(&entity.Season{}, nil)
		mockEpisodeRepo.On("InsertEpisode", context.Background(), mock.Anything).
			Return(nil)

		// Call the service method
		err := episodeSvc.AddEpisode(context.Background(), episodeDto)

		// Assertions
		assert.NoError(t, err)

	})
}

func TestEpisodeService_UpdateEpisode(t *testing.T) {
	t.Run("Invalid Input", func(t *testing.T) {
		_, _, _, _, _, _, _, episodeSvc := initMock()

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
		err := episodeSvc.UpdateEpisode(context.Background(), episodeDto)

		// Assertions
		assert.Error(t, err)
		assert.Equal(t, errors.ErrInvalidInput, err)
	})

	t.Run("Episode Not Found", func(t *testing.T) {
		mockCtrl, _, _, _, _, mockEpisodeRepo, _, episodeSvc := initMock()

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
		err := episodeSvc.UpdateEpisode(context.Background(), episodeDto)

		// Assertions
		assert.Error(t, err)
		assert.Equal(t, errors.ErrResourceNotFound, err)

	})

	t.Run("Unauthorized", func(t *testing.T) {
		mockCtrl, _, _, _, _, mockEpisodeRepo, _, episodeSvc := initMock()

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
			Return(&entity.Episode{}, nil)
		mockCtrl.On("CheckPrivilege", mock.Anything).Return(false)

		// Call the service method
		err := episodeSvc.UpdateEpisode(context.Background(), episodeDto)

		// Assertions
		assert.Error(t, err)
		assert.Equal(t, errors.ErrUnAuthorized, err)

	})

	t.Run("Error Getting Season", func(t *testing.T) {
		mockCtrl, _, mockSeasonRepo, _, _, mockEpisodeRepo, _, episodeSvc := initMock()

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
		mockEpisodeRepo.On("FindEpisodeByID", context.Background(), episodeDto.ID).Return(&entity.Episode{}, nil)
		mockSeasonRepo.On("FindSeasonByID", context.Background(), episodeDto.SeasonID).
			Return(nil, fmt.Errorf("failed to find season"))

		// Call the service method
		err := episodeSvc.UpdateEpisode(context.Background(), episodeDto)

		// Assertions
		assert.Error(t, err)
		assert.EqualError(t, err, "failed to find season")
	})

	t.Run("Success", func(t *testing.T) {
		mockCtrl, _, mockSeasonRepo, _, _, mockEpisodeRepo, _, episodeSvc := initMock()

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
		mockEpisodeRepo.On("FindEpisodeByID", context.Background(), mock.Anything).
			Return(&entity.Episode{}, nil)
		mockSeasonRepo.On("FindSeasonByID", context.Background(), mock.Anything).
			Return(&entity.Season{}, nil)
		mockEpisodeRepo.On("UpdateEpisode", context.Background(), mock.Anything).Return(nil)

		// Call the service method
		err := episodeSvc.UpdateEpisode(context.Background(), episodeDto)

		// Assertions
		assert.NoError(t, err)

	})
}

func TestEpisodeService_RemoveEpisodeByID(t *testing.T) {
	t.Run("Invalid Input", func(t *testing.T) {
		_, _, _, _, _, _, _, episodeSvc := initMock()

		// Set up input parameters
		episodeID := uint(0)

		// Call the service method
		err := episodeSvc.RemoveEpisodeByID(context.Background(), episodeID)

		// Assertions
		assert.Error(t, err)
		assert.Equal(t, errors.ErrInvalidInput, err)
	})

	t.Run("Added to Payment", func(t *testing.T) {
		_, _, _, _, mockPaymentRepo, _, _, episodeSvc := initMock()

		// Set up input parameters
		episodeID := uint(1)

		mockPaymentRepo.On("FindPaymentsByTypeCodeAndRefID", mock.Anything, "TV", mock.Anything).
			Return([]*entity.Payment{
				{
					ID: uint(1),
				},
			}, nil)

		// Call the service method
		err := episodeSvc.RemoveEpisodeByID(context.Background(), episodeID)

		// Assertions
		assert.Error(t, err)
		assert.Equal(t, errors.ErrCannotExecuteAction, err)
	})

	t.Run("Added to Collection", func(t *testing.T) {
		_, _, _, mockCollection, mockPaymentRepo, _, _, episodeSvc := initMock()

		// Set up input parameters
		episodeID := uint(1)

		mockPaymentRepo.On("FindPaymentsByTypeCodeAndRefID", mock.Anything, "TV", mock.Anything).
			Return([]*entity.Payment{}, nil)

		mockCollection.On("FindCollectionsByEpisodeID", mock.Anything, mock.Anything).
			Return([]*entity.Collection{
				{
					ID:     uint(2),
					UserID: uint(1),
					EpisodeID: sql.NullInt64{
						Int64: int64(1),
						Valid: false,
					},
					TypeCode: "TV",
				},
			}, nil)

		// Call the service method
		err := episodeSvc.RemoveEpisodeByID(context.Background(), episodeID)

		// Assertions
		assert.Error(t, err)
		assert.Equal(t, errors.ErrCannotExecuteAction, err)
	})

	t.Run("Invalid Client", func(t *testing.T) {
		mockCtrl, _, _, mockCollection, mockPaymentRepo, _, _, episodeSvc := initMock()

		// Set up input parameters
		episodeID := uint(1)

		// Set up expectations
		mockPaymentRepo.On("FindPaymentsByTypeCodeAndRefID", mock.Anything, "TV", mock.Anything).
			Return([]*entity.Payment{}, nil)

		mockCollection.On("FindCollectionsByEpisodeID", mock.Anything, mock.Anything).
			Return([]*entity.Collection{}, nil)

		mockCollection.On("FindCollectionsByEpisodeID", mock.Anything, mock.Anything).
			Return([]*entity.Collection{}, nil)

		mockCtrl.On("CheckPrivilege", mock.Anything).Return(false)

		// Call the service method
		err := episodeSvc.RemoveEpisodeByID(context.Background(), episodeID)

		// Assertions
		assert.Error(t, err)
		assert.Equal(t, errors.ErrInvalidClient, err)
	})

	t.Run("Error Deleting", func(t *testing.T) {
		mockCtrl, _, _, mockCollection, mockPaymentRepo, mockEpisodeRepo, _, episodeSvc := initMock()

		// Set up input parameters
		episodeID := uint(1)

		// Set up expectations
		mockPaymentRepo.On("FindPaymentsByTypeCodeAndRefID", mock.Anything, "TV", mock.Anything).
			Return([]*entity.Payment{}, nil)

		mockCollection.On("FindCollectionsByEpisodeID", mock.Anything, mock.Anything).
			Return([]*entity.Collection{}, nil)

		mockCtrl.On("CheckPrivilege", mock.Anything).Return(true)

		mockEpisodeRepo.On("FindEpisodeByID", mock.Anything, mock.Anything).
			Return(&entity.Episode{}, nil)

		mockEpisodeRepo.On("DeleteEpisodeByID", mock.Anything, mock.Anything).
			Return(fmt.Errorf("failed to delete episode"))

		// Call the service method
		err := episodeSvc.RemoveEpisodeByID(context.Background(), episodeID)

		// Assertions
		assert.Error(t, err)
		assert.EqualError(t, err, "failed to delete episode")
	})

	t.Run("Success", func(t *testing.T) {
		mockCtrl, _, _, mockCollection, mockPaymentRepo, mockEpisodeRepo, blobService, episodeSvc := initMock()

		// Set up input parameters
		episodeID := uint(1)

		// Set up expectations
		mockPaymentRepo.On("FindPaymentsByTypeCodeAndRefID", mock.Anything, "TV", mock.Anything).
			Return([]*entity.Payment{}, nil)

		mockCollection.On("FindCollectionsByEpisodeID", mock.Anything, mock.Anything).
			Return([]*entity.Collection{}, nil)

		mockCtrl.On("CheckPrivilege", mock.Anything).Return(true)

		mockEpisodeRepo.On("FindEpisodeByID", mock.Anything, mock.Anything).
			Return(&entity.Episode{
				VideoPath: sql.NullString{
					String: "/video/path",
					Valid:  true,
				},
			}, nil)

		blobService.On("DeleteFile", mock.Anything, mock.Anything, "video").
			Return("result", nil)

		mockEpisodeRepo.On("DeleteEpisodeByID", context.Background(), episodeID).Return(nil)

		// Call the service method
		err := episodeSvc.RemoveEpisodeByID(context.Background(), episodeID)

		// Assertions
		assert.NoError(t, err)
	})
}

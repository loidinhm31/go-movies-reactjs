package service

import (
	"context"
	"database/sql"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"movies-service/internal/collection"
	"movies-service/internal/common/dto"
	"movies-service/internal/common/entity"
	"movies-service/internal/errors"
	"movies-service/internal/test/helper"
	"movies-service/pkg/pagination"
	"testing"
)

func initMock() (*helper.MockUserRepository, *helper.MockMovieRepository, *helper.MockEpisodeRepository, *helper.MockPaymentRepository, *helper.MockCollectionRepository, collection.Service) {
	mockUserRepo := new(helper.MockUserRepository)
	mockMovieRepo := new(helper.MockMovieRepository)
	mockEpisodeRepo := new(helper.MockEpisodeRepository)
	mockPaymentRepo := new(helper.MockPaymentRepository)
	mockCollectionRepo := new(helper.MockCollectionRepository)

	collectionService := NewCollectionService(mockUserRepo, mockMovieRepo, mockEpisodeRepo, mockPaymentRepo, mockCollectionRepo)

	return mockUserRepo, mockMovieRepo, mockEpisodeRepo, mockPaymentRepo, mockCollectionRepo, collectionService
}

func TestCollectionService_AddCollection(t *testing.T) {
	t.Run("Invalid Type Code", func(t *testing.T) {
		_, _, _, _, _, collectionService := initMock()

		err := collectionService.AddCollection(context.Background(), &dto.CollectionDto{})
		assert.Error(t, err)
		assert.Equal(t, errors.ErrInvalidInput, err)
	})

	t.Run("Invalid Client", func(t *testing.T) {
		mockUserRepo, _, _, _, _, collectionService := initMock()

		mockUserRepo.On("FindUserByUsername", mock.Anything, mock.Anything).
			Return(&entity.User{Role: &entity.Role{RoleCode: "BANNED"}}, nil)

		err := collectionService.AddCollection(context.Background(), &dto.CollectionDto{
			TypeCode: "MOVIE",
		})
		assert.Error(t, err)
		assert.Equal(t, errors.ErrInvalidClient, err)

	})

	t.Run("Type Movie, Error Payment Not Found", func(t *testing.T) {
		mockUserRepo, mockMovieRepo, _, mockPaymentRepo, mockCollectionRepo, collectionService := initMock()

		mockUserRepo.On("FindUserByUsername", mock.Anything, mock.Anything).
			Return(&entity.User{
				ID:   uint(10),
				Role: &entity.Role{RoleCode: "GENERAL"},
			}, nil)

		mockMovieRepo.On("FindMovieByID", mock.Anything, mock.Anything).
			Return(&entity.Movie{
				ID:    uint(123),
				Price: sql.NullFloat64{Float64: 1.2, Valid: true},
			}, nil)

		mockPaymentRepo.On("FindPaymentByUserIDAndTypeCodeAndRefID", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
			Return(&entity.Payment{
				RefID:  uint(1),
				UserID: uint(9),
			}, nil)

		mockCollectionRepo.On("InsertCollection", mock.Anything, mock.Anything).
			Return(nil)

		err := collectionService.AddCollection(context.Background(), &dto.CollectionDto{
			TypeCode: "MOVIE",
			MovieID:  uint(123),
		})

		assert.Error(t, err)
		assert.Equal(t, errors.ErrPaymentNotFound, err)
	})

	t.Run("Type Movie, Inserted", func(t *testing.T) {
		mockUserRepo, mockMovieRepo, _, mockPaymentRepo, mockCollectionRepo, collectionService := initMock()

		refID := uint(123)
		userID := uint(10)

		mockUserRepo.On("FindUserByUsername", mock.Anything, mock.Anything).
			Return(&entity.User{
				ID:   userID,
				Role: &entity.Role{RoleCode: "GENERAL"},
			}, nil)

		mockMovieRepo.On("FindMovieByID", mock.Anything, mock.Anything).
			Return(&entity.Movie{
				ID:    refID,
				Price: sql.NullFloat64{Float64: 1.2, Valid: true},
			}, nil)

		mockPaymentRepo.On("FindPaymentByUserIDAndTypeCodeAndRefID", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
			Return(&entity.Payment{
				RefID:  refID,
				UserID: userID,
			}, nil)

		mockPaymentRepo.On("InsertPayment", mock.Anything, mock.Anything).
			Return(&entity.Payment{}, nil)

		mockCollectionRepo.On("InsertCollection", mock.Anything, mock.Anything).
			Return(nil)

		err := collectionService.AddCollection(context.Background(), &dto.CollectionDto{
			TypeCode: "MOVIE",
			MovieID:  refID,
		})

		assert.NoError(t, err)
	})

	t.Run("Type TV, Error Payment Not Found", func(t *testing.T) {
		mockUserRepo, _, mockEpisodeRepo, mockPaymentRepo, mockCollectionRepo, collectionService := initMock()

		refID := uint(123)
		userID := uint(10)

		mockUserRepo.On("FindUserByUsername", mock.Anything, mock.Anything).
			Return(&entity.User{
				ID:   userID,
				Role: &entity.Role{RoleCode: "GENERAL"},
			}, nil)

		mockEpisodeRepo.On("FindEpisodeByID", mock.Anything, mock.Anything).
			Return(&entity.Episode{
				ID:    refID,
				Price: sql.NullFloat64{Float64: 1.2, Valid: true},
			}, nil)

		mockPaymentRepo.On("FindPaymentByUserIDAndTypeCodeAndRefID", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
			Return(&entity.Payment{
				RefID:  refID,
				UserID: uint(9),
			}, nil)

		mockCollectionRepo.On("InsertCollection", mock.Anything, mock.Anything).
			Return(nil)

		err := collectionService.AddCollection(context.Background(), &dto.CollectionDto{
			TypeCode: "TV",
			MovieID:  uint(123),
		})

		assert.Error(t, err)
		assert.Equal(t, errors.ErrPaymentNotFound, err)
	})

	t.Run("Type TV, Inserted", func(t *testing.T) {
		mockUserRepo, _, mockEpisodeRepo, mockPaymentRepo, mockCollectionRepo, collectionService := initMock()

		refID := uint(123)
		userID := uint(10)

		mockUserRepo.On("FindUserByUsername", mock.Anything, mock.Anything).
			Return(&entity.User{
				ID:   userID,
				Role: &entity.Role{RoleCode: "GENERAL"},
			}, nil)

		mockEpisodeRepo.On("FindEpisodeByID", mock.Anything, mock.Anything).
			Return(&entity.Episode{
				ID:    refID,
				Price: sql.NullFloat64{Float64: 1.2, Valid: true},
			}, nil)

		mockPaymentRepo.On("FindPaymentByUserIDAndTypeCodeAndRefID", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
			Return(&entity.Payment{
				RefID:  refID,
				UserID: userID,
			}, nil)

		mockPaymentRepo.On("InsertPayment", mock.Anything, mock.Anything).
			Return(&entity.Payment{}, nil)

		mockCollectionRepo.On("InsertCollection", mock.Anything, mock.Anything).
			Return(nil)
		err := collectionService.AddCollection(context.Background(), &dto.CollectionDto{
			TypeCode: "TV",
			MovieID:  refID,
		})

		assert.NoError(t, err)
	})

}

func TestCollectionService_GetCollectionsByUserAndType(t *testing.T) {
	t.Run("Invalid Movie Type", func(t *testing.T) {
		_, _, _, _, _, collectionService := initMock()

		movieType := ""

		_, err := collectionService.GetCollectionsByUserAndType(context.Background(), movieType,
			mock.Anything, &pagination.PageRequest{})
		assert.Error(t, err)
		assert.Equal(t, errors.ErrInvalidInput, err)
	})

	t.Run("Invalid Client", func(t *testing.T) {
		mockUserRepo, _, _, _, _, collectionService := initMock()

		movieType := "TV"

		mockUserRepo.On("FindUserByUsername", mock.Anything, mock.Anything).
			Return(&entity.User{Role: &entity.Role{RoleCode: "BANNED"}}, nil)

		_, err := collectionService.GetCollectionsByUserAndType(context.Background(), movieType,
			mock.Anything, &pagination.PageRequest{})

		assert.Error(t, err)
		assert.Equal(t, errors.ErrInvalidClient, err)
	})

	t.Run("Success", func(t *testing.T) {
		mockUserRepo, _, _, _, mockCollectionRepo, collectionService := initMock()

		movieType := "TV"

		mockUserRepo.On("FindUserByUsername", mock.Anything, mock.Anything).
			Return(&entity.User{
				ID:   uint(10),
				Role: &entity.Role{RoleCode: "GENERAL"},
			}, nil)

		mockCollectionRepo.On("FindCollectionsByUserIDAndType", mock.Anything, mock.Anything,
			mock.Anything, mock.Anything, mock.Anything, mock.Anything).
			Return(&pagination.Page[*entity.CollectionDetail]{
				PageSize:      2,
				PageNumber:    0,
				TotalElements: 2,
				TotalPages:    1,
				Content: []*entity.CollectionDetail{
					{MovieID: uint(1), Title: "movie 1"},
					{MovieID: uint(2), Title: "movie 2"},
				},
			}, nil)

		page, err := collectionService.GetCollectionsByUserAndType(context.Background(), movieType,
			mock.Anything, &pagination.PageRequest{})

		assert.NoError(t, err)
		assert.Equal(t, 2, len(page.Content))
		assert.Equal(t, "movie 1", page.Content[0].Title)
		assert.Equal(t, uint(2), page.Content[1].MovieID)
	})
}

func TestCollectionService_GetCollectionByUserAndRefID(t *testing.T) {
	t.Run("Invalid Movie Type", func(t *testing.T) {
		_, _, _, _, _, collectionService := initMock()

		movieType := ""

		_, err := collectionService.GetCollectionByUserAndRefID(context.Background(), movieType, uint(10))
		assert.Error(t, err)
		assert.Equal(t, errors.ErrInvalidInput, err)
	})

	t.Run("Invalid Client", func(t *testing.T) {
		mockUserRepo, _, _, _, _, collectionService := initMock()

		movieType := "TV"

		mockUserRepo.On("FindUserByUsername", mock.Anything, mock.Anything).
			Return(&entity.User{Role: &entity.Role{RoleCode: "BANNED"}}, nil)

		_, err := collectionService.GetCollectionByUserAndRefID(context.Background(), movieType, uint(10))

		assert.Error(t, err)
		assert.Equal(t, errors.ErrInvalidClient, err)
	})

	t.Run("Movie Type, Success", func(t *testing.T) {
		mockUserRepo, _, _, _, mockCollectionRepo, collectionService := initMock()

		movieType := "MOVIE"
		refID := int64(50)
		userID := uint(10)

		mockUserRepo.On("FindUserByUsername", mock.Anything, mock.Anything).
			Return(&entity.User{Role: &entity.Role{RoleCode: "GENERAL"}}, nil)

		mockCollectionRepo.On("FindCollectionByUserIDAndMovieID", mock.Anything, mock.Anything, mock.Anything).
			Return(&entity.Collection{
				UserID: userID,
				MovieID: sql.NullInt64{
					Int64: refID,
					Valid: false,
				},
			}, nil)

		result, err := collectionService.GetCollectionByUserAndRefID(context.Background(), movieType, userID)

		assert.NoError(t, err)
		assert.Equal(t, userID, result.UserID)
		assert.Equal(t, uint(refID), result.MovieID)

	})

	t.Run("TV Type, Success", func(t *testing.T) {
		mockUserRepo, _, _, _, mockCollectionRepo, collectionService := initMock()

		movieType := "TV"
		refID := int64(50)
		userID := uint(10)

		mockUserRepo.On("FindUserByUsername", mock.Anything, mock.Anything).
			Return(&entity.User{Role: &entity.Role{RoleCode: "GENERAL"}}, nil)

		mockCollectionRepo.On("FindCollectionByUserIDAndEpisodeID", mock.Anything, mock.Anything, mock.Anything).
			Return(&entity.Collection{
				UserID: userID,
				EpisodeID: sql.NullInt64{
					Int64: refID,
					Valid: false,
				},
			}, nil)

		result, err := collectionService.GetCollectionByUserAndRefID(context.Background(), movieType, userID)

		assert.NoError(t, err)
		assert.Equal(t, userID, result.UserID)
		assert.Equal(t, uint(refID), result.EpisodeID)

	})
}

func TestCollectionService_RemoveCollectionByRefID(t *testing.T) {
	t.Run("Invalid Movie Type", func(t *testing.T) {
		_, _, _, _, _, collectionService := initMock()

		movieType := ""

		err := collectionService.RemoveCollectionByRefID(context.Background(), movieType, uint(10))
		assert.Error(t, err)
		assert.Equal(t, errors.ErrInvalidInput, err)
	})

	t.Run("Invalid Client", func(t *testing.T) {
		mockUserRepo, _, _, _, _, collectionService := initMock()

		movieType := "TV"

		mockUserRepo.On("FindUserByUsername", mock.Anything, mock.Anything).
			Return(&entity.User{Role: &entity.Role{RoleCode: "BANNED"}}, nil)

		err := collectionService.RemoveCollectionByRefID(context.Background(), movieType, uint(10))

		assert.Error(t, err)
		assert.Equal(t, errors.ErrInvalidClient, err)
	})

	t.Run("Movie Type, Removed", func(t *testing.T) {
		mockUserRepo, _, _, _, mockCollectionRepo, collectionService := initMock()

		movieType := "MOVIE"

		mockUserRepo.On("FindUserByUsername", mock.Anything, mock.Anything).
			Return(&entity.User{Role: &entity.Role{RoleCode: "GENERAL"}}, nil)

		mockCollectionRepo.On("DeleteCollectionByTypeCodeAndMovieID", mock.Anything, mock.Anything, mock.Anything).
			Return(nil)

		err := collectionService.RemoveCollectionByRefID(context.Background(), movieType, uint(10))

		assert.NoError(t, err)
	})

	t.Run("TV Type, Removed", func(t *testing.T) {
		mockUserRepo, _, _, _, mockCollectionRepo, collectionService := initMock()

		movieType := "TV"

		mockUserRepo.On("FindUserByUsername", mock.Anything, mock.Anything).
			Return(&entity.User{Role: &entity.Role{RoleCode: "GENERAL"}}, nil)

		mockCollectionRepo.On("DeleteCollectionByTypeCodeAndEpisodeID", mock.Anything, mock.Anything, mock.Anything).
			Return(nil)

		err := collectionService.RemoveCollectionByRefID(context.Background(), movieType, uint(10))

		assert.NoError(t, err)
	})
}

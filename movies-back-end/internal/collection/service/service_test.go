package service

import (
	"context"
	"database/sql"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"movies-service/internal/collection"
	"movies-service/internal/common/dto"
	model2 "movies-service/internal/common/model"
	"movies-service/internal/errors"
	"movies-service/internal/test/helper"
	"movies-service/pkg/pagination"
	"testing"
)

func initMock() (*helper.MockManagementCtrl, *helper.MockMovieRepository, *helper.MockEpisodeRepository, *helper.MockCollectionRepository, collection.Service) {
	mockCtrl := new(helper.MockManagementCtrl)
	mockMovieRepo := new(helper.MockMovieRepository)
	mockEpisodeRepo := new(helper.MockEpisodeRepository)
	mockCollectionRepo := new(helper.MockCollectionRepository)

	collectionService := NewCollectionService(mockMovieRepo, mockEpisodeRepo, mockCollectionRepo)

	return mockCtrl, mockMovieRepo, mockEpisodeRepo, mockCollectionRepo, collectionService
}

func TestCollectionService_AddCollection(t *testing.T) {
	t.Run("Invalid Type Code", func(t *testing.T) {
		_, _, _, _, collectionService := initMock()

		err := collectionService.AddCollection(context.Background(), &dto.CollectionDto{})
		assert.Error(t, err)
		assert.Equal(t, errors.ErrInvalidInput, err)
	})

	t.Run("Unauthorized", func(t *testing.T) {
		mockCtrl, _, _, _, collectionService := initMock()

		mockCtrl.On("CheckUser", mock.Anything).Return(false, false)

		err := collectionService.AddCollection(context.Background(), &dto.CollectionDto{
			TypeCode: "MOVIE",
		})
		assert.Error(t, err)
		assert.Equal(t, errors.ErrUnAuthorized, err)

	})

	t.Run("Type Movie", func(t *testing.T) {
		mockCtrl, mockMovieRepo, _, mockCollectionRepo, collectionService := initMock()

		mockCtrl.On("CheckUser", mock.Anything).Return(true, false)

		mockMovieRepo.On("FindMovieByID", mock.Anything, mock.Anything).
			Return(&model2.Movie{
				Price: sql.NullFloat64{Float64: 1.2, Valid: true},
			}, nil)

		mockCollectionRepo.On("InsertCollection", mock.Anything, mock.Anything).Return(nil)

		err := collectionService.AddCollection(context.Background(), &dto.CollectionDto{
			TypeCode: "MOVIE",
			MovieID:  uint(123),
		})
		assert.NoError(t, err)

	})

}

func TestCollectionService_GetCollectionByUsernameAndMovieID(t *testing.T) {
	t.Run("Invalid Movie Type", func(t *testing.T) {
		_, _, _, _, collectionService := initMock()

		movieType := ""

		_, err := collectionService.GetCollectionsByUserAndType(context.Background(), movieType,
			mock.Anything, &pagination.PageRequest{})
		assert.Error(t, err)
		assert.Equal(t, errors.ErrInvalidInput, err)
	})

	t.Run("Unauthorized", func(t *testing.T) {
		mockCtrl, _, _, _, collectionService := initMock()

		movieType := "TV"

		mockCtrl.On("CheckUser", mock.Anything).Return(false, false)

		_, err := collectionService.GetCollectionsByUserAndType(context.Background(), movieType,
			mock.Anything, &pagination.PageRequest{})
		assert.Error(t, err)
		assert.Equal(t, errors.ErrUnAuthorized, err)
	})

	t.Run("a", func(t *testing.T) {
		mockCtrl, _, _, mockRepo, collectionService := initMock()

		movieType := "TV"

		mockCtrl.On("CheckUser", mock.Anything).Return(true, false)

		mockRepo.On("FindCollectionsByUserIDAndType", mock.Anything, mock.Anything,
			mock.Anything, mock.Anything, mock.Anything, mock.Anything).
			Return(&pagination.Page[*model2.CollectionDetail]{
				PageSize:      2,
				PageNumber:    0,
				TotalElements: 2,
				TotalPages:    1,
				Content: []*model2.CollectionDetail{
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

func TestCollectionService_GetCollectionsByUsernameAndType(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		_, _, _, mockRepo, collectionService := initMock()

		mockRepo.On("FindCollectionByUserIDAndMovieID", mock.Anything, mock.Anything, mock.Anything).
			Return(&model2.Collection{
				UserID:  "testuser",
				MovieID: sql.NullInt64{},
			}, nil)

		result, err := collectionService.GetCollectionByUserAndRefID(context.Background(), "testuser", "", uint(1))

		assert.NoError(t, err)
		assert.Equal(t, uint(1), result.MovieID)
		assert.Equal(t, "testuser", result.UserID)

	})
}

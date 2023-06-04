package service

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"movies-service/internal/collection"
	"movies-service/internal/errors"
	"movies-service/internal/model"
	"movies-service/internal/test/helper"
	"movies-service/pkg/pagination"
	"testing"
)

func initMock() (*helper.MockManagementCtrl, *helper.MockCollectionRepository, collection.Service) {
	mockCtrl := new(helper.MockManagementCtrl)
	mockRepo := new(helper.MockCollectionRepository)

	collectionService := NewCollectionService(mockCtrl, mockRepo)

	return mockCtrl, mockRepo, collectionService
}

func TestCollectionService_AddCollection(t *testing.T) {
	t.Run("Invalid Movie ID", func(t *testing.T) {
		_, _, collectionService := initMock()

		movieID := uint(0)

		err := collectionService.AddCollection(context.Background(), movieID)
		assert.Error(t, err)
		assert.Equal(t, errors.ErrInvalidInput, err)
	})

	t.Run("Unauthorized", func(t *testing.T) {
		mockCtrl, _, collectionService := initMock()

		movieID := uint(123)

		mockCtrl.On("CheckUser", mock.Anything).Return(false, false)

		err := collectionService.AddCollection(context.Background(), movieID)
		assert.Error(t, err)
		assert.Equal(t, errors.ErrUnAuthorized, err)

	})

	t.Run("Valid Data", func(t *testing.T) {
		mockCtrl, mockRepo, collectionService := initMock()

		movieID := uint(123)

		mockCtrl.On("CheckUser", mock.Anything).Return(true, false)

		mockRepo.On("InsertCollection", mock.Anything, mock.Anything).Return(nil)

		err := collectionService.AddCollection(context.Background(), movieID)
		assert.NoError(t, err)

	})

}

func TestCollectionService_GetCollectionByUsernameAndMovieID(t *testing.T) {
	t.Run("Invalid Movie Type", func(t *testing.T) {
		_, _, collectionService := initMock()

		movieType := ""

		_, err := collectionService.GetCollectionsByUsernameAndType(context.Background(), movieType,
			mock.Anything, &pagination.PageRequest{})
		assert.Error(t, err)
		assert.Equal(t, errors.ErrInvalidInput, err)
	})

	t.Run("Unauthorized", func(t *testing.T) {
		mockCtrl, _, collectionService := initMock()

		movieType := "TV"

		mockCtrl.On("CheckUser", mock.Anything).Return(false, false)

		_, err := collectionService.GetCollectionsByUsernameAndType(context.Background(), movieType,
			mock.Anything, &pagination.PageRequest{})
		assert.Error(t, err)
		assert.Equal(t, errors.ErrUnAuthorized, err)
	})

	t.Run("a", func(t *testing.T) {
		mockCtrl, mockRepo, collectionService := initMock()

		movieType := "TV"

		mockCtrl.On("CheckUser", mock.Anything).Return(true, false)

		mockRepo.On("FindCollectionsByUsernameAndType", mock.Anything, mock.Anything,
			mock.Anything, mock.Anything, mock.Anything, mock.Anything).
			Return(&pagination.Page[*model.CollectionDetail]{
				PageSize:      2,
				PageNumber:    0,
				TotalElements: 2,
				TotalPages:    1,
				Content: []*model.CollectionDetail{
					{MovieID: uint(1), Title: "movie 1"},
					{MovieID: uint(2), Title: "movie 2"},
				},
			}, nil)

		page, err := collectionService.GetCollectionsByUsernameAndType(context.Background(), movieType,
			mock.Anything, &pagination.PageRequest{})

		assert.NoError(t, err)
		assert.Equal(t, 2, len(page.Content))
		assert.Equal(t, "movie 1", page.Content[0].Title)
		assert.Equal(t, uint(2), page.Content[1].MovieID)
	})
}

func TestCollectionService_GetCollectionsByUsernameAndType(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		_, mockRepo, collectionService := initMock()

		mockRepo.On("FindCollectionByUsernameAndMovieID", mock.Anything, mock.Anything, mock.Anything).
			Return(&model.Collection{
				Username: "testuser",
				MovieID:  uint(1),
			}, nil)

		result, err := collectionService.GetCollectionByUsernameAndRefID(context.Background(), "testuser", "", uint(1))

		assert.NoError(t, err)
		assert.Equal(t, uint(1), result.MovieID)
		assert.Equal(t, "testuser", result.Username)

	})
}

package service

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"movies-service/internal/common/dto"
	"movies-service/internal/errors"
	"movies-service/internal/test/helper"
	"movies-service/internal/view"
	"testing"
)

func initMock() (*helper.MockManagementCtrl, *helper.MockViewRepository, view.Service) {
	mockCtrl := new(helper.MockManagementCtrl)
	mockRepo := new(helper.MockViewRepository)

	viewService := NewViewService(mockCtrl, mockRepo)

	return mockCtrl, mockRepo, viewService
}

func TestViewService_RecognizeViewForMovie(t *testing.T) {
	t.Run("Valid Viewer", func(t *testing.T) {
		mockCtrl, mockRepo, viewService := initMock()

		ctx := context.Background()
		viewer := &dto.Viewer{
			Viewer: "user123",
		}

		mockCtrl.On("CheckUser", mock.Anything).Return(true, false)
		mockRepo.On("InsertView", ctx, viewer).Return(nil)

		err := viewService.RecognizeViewForMovie(ctx, viewer)

		// Assert
		assert.NoError(t, err)
		mockCtrl.AssertExpectations(t)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Invalid Viewer", func(t *testing.T) {
		mockCtrl, mockRepo, viewService := initMock()

		ctx := context.Background()
		viewer := &dto.Viewer{
			Viewer: "",
		}

		err := viewService.RecognizeViewForMovie(ctx, viewer)

		// Assert
		assert.Error(t, err)
		assert.Equal(t, errors.ErrInvalidClient.Error(), err.Error())
		mockCtrl.AssertExpectations(t)
		mockRepo.AssertNotCalled(t, "InsertView", ctx, viewer)
	})

	t.Run("Invalid User", func(t *testing.T) {
		mockCtrl, mockRepo, viewService := initMock()

		ctx := context.Background()
		viewer := &dto.Viewer{
			Viewer: "user123",
		}

		mockCtrl.On("CheckUser", mock.Anything).Return(false, false)

		err := viewService.RecognizeViewForMovie(ctx, viewer)

		// Assert
		assert.Error(t, err)
		assert.Equal(t, errors.ErrInvalidClient.Error(), err.Error())
		mockCtrl.AssertExpectations(t)
		mockRepo.AssertNotCalled(t, "InsertView", ctx, viewer)
	})
}

func TestViewService_GetNumberOfViewsByMovieId(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		_, mockRepo, viewService := initMock()

		ctx := context.Background()
		movieID := uint(123)
		totalViews := int64(42)

		mockRepo.On("CountViewsByMovieId", ctx, movieID).Return(totalViews, nil)

		// Act
		result, err := viewService.GetNumberOfViewsByMovieId(ctx, uint(movieID))

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, totalViews, result)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Error", func(t *testing.T) {
		_, mockRepo, viewService := initMock()

		ctx := context.Background()
		movieID := uint(123)
		expectedError := errors.ErrResourceNotFound

		mockRepo.On("CountViewsByMovieId", ctx, movieID).Return(int64(0), expectedError)

		// Act
		result, err := viewService.GetNumberOfViewsByMovieId(ctx, uint(movieID))

		// Assert
		assert.EqualError(t, err, expectedError.Error())
		assert.Zero(t, result)
		mockRepo.AssertExpectations(t)
	})

}

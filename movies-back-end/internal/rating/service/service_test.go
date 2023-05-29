package service

import (
	"context"
	"github.com/stretchr/testify/assert"
	"movies-service/internal/model"
	"movies-service/internal/rating"
	"movies-service/internal/test/helper"
	"testing"
)

func initMock() (*helper.MockRatingRepository, rating.Service) {
	// Create a mock rating repository
	mockRepo := new(helper.MockRatingRepository)

	// Create a genre service instance with the mock repository and controller
	ratingService := NewRatingService(mockRepo)

	return mockRepo, ratingService
}

func TestGetAllRatings(t *testing.T) {
	mockRepo, ratingService := initMock()

	data := []*model.Rating{
		{
			ID:   1,
			Code: "R",
			Name: "R",
		},
		{
			ID:   2,
			Code: "18A",
			Name: "18A",
		},
	}

	mockRepo.On("FindRatings", context.Background()).
		Return(data, nil).Once()

	results, err := ratingService.GetAllRatings(context.Background())

	assert.NoError(t, err)
	assert.Equal(t, 2, len(results))
	assert.Equal(t, "R", results[0].Code)
	assert.Equal(t, "18A", results[1].Code)

}

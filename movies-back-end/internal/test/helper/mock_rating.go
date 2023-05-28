package helper

import (
	"context"
	"github.com/stretchr/testify/mock"
	"movies-service/internal/model"
)

// MockRatingRepository Mock genreRepository for testing
type MockRatingRepository struct {
	mock.Mock
}

func (m *MockRatingRepository) FindRatings(ctx context.Context) ([]*model.Rating, error) {
	args := m.Called(ctx)
	return args.Get(0).([]*model.Rating), args.Error(1)
}

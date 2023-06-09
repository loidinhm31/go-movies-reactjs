package helper

import (
	"context"
	"github.com/stretchr/testify/mock"
	"movies-service/internal/common/entity"
)

// MockRatingRepository Mock genreRepository for testing
type MockRatingRepository struct {
	mock.Mock
}

func (m *MockRatingRepository) FindRatings(ctx context.Context) ([]*entity.Rating, error) {
	args := m.Called(ctx)
	return args.Get(0).([]*entity.Rating), args.Error(1)
}

package helper

import (
	"context"
	"github.com/stretchr/testify/mock"
	"movies-service/internal/model"
	"movies-service/pkg/pagination"
)

// MockSearchRepository Mock implementation of the search.Repository interface
type MockSearchRepository struct {
	mock.Mock
}

func (m *MockSearchRepository) SearchMovie(ctx context.Context, searchParams *model.SearchParams) (*pagination.Page[*model.Movie], error) {
	args := m.Called(ctx, searchParams)
	return args.Get(0).(*pagination.Page[*model.Movie]), args.Error(1)
}

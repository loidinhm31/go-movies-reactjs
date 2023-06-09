package helper

import (
	"context"
	"github.com/stretchr/testify/mock"
	"movies-service/internal/common/entity"
	model2 "movies-service/internal/common/model"
	"movies-service/pkg/pagination"
)

// MockSearchRepository Mock implementation of the search.Repository interface
type MockSearchRepository struct {
	mock.Mock
}

func (m *MockSearchRepository) SearchMovie(ctx context.Context, searchParams *model2.SearchParams) (*pagination.Page[*entity.Movie], error) {
	args := m.Called(ctx, searchParams)
	return args.Get(0).(*pagination.Page[*entity.Movie]), args.Error(1)
}

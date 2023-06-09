package helper

import (
	"context"
	"github.com/stretchr/testify/mock"
	"movies-service/internal/common/entity"
)

// MockGenreRepository Mock genreRepository for testing
type MockGenreRepository struct {
	mock.Mock
}

func (m *MockGenreRepository) FindAllGenresByTypeCode(ctx context.Context, movieType string) ([]*entity.Genre, error) {
	args := m.Called(ctx, movieType)
	return args.Get(0).([]*entity.Genre), args.Error(1)
}

func (m *MockGenreRepository) FindGenreByNameAndTypeCode(ctx context.Context, genre *entity.Genre) (*entity.Genre, error) {
	args := m.Called(ctx, genre)
	return args.Get(0).(*entity.Genre), args.Error(1)
}

func (m *MockGenreRepository) FindAllGenres(ctx context.Context) ([]*entity.Genre, error) {
	args := m.Called(ctx)
	return args.Get(0).([]*entity.Genre), args.Error(1)
}

func (m *MockGenreRepository) InsertGenres(ctx context.Context, genres []*entity.Genre) error {
	args := m.Called(ctx, genres)
	return args.Error(0)
}

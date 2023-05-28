package helper

import (
	"context"
	"github.com/stretchr/testify/mock"
	"movies-service/internal/model"
)

// MockGenreRepository Mock genreRepository for testing
type MockGenreRepository struct {
	mock.Mock
}

func (m *MockGenreRepository) FindAllGenresByTypeCode(ctx context.Context, movieType string) ([]*model.Genre, error) {
	args := m.Called(ctx, movieType)
	return args.Get(0).([]*model.Genre), args.Error(1)
}

func (m *MockGenreRepository) FindGenreByNameAndTypeCode(ctx context.Context, genre *model.Genre) (*model.Genre, error) {
	args := m.Called(ctx, genre)
	return args.Get(0).(*model.Genre), args.Error(1)
}

func (m *MockGenreRepository) FindAllGenres(ctx context.Context) ([]*model.Genre, error) {
	args := m.Called(ctx)
	return args.Get(0).([]*model.Genre), args.Error(1)
}

func (m *MockGenreRepository) InsertGenres(ctx context.Context, genres []*model.Genre) error {
	args := m.Called(ctx, genres)
	return args.Error(0)
}

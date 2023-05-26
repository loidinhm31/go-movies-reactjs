package test_helper

import (
	"context"
	"github.com/stretchr/testify/mock"
	"movies-service/internal/models"
	"movies-service/pkg/pagination"
)

// MockMovieRepository Mock movieRepository for testing
type MockMovieRepository struct {
	mock.Mock
}

func (m *MockMovieRepository) InsertMovie(ctx context.Context, movie *models.Movie) error {
	args := m.Called(ctx, movie)
	return args.Error(0)
}

func (m *MockMovieRepository) FindAllMovies(ctx context.Context,
	pageRequest *pagination.PageRequest,
	page *pagination.Page[*models.Movie]) (*pagination.Page[*models.Movie], error) {
	args := m.Called(ctx, pageRequest, page)
	return args.Get(0).(*pagination.Page[*models.Movie]), args.Error(1)
}

func (m *MockMovieRepository) FindAllMoviesByType(ctx context.Context,
	movieType string,
	pageRequest *pagination.PageRequest,
	page *pagination.Page[*models.Movie]) (*pagination.Page[*models.Movie], error) {
	args := m.Called(ctx, movieType, pageRequest, page)
	return args.Get(0).(*pagination.Page[*models.Movie]), args.Error(1)
}

func (m *MockMovieRepository) FindMovieById(ctx context.Context, id int) (*models.Movie, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(*models.Movie), args.Error(1)
}

func (m *MockMovieRepository) FindMoviesByGenre(ctx context.Context,
	pageRequest *pagination.PageRequest,
	page *pagination.Page[*models.Movie],
	genreId int) (*pagination.Page[*models.Movie], error) {
	args := m.Called(ctx, pageRequest, page, genreId)
	return args.Get(0).(*pagination.Page[*models.Movie]), args.Error(1)
}

func (m *MockMovieRepository) UpdateMovie(ctx context.Context, movie *models.Movie) error {
	args := m.Called(ctx, movie)
	return args.Error(0)
}

func (m *MockMovieRepository) UpdateMovieGenres(ctx context.Context, movie *models.Movie, genres []*models.Genre) error {
	args := m.Called(ctx, movie, genres)
	return args.Error(0)
}

func (m *MockMovieRepository) DeleteMovieById(ctx context.Context, id int) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

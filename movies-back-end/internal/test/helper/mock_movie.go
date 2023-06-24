package helper

import (
	"context"
	"github.com/stretchr/testify/mock"
	"movies-service/internal/common/entity"
	"movies-service/pkg/pagination"
)

// MockMovieRepository Mock movieRepository for testing
type MockMovieRepository struct {
	mock.Mock
}

func (m *MockMovieRepository) InsertMovie(ctx context.Context, movie *entity.Movie) error {
	args := m.Called(ctx, movie)
	return args.Error(0)
}

func (m *MockMovieRepository) FindAllMovies(ctx context.Context, keyword string, pageRequest *pagination.PageRequest, page *pagination.Page[*entity.Movie]) (*pagination.Page[*entity.Movie], error) {
	args := m.Called(ctx, keyword, pageRequest, page)
	return args.Get(0).(*pagination.Page[*entity.Movie]), args.Error(1)
}

func (m *MockMovieRepository) FindAllMoviesByType(ctx context.Context, keyword, movieType string, pageRequest *pagination.PageRequest, page *pagination.Page[*entity.Movie]) (*pagination.Page[*entity.Movie], error) {
	args := m.Called(ctx, keyword, movieType, pageRequest, page)
	return args.Get(0).(*pagination.Page[*entity.Movie]), args.Error(1)
}

func (m *MockMovieRepository) FindMovieByID(ctx context.Context, id uint) (*entity.Movie, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(*entity.Movie), args.Error(1)
}

func (m *MockMovieRepository) FindMoviesByGenre(ctx context.Context,
	pageRequest *pagination.PageRequest,
	page *pagination.Page[*entity.Movie],
	genreId uint) (*pagination.Page[*entity.Movie], error) {
	args := m.Called(ctx, pageRequest, page, genreId)
	return args.Get(0).(*pagination.Page[*entity.Movie]), args.Error(1)
}

func (m *MockMovieRepository) UpdateMovie(ctx context.Context, movie *entity.Movie) error {
	args := m.Called(ctx, movie)
	return args.Error(0)
}

func (m *MockMovieRepository) UpdateMovieGenres(ctx context.Context, movie *entity.Movie, genres []*entity.Genre) error {
	args := m.Called(ctx, movie, genres)
	return args.Error(0)
}

func (m *MockMovieRepository) DeleteMovieByID(ctx context.Context, id uint) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func (m *MockMovieRepository) FindMovieByEpisodeID(ctx context.Context, episdoeID uint) (*entity.Movie, error) {
	args := m.Called(ctx, episdoeID)
	return args.Get(0).(*entity.Movie), args.Error(1)
}

func (m *MockMovieRepository) UpdatePriceWithAverageEpisodePrice(ctx context.Context, movieID uint) error {
	args := m.Called(ctx, movieID)
	return args.Error(0)
}

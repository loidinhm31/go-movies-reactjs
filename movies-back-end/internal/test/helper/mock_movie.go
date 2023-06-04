package helper

import (
	"context"
	"github.com/stretchr/testify/mock"
	"movies-service/internal/model"
	"movies-service/pkg/pagination"
)

// MockMovieRepository Mock movieRepository for testing
type MockMovieRepository struct {
	mock.Mock
}

func (m *MockMovieRepository) FindMovieByID(ctx context.Context, id uint) (*model.Movie, error) {
	//TODO implement me
	panic("implement me")
}

func (m *MockMovieRepository) DeleteMovieByID(ctx context.Context, id uint) error {
	//TODO implement me
	panic("implement me")
}

func (m *MockMovieRepository) FindMovieByEpisodeID(ctx context.Context, episdoeID uint) (*model.Movie, error) {
	//TODO implement me
	panic("implement me")
}

func (m *MockMovieRepository) InsertMovie(ctx context.Context, movie *model.Movie) error {
	args := m.Called(ctx, movie)
	return args.Error(0)
}

func (m *MockMovieRepository) FindAllMovies(ctx context.Context, keyword string, pageRequest *pagination.PageRequest, page *pagination.Page[*model.Movie]) (*pagination.Page[*model.Movie], error) {
	args := m.Called(ctx, keyword, pageRequest, page)
	return args.Get(0).(*pagination.Page[*model.Movie]), args.Error(1)
}

func (m *MockMovieRepository) FindAllMoviesByType(ctx context.Context, keyword, movieType string, pageRequest *pagination.PageRequest, page *pagination.Page[*model.Movie]) (*pagination.Page[*model.Movie], error) {
	args := m.Called(ctx, keyword, movieType, pageRequest, page)
	return args.Get(0).(*pagination.Page[*model.Movie]), args.Error(1)
}

func (m *MockMovieRepository) FindMovieById(ctx context.Context, id uint) (*model.Movie, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(*model.Movie), args.Error(1)
}

func (m *MockMovieRepository) FindMoviesByGenre(ctx context.Context,
	pageRequest *pagination.PageRequest,
	page *pagination.Page[*model.Movie],
	genreId uint) (*pagination.Page[*model.Movie], error) {
	args := m.Called(ctx, pageRequest, page, genreId)
	return args.Get(0).(*pagination.Page[*model.Movie]), args.Error(1)
}

func (m *MockMovieRepository) UpdateMovie(ctx context.Context, movie *model.Movie) error {
	args := m.Called(ctx, movie)
	return args.Error(0)
}

func (m *MockMovieRepository) UpdateMovieGenres(ctx context.Context, movie *model.Movie, genres []*model.Genre) error {
	args := m.Called(ctx, movie, genres)
	return args.Error(0)
}

func (m *MockMovieRepository) DeleteMovieById(ctx context.Context, id uint) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

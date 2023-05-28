package helper

import (
	"context"
	"github.com/stretchr/testify/mock"
	"movies-service/internal/dto"
	"movies-service/internal/model"
)

type MockAnalysisRepository struct {
	mock.Mock
}

func (m *MockAnalysisRepository) CountMoviesByGenre(ctx context.Context, movieType string) ([]*model.GenreCount, error) {
	args := m.Called(ctx, movieType)
	return args.Get(0).([]*model.GenreCount), args.Error(1)
}

func (m *MockAnalysisRepository) CountMoviesByReleaseDate(ctx context.Context, year string, months []string) ([]*model.MovieCount, error) {
	args := m.Called(ctx, year, months)
	return args.Get(0).([]*model.MovieCount), args.Error(1)
}

func (m *MockAnalysisRepository) CountMoviesByCreatedDate(ctx context.Context, year string, months []string) ([]*model.MovieCount, error) {
	args := m.Called(ctx, year, months)
	return args.Get(0).([]*model.MovieCount), args.Error(1)
}

func (m *MockAnalysisRepository) CountViewsByGenreAndViewedDate(ctx context.Context, request *dto.RequestData) ([]*model.ViewCount, error) {
	args := m.Called(ctx, request)
	return args.Get(0).([]*model.ViewCount), args.Error(1)
}

func (m *MockAnalysisRepository) CountCumulativeViewsByGenreAndViewedDate(ctx context.Context, request *dto.RequestData) ([]*model.ViewCount, error) {
	args := m.Called(ctx, request)
	return args.Get(0).([]*model.ViewCount), args.Error(1)
}

func (m *MockAnalysisRepository) CountViewsByViewedDate(ctx context.Context, request *dto.RequestData) ([]*model.ViewCount, error) {
	args := m.Called(ctx, request)
	return args.Get(0).([]*model.ViewCount), args.Error(1)
}

func (m *MockAnalysisRepository) CountMoviesByGenreAndReleasedDate(ctx context.Context, request *dto.RequestData) ([]*model.MovieCount, error) {
	args := m.Called(ctx, request)
	return args.Get(0).([]*model.MovieCount), args.Error(1)
}

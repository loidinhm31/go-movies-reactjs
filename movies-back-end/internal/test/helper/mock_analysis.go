package helper

import (
	"context"
	"github.com/stretchr/testify/mock"
	"movies-service/internal/common/dto"
	"movies-service/internal/common/entity"
)

type MockAnalysisRepository struct {
	mock.Mock
}

func (m *MockAnalysisRepository) CountMoviesByGenre(ctx context.Context, movieType string) ([]*entity.GenreCount, error) {
	args := m.Called(ctx, movieType)
	return args.Get(0).([]*entity.GenreCount), args.Error(1)
}

func (m *MockAnalysisRepository) CountMoviesByReleaseDate(ctx context.Context, year string, months []string) ([]*entity.MovieCount, error) {
	args := m.Called(ctx, year, months)
	return args.Get(0).([]*entity.MovieCount), args.Error(1)
}

func (m *MockAnalysisRepository) CountMoviesByCreatedDate(ctx context.Context, year string, months []string) ([]*entity.MovieCount, error) {
	args := m.Called(ctx, year, months)
	return args.Get(0).([]*entity.MovieCount), args.Error(1)
}

func (m *MockAnalysisRepository) CountViewsByGenreAndViewedDate(ctx context.Context, request *dto.RequestData) ([]*entity.ViewCount, error) {
	args := m.Called(ctx, request)
	return args.Get(0).([]*entity.ViewCount), args.Error(1)
}

func (m *MockAnalysisRepository) CountCumulativeViewsByGenreAndViewedDate(ctx context.Context, request *dto.RequestData) ([]*entity.ViewCount, error) {
	args := m.Called(ctx, request)
	return args.Get(0).([]*entity.ViewCount), args.Error(1)
}

func (m *MockAnalysisRepository) CountViewsByViewedDate(ctx context.Context, request *dto.RequestData) ([]*entity.ViewCount, error) {
	args := m.Called(ctx, request)
	return args.Get(0).([]*entity.ViewCount), args.Error(1)
}

func (m *MockAnalysisRepository) CountMoviesByGenreAndReleasedDate(ctx context.Context, request *dto.RequestData) ([]*entity.MovieCount, error) {
	args := m.Called(ctx, request)
	return args.Get(0).([]*entity.MovieCount), args.Error(1)
}

func (m *MockAnalysisRepository) SumTotalAmountAndTotalReceivedPayment(ctx context.Context, typeCode string) (*entity.TotalPayment, error) {
	args := m.Called(ctx, typeCode)
	return args.Get(0).(*entity.TotalPayment), args.Error(1)
}

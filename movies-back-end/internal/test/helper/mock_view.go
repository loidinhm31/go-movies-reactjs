package helper

import (
	"context"
	"github.com/stretchr/testify/mock"
	"movies-service/internal/dto"
)

type MockViewRepository struct {
	mock.Mock
}

func (m *MockViewRepository) InsertView(ctx context.Context, viewer *dto.Viewer) error {
	args := m.Called(ctx, viewer)
	return args.Error(0)
}

func (m *MockViewRepository) CountViewsByMovieId(ctx context.Context, movieId int) (int64, error) {
	args := m.Called(ctx, movieId)
	return args.Get(0).(int64), args.Error(1)
}

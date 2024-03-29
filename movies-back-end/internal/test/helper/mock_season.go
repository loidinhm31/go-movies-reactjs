package helper

import (
	"context"
	"github.com/stretchr/testify/mock"
	"movies-service/internal/common/entity"
)

type MockSeasonRepository struct {
	mock.Mock
}

func (m *MockSeasonRepository) FindSeasonByID(ctx context.Context, id uint) (*entity.Season, error) {
	args := m.Called(ctx, id)
	result := args.Get(0)
	err := args.Error(1)
	if result != nil {
		return result.(*entity.Season), err
	}
	return nil, err
}

func (m *MockSeasonRepository) FindSeasonsByMovieID(ctx context.Context, movieID uint) ([]*entity.Season, error) {
	args := m.Called(ctx, movieID)
	result := args.Get(0)
	err := args.Error(1)
	if result != nil {
		return result.([]*entity.Season), err
	}
	return nil, err
}

func (m *MockSeasonRepository) InsertSeason(ctx context.Context, season *entity.Season) error {
	args := m.Called(ctx, season)
	return args.Error(0)
}

func (m *MockSeasonRepository) UpdateSeason(ctx context.Context, season *entity.Season) error {
	args := m.Called(ctx, season)
	return args.Error(0)
}

func (m *MockSeasonRepository) DeleteSeasonByID(ctx context.Context, id uint) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

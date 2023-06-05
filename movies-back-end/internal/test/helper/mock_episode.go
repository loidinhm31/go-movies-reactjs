package helper

import (
	"context"
	"github.com/stretchr/testify/mock"
	"movies-service/internal/common/model"
)

type MockEpisodeRepository struct {
	mock.Mock
}

func (m *MockEpisodeRepository) FindEpisodeByID(ctx context.Context, id uint) (*model.Episode, error) {
	args := m.Called(ctx, id)
	result := args.Get(0)
	err := args.Error(1)
	if result != nil {
		return result.(*model.Episode), err
	}
	return nil, err
}

func (m *MockEpisodeRepository) FindEpisodesBySeasonID(ctx context.Context, seasonID uint) ([]*model.Episode, error) {
	args := m.Called(ctx, seasonID)
	return args.Get(0).([]*model.Episode), args.Error(1)
}

func (m *MockEpisodeRepository) InsertEpisode(ctx context.Context, episode *model.Episode) error {
	args := m.Called(ctx, episode)
	return args.Error(0)
}

func (m *MockEpisodeRepository) UpdateEpisode(ctx context.Context, episode *model.Episode) error {
	args := m.Called(ctx, episode)
	return args.Error(0)
}

func (m *MockEpisodeRepository) DeleteEpisodeByID(ctx context.Context, id uint) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func (m *MockEpisodeRepository) DeleteEpisodeBySeasonID(ctx context.Context, seasonID uint) error {
	args := m.Called(ctx, seasonID)
	return args.Error(0)
}

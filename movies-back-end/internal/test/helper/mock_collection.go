package helper

import (
	"context"
	"github.com/stretchr/testify/mock"
	"movies-service/internal/common/model"
	"movies-service/pkg/pagination"
)

type MockCollectionRepository struct {
	mock.Mock
}

func (m *MockCollectionRepository) FindCollectionByMovieID(ctx context.Context, movieID uint) (*model.Collection, error) {
	args := m.Called(ctx, movieID)
	return args.Get(0).(*model.Collection), args.Error(1)
}

func (m *MockCollectionRepository) FindCollectionsByEpisodeID(ctx context.Context, episodeID uint) ([]*model.Collection, error) {
	args := m.Called(ctx, episodeID)
	return args.Get(0).([]*model.Collection), args.Error(1)
}

func (m *MockCollectionRepository) InsertCollection(ctx context.Context, collection *model.Collection) error {
	args := m.Called(ctx, collection)
	return args.Error(0)
}

func (m *MockCollectionRepository) FindCollectionsByUsernameAndType(ctx context.Context, username string, movieType string, keyword string, pageRequest *pagination.PageRequest, page *pagination.Page[*model.CollectionDetail]) (*pagination.Page[*model.CollectionDetail], error) {
	args := m.Called(ctx, username, movieType, keyword, pageRequest, page)
	return args.Get(0).(*pagination.Page[*model.CollectionDetail]), args.Error(1)
}

func (m *MockCollectionRepository) FindCollectionByUserIDAndMovieID(ctx context.Context, username string, movieID uint) (*model.Collection, error) {
	args := m.Called(ctx, username, movieID)
	return args.Get(0).(*model.Collection), args.Error(1)
}

func (m *MockCollectionRepository) FindCollectionByPaymentID(ctx context.Context, paymentID uint) (*model.Collection, error) {
	args := m.Called(ctx, paymentID)
	return args.Get(0).(*model.Collection), args.Error(1)
}

func (m *MockCollectionRepository) FindCollectionsByMovieID(ctx context.Context, movieID uint) ([]*model.Collection, error) {
	args := m.Called(ctx, movieID)
	return args.Get(0).([]*model.Collection), args.Error(1)
}

func (m *MockCollectionRepository) FindCollectionByEpisodeID(ctx context.Context, episodeID uint) (*model.Collection, error) {
	args := m.Called(ctx, episodeID)
	return args.Get(0).(*model.Collection), args.Error(1)
}

func (m *MockCollectionRepository) FindCollectionsByID(ctx context.Context, id uint) (*model.Collection, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(*model.Collection), args.Error(1)
}

func (m *MockCollectionRepository) FindCollectionByUsernameAndEpisodeID(ctx context.Context, username string, episodeID uint) (*model.Collection, error) {
	args := m.Called(ctx, username, episodeID)
	return args.Get(0).(*model.Collection), args.Error(1)
}

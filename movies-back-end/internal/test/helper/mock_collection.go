package helper

import (
	"context"
	"github.com/stretchr/testify/mock"
	"movies-service/internal/common/entity"
	"movies-service/pkg/pagination"
)

type MockCollectionRepository struct {
	mock.Mock
}

func (m *MockCollectionRepository) InsertCollection(ctx context.Context, collection *entity.Collection) error {
	args := m.Called(ctx, collection)
	return args.Error(0)
}

func (m *MockCollectionRepository) FindCollectionsByUsernameAndType(ctx context.Context, username string, movieType string, keyword string, pageRequest *pagination.PageRequest, page *pagination.Page[*entity.CollectionDetail]) (*pagination.Page[*entity.CollectionDetail], error) {
	args := m.Called(ctx, username, movieType, keyword, pageRequest, page)
	return args.Get(0).(*pagination.Page[*entity.CollectionDetail]), args.Error(1)
}

func (m *MockCollectionRepository) FindCollectionByUserIDAndMovieID(ctx context.Context, userID uint, movieID uint) (*entity.Collection, error) {
	args := m.Called(ctx, userID, movieID)
	return args.Get(0).(*entity.Collection), args.Error(1)
}

func (m *MockCollectionRepository) FindCollectionByPaymentID(ctx context.Context, paymentID uint) (*entity.Collection, error) {
	args := m.Called(ctx, paymentID)
	return args.Get(0).(*entity.Collection), args.Error(1)
}

func (m *MockCollectionRepository) FindCollectionsByMovieID(ctx context.Context, movieID uint) ([]*entity.Collection, error) {
	args := m.Called(ctx, movieID)
	return args.Get(0).([]*entity.Collection), args.Error(1)
}

func (m *MockCollectionRepository) FindCollectionsByEpisodeID(ctx context.Context, episodeID uint) ([]*entity.Collection, error) {
	args := m.Called(ctx, episodeID)
	return args.Get(0).([]*entity.Collection), args.Error(1)
}

func (m *MockCollectionRepository) FindCollectionByUsernameAndEpisodeID(ctx context.Context, username string, episodeID uint) (*entity.Collection, error) {
	args := m.Called(ctx, username, episodeID)
	return args.Get(0).(*entity.Collection), args.Error(1)
}

func (m *MockCollectionRepository) FindCollectionsByUserIDAndType(ctx context.Context, userID uint, movieType string, keyword string, pageRequest *pagination.PageRequest, page *pagination.Page[*entity.CollectionDetail]) (*pagination.Page[*entity.CollectionDetail], error) {
	args := m.Called(ctx, userID, movieType, keyword, pageRequest, page)
	return args.Get(0).(*pagination.Page[*entity.CollectionDetail]), args.Error(1)
}

func (m *MockCollectionRepository) FindCollectionByUserIDAndEpisodeID(ctx context.Context, userID uint, episodeID uint) (*entity.Collection, error) {
	args := m.Called(ctx, userID, episodeID)
	return args.Get(0).(*entity.Collection), args.Error(1)
}

func (m *MockCollectionRepository) DeleteCollectionByTypeCodeAndMovieID(ctx context.Context, typeCode string, movieID uint) error {
	args := m.Called(ctx, typeCode, movieID)
	return args.Error(0)
}

func (m *MockCollectionRepository) DeleteCollectionByTypeCodeAndEpisodeID(ctx context.Context, typeCode string, episodeID uint) error {
	args := m.Called(ctx, typeCode, episodeID)
	return args.Error(0)
}

package helper

import (
	"context"
	"github.com/stretchr/testify/mock"
	"movies-service/internal/model"
	"movies-service/pkg/pagination"
)

type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) FindUserByID(ctx context.Context, userID uint) (*model.User, error) {
	args := m.Called(ctx, userID)
	result := args.Get(0)
	if result != nil {
		return result.(*model.User), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockUserRepository) FindAllUsers(ctx context.Context, pageRequest *pagination.PageRequest, page *pagination.Page[*model.User], key string, isNew bool) (*pagination.Page[*model.User], error) {
	args := m.Called(ctx, pageRequest, page, key, isNew)
	return args.Get(0).(*pagination.Page[*model.User]), args.Error(1)
}

func (m *MockUserRepository) UpdateUserRole(ctx context.Context, userID uint, roleID uint) error {
	args := m.Called(ctx, userID, roleID)
	return args.Error(0)
}

func (m *MockUserRepository) FindUserByUsername(ctx context.Context, user *model.User) (*model.User, error) {
	args := m.Called(ctx, user)
	result := args.Get(0)
	if result != nil {
		return result.(*model.User), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockUserRepository) InsertUser(ctx context.Context, user *model.User) error {
	args := m.Called(ctx, user)
	return args.Error(0)
}

package helper

import (
	"context"
	"github.com/stretchr/testify/mock"
	"movies-service/internal/common/entity"
	"movies-service/pkg/pagination"
)

type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) FindUserByID(ctx context.Context, userID uint) (*entity.User, error) {
	args := m.Called(ctx, userID)
	result := args.Get(0)
	if result != nil {
		return result.(*entity.User), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockUserRepository) FindAllUsers(ctx context.Context, pageRequest *pagination.PageRequest, page *pagination.Page[*entity.User], key string, isNew bool) (*pagination.Page[*entity.User], error) {
	args := m.Called(ctx, pageRequest, page, key, isNew)
	return args.Get(0).(*pagination.Page[*entity.User]), args.Error(1)
}

func (m *MockUserRepository) UpdateUserRole(ctx context.Context, userID uint, roleID uint) error {
	args := m.Called(ctx, userID, roleID)
	return args.Error(0)
}

func (m *MockUserRepository) FindUserByUsername(ctx context.Context, username string) (*entity.User, error) {
	args := m.Called(ctx, username)
	result := args.Get(0)
	if result != nil {
		return result.(*entity.User), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockUserRepository) InsertUser(ctx context.Context, user *entity.User) error {
	args := m.Called(ctx, user)
	return args.Error(0)
}

func (m *MockUserRepository) FindUserByUsernameAndIsNew(ctx context.Context, username string, isNew bool) (*entity.User, error) {
	args := m.Called(ctx, username, isNew)
	result := args.Get(0)
	if result != nil {
		return result.(*entity.User), args.Error(1)
	}
	return nil, args.Error(1)
}

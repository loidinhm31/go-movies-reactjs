package helper

import (
	"context"
	"github.com/stretchr/testify/mock"
	"movies-service/internal/common/entity"
)

type MockRoleRepository struct {
	mock.Mock
}

func (m *MockRoleRepository) FindRoleByRoleCode(ctx context.Context, roleCode string) (*entity.Role, error) {
	args := m.Called(ctx, roleCode)
	result := args.Get(0)
	if result != nil {
		return result.(*entity.Role), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockRoleRepository) FindRoles(ctx context.Context) ([]*entity.Role, error) {
	args := m.Called(ctx)
	result := args.Get(0)
	if result != nil {
		return result.([]*entity.Role), args.Error(1)
	}
	return nil, args.Error(1)
}

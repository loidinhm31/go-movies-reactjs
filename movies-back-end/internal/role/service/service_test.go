package service

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"movies-service/internal/model"
	"testing"
)

type MockRoleRepository struct {
	mock.Mock
}

func (m *MockRoleRepository) FindRoleByRoleCode(ctx context.Context, roleCode string) (*model.Role, error) {
	args := m.Called(ctx, roleCode)
	result := args.Get(0)
	if result != nil {
		return result.(*model.Role), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockRoleRepository) FindRoles(ctx context.Context) ([]*model.Role, error) {
	args := m.Called(ctx)
	result := args.Get(0)
	if result != nil {
		return result.([]*model.Role), args.Error(1)
	}
	return nil, args.Error(1)
}

func TestGetAllRoles(t *testing.T) {
	mockRepo := new(MockRoleRepository)

	// Create a genre service instance with the mock repository and controller
	roleService := NewRoleService(mockRepo)

	mockRepo.On("FindRoles", context.Background()).
		Return([]*model.Role{
			{
				ID:       1,
				RoleName: "Admin",
				RoleCode: "ADMIN",
			},
			{
				ID:       2,
				RoleName: "General",
				RoleCode: "GENERAL",
			},
		}, nil)

	results, err := roleService.GetAllRoles(context.Background())

	assert.NoError(t, err)
	assert.Equal(t, 2, len(results))
	assert.Equal(t, "ADMIN", results[0].RoleCode)
	assert.Equal(t, "GENERAL", results[1].RoleCode)

}

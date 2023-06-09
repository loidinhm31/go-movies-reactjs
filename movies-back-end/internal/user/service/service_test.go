package service

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"movies-service/config"
	"movies-service/internal/common/dto"
	"movies-service/internal/common/entity"
	"movies-service/internal/errors"
	"movies-service/internal/test/helper"
	"movies-service/internal/user"
	"movies-service/pkg/pagination"
	"testing"
	"time"
)

func initMock() (*helper.MockManagementCtrl, *helper.MockRoleRepository, *helper.MockUserRepository, *helper.MockMailService, user.Service) {
	cfg := &config.Config{
		Mail: config.Mail{
			From: "from",
		},
	}

	mockCtrl := new(helper.MockManagementCtrl)
	mockRoleRepo := new(helper.MockRoleRepository)
	mockUserRepo := new(helper.MockUserRepository)
	mockMailSvc := new(helper.MockMailService)

	userService := NewUserService(cfg, mockCtrl, mockRoleRepo, mockUserRepo, mockMailSvc)

	return mockCtrl, mockRoleRepo, mockUserRepo, mockMailSvc, userService
}

func TestUserService_GetUsers(t *testing.T) {
	t.Run("Unauthorized", func(t *testing.T) {
		mockCtrl, _, _, _, userService := initMock()

		// Set up test data
		ctx := context.Background()

		key := "search key"

		// Set up mock expectations
		mockCtrl.On("CheckAdminPrivilege", mock.Anything).Return(false)

		// Call the method being tested
		result, err := userService.GetUsers(ctx, &pagination.PageRequest{}, key, true)

		// Assert the result
		assert.Error(t, err)
		assert.Equal(t, errors.ErrUnAuthorized.Error(), err.Error())
		assert.Nil(t, result)
	})

	t.Run("Error Getting User", func(t *testing.T) {
		mockCtrl, _, mockUserRepo, _, userService := initMock()

		key := "search key"

		// Set up mock expectations
		mockCtrl.On("CheckAdminPrivilege", mock.Anything).Return(true)

		mockUserRepo.On("FindAllUsers", context.Background(), &pagination.PageRequest{}, &pagination.Page[*entity.User]{}, key, true).
			Return(&pagination.Page[*entity.User]{}, errors.ErrResourceNotFound)

		// Call the method being tested
		result, err := userService.GetUsers(context.Background(), &pagination.PageRequest{}, key, true)

		// Assert the result
		assert.Error(t, err)
		assert.Equal(t, errors.ErrResourceNotFound.Error(), err.Error())
		assert.Nil(t, result)
	})

	t.Run("Success", func(t *testing.T) {
		mockCtrl, _, mockUserRepo, _, userService := initMock()

		key := "search key"

		// Set up mock expectations
		mockCtrl.On("CheckAdminPrivilege", mock.Anything).Return(true)

		mockUserRepo.On("FindAllUsers", context.Background(), &pagination.PageRequest{}, &pagination.Page[*entity.User]{}, key, true).
			Return(&pagination.Page[*entity.User]{
				PageSize:      1,
				PageNumber:    0,
				Sort:          pagination.Sort{},
				TotalElements: 2,
				TotalPages:    1,
				Content: []*entity.User{
					{ID: 1, Username: "f1", Email: "f1@example.com", FirstName: "A", LastName: "AA", IsNew: true, Role: &entity.Role{ID: 1, RoleCode: "BANNED"}},
					{ID: 2, Username: "f1", Email: "f2@example.com", FirstName: "B", LastName: "BB", IsNew: true, Role: &entity.Role{ID: 1, RoleCode: "BANNED"}},
				},
			}, nil)

		// Call the method being tested
		result, err := userService.GetUsers(context.Background(), &pagination.PageRequest{}, key, true)

		// Assert the result
		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, int64(2), result.TotalElements)
		assert.Equal(t, 1, result.TotalPages)
		assert.Len(t, result.Content, 2)
		assert.Equal(t, "f1", result.Content[0].Username)
		assert.Equal(t, "f1@example.com", result.Content[0].Email)
		assert.Equal(t, "A", result.Content[0].FirstName)
		assert.Equal(t, "AA", result.Content[0].LastName)
		assert.Equal(t, true, result.Content[0].IsNew)
		assert.Equal(t, uint(1), result.Content[0].Role.ID)
		assert.Equal(t, "BANNED", result.Content[0].Role.RoleCode)

	})
}

func TestUserService_UpdateUserRole(t *testing.T) {
	t.Run("Unauthorized", func(t *testing.T) {
		mockCtrl, _, _, _, userService := initMock()

		// Set up test data
		ctx := context.TODO()
		userDto := &dto.UserDto{
			ID: 1,
			Role: dto.RoleDto{
				RoleCode: "ADMIN",
			},
		}

		// Set up mock expectations
		mockCtrl.On("CheckAdminPrivilege", mock.AnythingOfType("string")).Return(false)

		// Call the method being tested
		err := userService.UpdateUserRole(ctx, userDto)

		// Assert the result
		assert.Error(t, err)
		assert.Equal(t, errors.ErrUnAuthorized, err)
	})

	t.Run("Error Getting User", func(t *testing.T) {
		mockCtrl, mockRoleRepo, mockUserRepo, _, userService := initMock()

		// Set up test data
		ctx := context.Background()
		userDto := &dto.UserDto{
			ID: 1,
			Role: dto.RoleDto{
				RoleCode: "ADMIN",
			},
		}

		// Set up mock expectations
		mockCtrl.On("CheckAdminPrivilege", mock.AnythingOfType("string")).Return(true)
		mockUserRepo.On("FindUserByID", ctx, userDto.ID).Return(nil, errors.ErrResourceNotFound)

		// Call the method being tested
		err := userService.UpdateUserRole(ctx, userDto)

		// Assert the result
		assert.Error(t, err)
		assert.Equal(t, errors.ErrResourceNotFound, err)

		// Assert that the expected methods were called on the mocks
		mockCtrl.AssertExpectations(t)
		mockUserRepo.AssertExpectations(t)
		mockRoleRepo.AssertExpectations(t)
	})

	t.Run("Error Getting Role", func(t *testing.T) {
		mockCtrl, mockRoleRepo, mockUserRepo, _, userService := initMock()

		// Set up test data
		ctx := context.TODO()
		userDto := &dto.UserDto{
			ID: 1,
			Role: dto.RoleDto{
				RoleCode: "ADMIN",
			},
		}

		// Set up mock expectations
		mockCtrl.On("CheckAdminPrivilege", mock.Anything).Return(true)
		mockUserRepo.On("FindUserByID", ctx, userDto.ID).Return(&entity.User{}, nil)
		mockRoleRepo.On("FindRoleByRoleCode", ctx, userDto.Role.RoleCode).Return(nil, errors.ErrResourceNotFound)

		// Call the method being tested
		err := userService.UpdateUserRole(ctx, userDto)

		// Assert the result
		assert.Error(t, err)
		assert.Equal(t, errors.ErrResourceNotFound, err)

		// Assert that the expected methods were called on the mocks
		mockCtrl.AssertExpectations(t)
		mockUserRepo.AssertExpectations(t)
		mockRoleRepo.AssertExpectations(t)
	})

	t.Run("Success Update User Role", func(t *testing.T) {
		mockCtrl, mockRoleRepo, mockUserRepo, mockMailSvc, userService := initMock()

		// Set up test data
		ctx := context.Background()
		userDto := &dto.UserDto{
			ID: 1,
			Role: dto.RoleDto{
				RoleCode: "ADMIN",
			},
		}

		// Set up mock expectations
		expectedUser := &entity.User{
			ID:        1,
			Username:  "john",
			Email:     "john@example.com",
			FirstName: "John",
			LastName:  "Doe",
			IsNew:     true,
			Role: &entity.Role{
				ID:       1,
				RoleName: "Admin",
				RoleCode: "ADMIN",
			},
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}
		mockCtrl.On("CheckAdminPrivilege", mock.Anything).Return(true)
		mockUserRepo.On("FindUserByID", ctx, userDto.ID).Return(expectedUser, nil)
		mockRoleRepo.On("FindRoleByRoleCode", ctx, userDto.Role.RoleCode).Return(expectedUser.Role, nil)
		mockUserRepo.On("UpdateUserRole", ctx, expectedUser.ID, expectedUser.Role.ID).Return(nil)
		mockMailSvc.On("SendMessage", ctx, mock.Anything).
			Return(nil)

		// Call the method being tested
		err := userService.UpdateUserRole(ctx, userDto)

		// Assert the result
		assert.NoError(t, err)

		// Assert that the expected methods were called on the mocks
		mockCtrl.AssertExpectations(t)
		mockUserRepo.AssertExpectations(t)
		mockRoleRepo.AssertExpectations(t)
	})
}

func TestUserService_AddOidcUser(t *testing.T) {
	t.Run("User exists", func(t *testing.T) {
		_, _, mockUserRepo, _, userService := initMock()

		mockUserRepo.On("FindUserByUsernameAndIsNew", context.Background(), mock.Anything, false).
			Return(&entity.User{}, nil)

		// Call the method being tested
		_, err := userService.AddOidcUser(context.Background(), &dto.UserDto{})

		// Assert the result
		assert.Error(t, err)
		assert.Equal(t, errors.ErrUserExisted.Error(), err.Error())
	})

	t.Run("Unauthorized", func(t *testing.T) {
		mockCtrl, _, mockUserRepo, _, userService := initMock()

		mockUserRepo.On("FindUserByUsernameAndIsNew", context.Background(), mock.Anything, false).
			Return(nil, nil)

		mockCtrl.On("CheckAdminPrivilege", mock.Anything).Return(false)

		// Call the method being tested
		_, err := userService.AddOidcUser(context.Background(), &dto.UserDto{})

		// Assert the result
		assert.Error(t, err)
		assert.Equal(t, errors.ErrUnAuthorized.Error(), err.Error())
	})

	t.Run("Success Add", func(t *testing.T) {
		mockCtrl, mockRoleRepo, mockUserRepo, _, userService := initMock()

		// Set up test data
		ctx := context.Background()
		userDto := &dto.UserDto{
			Username:  "john",
			Email:     "john@example.com",
			FirstName: "John",
			LastName:  "Doe",
			Role: dto.RoleDto{
				RoleCode: "ADMIN",
			},
		}

		// Set up mock expectations
		expectedUser := &entity.User{
			Username:  userDto.Username,
			Email:     userDto.Email,
			FirstName: userDto.FirstName,
			LastName:  userDto.LastName,
			Role: &entity.Role{
				ID:       1,
				RoleName: "Admin",
				RoleCode: "ADMIN",
			},
			IsNew:     false,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}
		mockUserRepo.On("FindUserByUsernameAndIsNew", ctx, mock.Anything, false).
			Return(nil, nil)

		mockCtrl.On("CheckAdminPrivilege", mock.Anything).Return(true)

		mockRoleRepo.On("FindRoleByRoleCode", ctx, userDto.Role.RoleCode).Return(expectedUser.Role, nil)

		mockUserRepo.On("InsertUser", ctx, mock.Anything).
			Return(nil)

		// Call the method being tested
		result, err := userService.AddOidcUser(ctx, userDto)

		// Assert the result
		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, userDto.Username, result.Username)
		assert.Equal(t, userDto.Email, result.Email)
		assert.Equal(t, userDto.FirstName, result.FirstName)
		assert.Equal(t, userDto.LastName, result.LastName)
		assert.Equal(t, userDto.Role.RoleCode, result.Role.RoleCode)
		assert.False(t, result.IsNew)

		// Assert that the expected methods were called on the mocks
		mockUserRepo.AssertExpectations(t)
		mockCtrl.AssertExpectations(t)
		mockRoleRepo.AssertExpectations(t)
	})
}

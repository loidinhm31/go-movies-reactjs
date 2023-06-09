package service

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"movies-service/internal/common/entity"
	"movies-service/internal/control"
	"movies-service/internal/test/helper"
	"testing"
)

func initMock() (*helper.MockUserRepository, control.Service) {
	mockRepo := new(helper.MockUserRepository)
	controlService := NewManagementCtrl(mockRepo)
	return mockRepo, controlService
}

func TestManagementCtrl_CheckPrivilege(t *testing.T) {
	t.Run("True", func(t *testing.T) {
		mockRepo, controlService := initMock()

		expectedUser := &entity.User{
			Username: "testuser",
			IsNew:    false,
			Role: &entity.Role{
				RoleCode: "ADMIN",
			},
		}

		mockRepo.On("FindUserByUsernameAndIsNew", mock.Anything, mock.Anything, mock.Anything).
			Return(expectedUser, nil)

		result := controlService.CheckPrivilege("testuser")

		// Assert
		assert.True(t, result)
	})

	t.Run("False", func(t *testing.T) {
		mockRepo, controlService := initMock()

		expectedUser := &entity.User{
			Username: "user",
			IsNew:    false,
			Role: &entity.Role{
				RoleCode: "GENERAL",
			},
		}

		mockRepo.On("FindUserByUsernameAndIsNew", mock.Anything, mock.Anything, mock.Anything).
			Return(expectedUser, nil)

		result := controlService.CheckPrivilege("user")

		// Assert
		assert.False(t, result)
	})

	t.Run("Error Repo", func(t *testing.T) {
		mockRepo, controlService := initMock()

		expectedError := errors.New("repository error")

		mockRepo.On("FindUserByUsernameAndIsNew", mock.Anything, mock.Anything, mock.Anything).
			Return(nil, expectedError)

		result := controlService.CheckPrivilege("testuser")

		// Assert
		assert.False(t, result)
	})
}

func TestManagementCtrl_CheckAdminPrivilege(t *testing.T) {
	t.Run("True", func(t *testing.T) {
		mockRepo, controlService := initMock()

		expectedUser := &entity.User{
			Username: "testuser",
			IsNew:    false,
			Role: &entity.Role{
				RoleCode: "ADMIN",
			},
		}
		mockRepo.On("FindUserByUsernameAndIsNew", mock.Anything, mock.Anything, mock.Anything).
			Return(expectedUser, nil)

		result := controlService.CheckAdminPrivilege("testuser")

		// Assert
		assert.True(t, result)

	})

	t.Run("False", func(t *testing.T) {
		mockRepo, controlService := initMock()

		expectedUser := &entity.User{
			Username: "user",
			IsNew:    false,
			Role: &entity.Role{
				RoleCode: "GENERAL",
			},
		}
		mockRepo.On("FindUserByUsernameAndIsNew", mock.Anything, mock.Anything, mock.Anything).
			Return(expectedUser, nil)

		result := controlService.CheckAdminPrivilege("user")

		// Assert
		assert.False(t, result)

	})

	t.Run("Error Repo", func(t *testing.T) {
		mockRepo, controlService := initMock()

		expectedError := errors.New("repository error")
		mockRepo.On("FindUserByUsernameAndIsNew", mock.Anything, mock.Anything, mock.Anything).
			Return(nil, expectedError)

		result := controlService.CheckAdminPrivilege("admin")

		// Assert
		assert.False(t, result)

	})
}

func TestManagementCtrl_CheckUser(t *testing.T) {
	t.Run("True User", func(t *testing.T) {
		mockRepo, controlService := initMock()

		expectedUser := &entity.User{
			Username: "user1",
			IsNew:    false,
			Role: &entity.Role{
				RoleCode: "GENERAL",
			},
		}
		mockRepo.On("FindUserByUsernameAndIsNew", mock.Anything, mock.Anything, mock.Anything).
			Return(expectedUser, nil)

		isValidUser, isPrivilege := controlService.CheckUser("user1")

		// Assert
		assert.True(t, isValidUser)
		assert.False(t, isPrivilege)
	})

	t.Run("True Privilege", func(t *testing.T) {
		mockRepo, controlService := initMock()

		expectedUser := &entity.User{
			Username: "user1",
			IsNew:    false,
			Role: &entity.Role{
				RoleCode: "ADMIN",
			},
		}
		mockRepo.On("FindUserByUsernameAndIsNew", mock.Anything, mock.Anything, mock.Anything).
			Return(expectedUser, nil)

		isValidUser, isPrivilege := controlService.CheckUser("user1")

		// Assert
		assert.True(t, isValidUser)
		assert.True(t, isPrivilege)
	})

	t.Run("False", func(t *testing.T) {
		mockRepo, controlService := initMock()

		expectedUser := &entity.User{
			Username: "banneduser",
			IsNew:    false,
			Role: &entity.Role{
				RoleCode: "BANNED",
			},
		}
		mockRepo.On("FindUserByUsernameAndIsNew", mock.Anything, mock.Anything, mock.Anything).
			Return(expectedUser, nil)

		isValidUser, isPrivilege := controlService.CheckUser("banneduser")

		// Assert
		assert.False(t, isValidUser)
		assert.False(t, isPrivilege)
	})

	t.Run("Error Repo", func(t *testing.T) {
		mockRepo, controlService := initMock()

		expectedError := errors.New("repository error")
		mockRepo.On("FindUserByUsernameAndIsNew", mock.Anything, mock.Anything, mock.Anything).
			Return(nil, expectedError)

		isValidUser, isPrivilege := controlService.CheckUser("user1")

		// Assert
		assert.False(t, isValidUser)
		assert.False(t, isPrivilege)
	})
}

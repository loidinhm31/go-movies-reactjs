package service

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"movies-service/internal/control"
	"movies-service/internal/model"
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

		expectedUser := &model.User{
			Username: "testuser",
			IsNew:    false,
			Role: &model.Role{
				RoleCode: "ADMIN",
			},
		}
		mockRepo.On("FindUserByUsername", mock.Anything, mock.Anything).
			Return(expectedUser, nil)

		result := controlService.CheckPrivilege("testuser")

		// Assert
		assert.True(t, result)
		mockRepo.AssertCalled(t, "FindUserByUsername", mock.Anything, &model.User{Username: "testuser", IsNew: false})
	})

	t.Run("False", func(t *testing.T) {
		mockRepo, controlService := initMock()

		expectedUser := &model.User{
			Username: "user",
			IsNew:    false,
			Role: &model.Role{
				RoleCode: "GENERAL",
			},
		}
		mockRepo.On("FindUserByUsername", mock.Anything, mock.Anything).
			Return(expectedUser, nil)

		result := controlService.CheckPrivilege("user")

		// Assert
		assert.False(t, result)
		mockRepo.AssertCalled(t, "FindUserByUsername", mock.Anything, &model.User{Username: "user", IsNew: false})

	})

	t.Run("Error Repo", func(t *testing.T) {
		mockRepo, controlService := initMock()

		expectedError := errors.New("repository error")
		mockRepo.On("FindUserByUsername", mock.Anything, mock.Anything).
			Return(nil, expectedError)

		result := controlService.CheckPrivilege("testuser")

		// Assert
		assert.False(t, result)
		mockRepo.AssertCalled(t, "FindUserByUsername", mock.Anything, &model.User{Username: "testuser", IsNew: false})

	})
}

func TestManagementCtrl_CheckAdminPrivilege(t *testing.T) {
	t.Run("True", func(t *testing.T) {
		mockRepo, controlService := initMock()

		expectedUser := &model.User{
			Username: "testuser",
			IsNew:    false,
			Role: &model.Role{
				RoleCode: "ADMIN",
			},
		}
		mockRepo.On("FindUserByUsername", mock.Anything, mock.Anything).
			Return(expectedUser, nil)

		result := controlService.CheckAdminPrivilege("testuser")

		// Assert
		assert.True(t, result)
		mockRepo.AssertCalled(t, "FindUserByUsername", mock.Anything, &model.User{Username: "testuser", IsNew: false})

	})

	t.Run("False", func(t *testing.T) {
		mockRepo, controlService := initMock()

		expectedUser := &model.User{
			Username: "user",
			IsNew:    false,
			Role: &model.Role{
				RoleCode: "USER",
			},
		}
		mockRepo.On("FindUserByUsername", mock.Anything, mock.Anything).
			Return(expectedUser, nil)

		result := controlService.CheckAdminPrivilege("user")

		// Assert
		assert.False(t, result)
		mockRepo.AssertCalled(t, "FindUserByUsername", mock.Anything, &model.User{Username: "user", IsNew: false})

	})

	t.Run("Error Repo", func(t *testing.T) {
		mockRepo, controlService := initMock()

		expectedError := errors.New("repository error")
		mockRepo.On("FindUserByUsername", mock.Anything, mock.Anything).
			Return(nil, expectedError)

		result := controlService.CheckAdminPrivilege("admin")

		// Assert
		assert.False(t, result)
		mockRepo.AssertCalled(t, "FindUserByUsername", mock.Anything, &model.User{Username: "admin", IsNew: false})

	})
}

func TestManagementCtrl_CheckUser(t *testing.T) {
	t.Run("True", func(t *testing.T) {
		mockRepo, controlService := initMock()

		expectedUser := &model.User{
			Username: "user1",
			IsNew:    false,
			Role: &model.Role{
				RoleCode: "USER",
			},
		}
		mockRepo.On("FindUserByUsername", mock.Anything, mock.Anything).
			Return(expectedUser, nil)

		result := controlService.CheckUser("user1")

		// Assert
		assert.True(t, result)
		mockRepo.AssertCalled(t, "FindUserByUsername", mock.Anything, &model.User{Username: "user1", IsNew: false})
	})

	t.Run("False", func(t *testing.T) {
		mockRepo, controlService := initMock()

		expectedUser := &model.User{
			Username: "banneduser",
			IsNew:    false,
			Role: &model.Role{
				RoleCode: "BANNED",
			},
		}
		mockRepo.On("FindUserByUsername", mock.Anything, mock.Anything).
			Return(expectedUser, nil)

		result := controlService.CheckUser("banneduser")

		// Assert
		assert.False(t, result)
		mockRepo.AssertCalled(t, "FindUserByUsername", mock.Anything, &model.User{Username: "banneduser", IsNew: false})
	})

	t.Run("Error Repo", func(t *testing.T) {
		mockRepo, controlService := initMock()

		expectedError := errors.New("repository error")
		mockRepo.On("FindUserByUsername", mock.Anything, mock.Anything).
			Return(nil, expectedError)

		result := controlService.CheckUser("user1")

		// Assert
		assert.False(t, result)
		mockRepo.AssertCalled(t, "FindUserByUsername", mock.Anything, &model.User{Username: "user1", IsNew: false})
	})
}

package service

import (
	"context"
	"github.com/Nerzal/gocloak/v13"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"movies-service/config"
	"movies-service/internal/auth"
	"movies-service/internal/dto"
	"movies-service/internal/errors"
	"movies-service/internal/model"
	"movies-service/internal/test/helper"
	"testing"
	"time"
)

func initMock() (*helper.MockGoCloak, *helper.MockManagementCtrl, *helper.MockUserRepository, *helper.MockRoleRepository, auth.Service) {
	// Create mock dependencies
	mockUserRepo := new(helper.MockUserRepository)
	mockRoleRepo := new(helper.MockRoleRepository)
	mockCloak := new(helper.MockGoCloak)

	mockCtrl := new(helper.MockManagementCtrl)
	// Create the authService instance with the mock dependencies

	authService := NewAuthService(config.KeycloakConfig{}, mockCloak, mockCtrl, mockRoleRepo, mockUserRepo)
	return mockCloak, mockCtrl, mockUserRepo, mockRoleRepo, authService
}

func TestSignUp(t *testing.T) {

	t.Run("Added user", func(t *testing.T) {
		_, _, mockUserRepo, mockRoleRepo, authService := initMock()

		userDto := &dto.UserDto{
			Username:  "john",
			FirstName: "John",
			LastName:  "Doe",
			Email:     "john@example.com",
			IsNew:     true,
		}

		mockUserRepo.On("FindUserByUsername", context.Background(), mock.Anything).
			Return(&model.User{ID: 1, Username: "john", IsNew: false}, nil)

		_, err := authService.SignUp(context.Background(), userDto)

		// Assert that the userDto is returned without error
		assert.Error(t, err)
		assert.Equal(t, errors.ErrUserExisted.Error(), err.Error())

		// Assert that there are no more calls to the UserRepository and RoleRepository
		mockUserRepo.AssertExpectations(t)
		mockRoleRepo.AssertExpectations(t)
	})

	t.Run("Add User successful", func(t *testing.T) {
		_, _, mockUserRepo, mockRoleRepo, authService := initMock()

		userDto := &dto.UserDto{
			Username:  "john",
			FirstName: "John",
			LastName:  "Doe",
			Email:     "john@example.com",
			IsNew:     true,
		}

		mockUserRepo.On("FindUserByUsername", context.Background(), mock.Anything).
			Return(nil, nil)

		// Mock the RoleRepository's FindRoleByRoleCode method to return a role
		mockRoleRepo.On("FindRoleByRoleCode", context.Background(), "BANNED").
			Return(&model.Role{}, nil)

		// Mock the UserRepository's InsertUser method to succeed
		mockUserRepo.On("InsertUser", context.Background(), mock.Anything).
			Return(nil)

		// Call the SignUp function
		result, err := authService.SignUp(context.Background(), userDto)

		// Assert that the userDto is returned without error
		assert.NoError(t, err)
		assert.Equal(t, userDto, result)

		// Assert that the UserRepository's methods were called with the expected arguments
		mockUserRepo.AssertCalled(t, "FindUserByUsername", context.Background(), &model.User{Username: "john"})
		mockRoleRepo.AssertCalled(t, "FindRoleByRoleCode", context.Background(), "BANNED")
		mockUserRepo.AssertCalled(t, "InsertUser", context.Background(), mock.Anything)

		// Assert that there are no more calls to the UserRepository and RoleRepository
		mockUserRepo.AssertExpectations(t)
		mockRoleRepo.AssertExpectations(t)
	})
}

func TestSignIn(t *testing.T) {

	// Prepare test data
	username := "testuser"
	expectedUser := &model.User{
		ID:        123,
		Username:  username,
		Email:     "testuser@example.com",
		FirstName: "Test",
		LastName:  "User",
		Role: &model.Role{
			ID:       456,
			RoleName: "admin",
			RoleCode: "ADMIN",
		},
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	t.Run("Existing user", func(t *testing.T) {
		_, _, mockUserRepo, _, authService := initMock()

		// Mock the FindUserByUsername method to return an existing user
		mockUserRepo.On("FindUserByUsername", context.Background(), mock.AnythingOfType("*model.User")).
			Return(expectedUser, nil)

		// Call the SignIn method
		result, err := authService.SignIn(context.Background(), username)

		// Assert the result and error
		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, expectedUser.Username, result.Username)
		assert.Equal(t, expectedUser.Email, result.Email)
		assert.Equal(t, expectedUser.FirstName, result.FirstName)
		assert.Equal(t, expectedUser.LastName, result.LastName)
		assert.Equal(t, expectedUser.Role.ID, result.Role.ID)
		assert.Equal(t, expectedUser.Role.RoleName, result.Role.RoleName)
		assert.Equal(t, expectedUser.Role.RoleCode, result.Role.RoleCode)

		// Assert that the FindUserByUsername method was called with the expected arguments
		mockUserRepo.AssertCalled(t, "FindUserByUsername", context.Background(), &model.User{Username: username, IsNew: false})

		// Assert that there are no more calls to the UserRepository
		mockUserRepo.AssertExpectations(t)
	})

	t.Run("Non-existing user", func(t *testing.T) {
		_, _, mockUserRepo, _, authService := initMock()

		// Mock the FindUserByUsername method to return a non-existing user
		mockUserRepo.On("FindUserByUsername", context.Background(), mock.AnythingOfType("*model.User")).
			Return(nil, errors.ErrResourceNotFound)

		// Call the SignIn method
		result, err := authService.SignIn(context.Background(), username)

		// Assert the error
		assert.Error(t, err)
		assert.Nil(t, result)

		// Assert that the FindUserByUsername method was called with the expected arguments
		mockUserRepo.AssertCalled(t, "FindUserByUsername", context.Background(), &model.User{Username: username, IsNew: false})

		// Assert that there are no more calls to the UserRepository
		mockUserRepo.AssertExpectations(t)
	})
}

func TestFindUserFromODIC(t *testing.T) {

	// Prepare test data
	ctx := context.Background()
	username := "testuser"
	email := "test@example.com"
	firstName := "John"
	lastName := "Doe"

	t.Run("Unauthorized", func(t *testing.T) {
		_, mockCtrl, _, _, authService := initMock()

		// Mock the CheckPrivilege method to return true
		mockCtrl.On("CheckPrivilege", mock.Anything).
			Return(false)

		// Call the FindUserFromODIC method
		_, err := authService.FindUserFromODIC(ctx, &username)

		// Assert the result and error
		assert.Error(t, err)
		assert.Equal(t, errors.ErrUnAuthorized.Error(), err.Error())
	})

	t.Run("Empty user from OIDC", func(t *testing.T) {
		mockCloak, mockCtrl, _, _, authService := initMock()

		// Mock the CheckPrivilege method to return true
		mockCtrl.On("CheckPrivilege", mock.Anything).
			Return(true)

		// Mock the GetUsers method to return a single enabled user
		mockCloak.On("GetUsers", ctx, mock.Anything, mock.Anything, gocloak.GetUsersParams{
			Username: &username,
		}).Return([]*gocloak.User{}, nil)

		// Call the FindUserFromODIC method
		_, err := authService.FindUserFromODIC(ctx, &username)

		// Assert the result and error
		assert.Error(t, err)
		assert.Equal(t, errors.ErrUserNotFound.Error(), err.Error())
	})

	t.Run("Not enable user from OIDC", func(t *testing.T) {
		mockCloak, mockCtrl, _, _, authService := initMock()

		// Mock the CheckPrivilege method to return true
		mockCtrl.On("CheckPrivilege", mock.Anything).
			Return(true)

		// Mock the GetUsers method to return a single enabled user
		mockCloak.On("GetUsers", ctx, mock.Anything, mock.Anything, gocloak.GetUsersParams{
			Username: &username,
		}).Return([]*gocloak.User{
			{
				Username:  &username,
				Email:     &email,
				FirstName: &firstName,
				LastName:  &lastName,
				Enabled:   gocloak.BoolP(false),
			},
		}, nil)

		// Call the FindUserFromODIC method
		_, err := authService.FindUserFromODIC(ctx, &username)

		// Assert the result and error
		assert.Error(t, err)
		assert.Equal(t, errors.ErrUserNotFound.Error(), err.Error())
	})

	t.Run("Admin privilege, user found and enabled", func(t *testing.T) {
		mockCloak, mockCtrl, _, _, authService := initMock()

		// Mock the CheckPrivilege method to return true
		mockCtrl.On("CheckPrivilege", mock.Anything).
			Return(true)

		// Mock the GetUsers method to return a single enabled user
		mockCloak.On("GetUsers", ctx, mock.Anything, mock.Anything, gocloak.GetUsersParams{
			Username: &username,
		}).Return([]*gocloak.User{
			{
				Username:  &username,
				Email:     &email,
				FirstName: &firstName,
				LastName:  &lastName,
				Enabled:   gocloak.BoolP(true),
			},
		}, nil)

		// Call the FindUserFromODIC method
		result, err := authService.FindUserFromODIC(ctx, &username)

		// Assert the result and error
		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, username, result.Username)
	})
}

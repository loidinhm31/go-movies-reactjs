package service

import (
	"context"
	"movies-service/internal/control"
	"movies-service/internal/models"
	"movies-service/internal/users"
)

type managementCtrl struct {
	userRepository users.UserRepository
}

func NewManagementCtrl(userRepository users.UserRepository) control.Service {
	return &managementCtrl{
		userRepository: userRepository,
	}
}

func (mc *managementCtrl) CheckPrivilege(username string) bool {
	user, err := mc.userRepository.FindUserByUsername(context.Background(), &models.User{
		Username: username,
		IsNew:    false,
	})
	if err != nil {
		return false
	}
	if user.Role.RoleCode == "ADMIN" || user.Role.RoleCode == "MOD" {
		return true
	}
	return false
}

func (mc *managementCtrl) CheckAdminPrivilege(username string) bool {
	user, err := mc.userRepository.FindUserByUsername(context.Background(), &models.User{
		Username: username,
		IsNew:    false,
	})
	if err != nil {
		return false
	}
	if user.Role.RoleCode == "ADMIN" {
		return true
	}
	return false
}

func (mc *managementCtrl) CheckUser(username string) bool {
	user, err := mc.userRepository.FindUserByUsername(context.Background(), &models.User{
		Username: username,
		IsNew:    false,
	})
	if err != nil {
		return false
	}
	if user.Username == username && user.Role.RoleCode != "BANNED" {
		return true
	}
	return false
}

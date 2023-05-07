package service

import (
	"context"
	"movies-service/internal/auth"
	"movies-service/internal/control"
)

type managementCtrl struct {
	userRepository auth.UserRepository
}

func NewManagementCtrl(userRepository auth.UserRepository) control.Service {
	return &managementCtrl{
		userRepository: userRepository,
	}
}

func (mc *managementCtrl) CheckPrivilege(username string) bool {
	user, err := mc.userRepository.FindUserByUsername(context.Background(), username)
	if err != nil {
		return false
	}
	if user.Role.RoleCode == "ADMIN" || user.Role.RoleCode == "MOD" {
		return true
	}
	return false
}

func (mc *managementCtrl) CheckUser(username string) bool {
	user, err := mc.userRepository.FindUserByUsername(context.Background(), username)
	if err != nil {
		return false
	}
	if user.Username == username && user.Role.RoleCode != "BANNED" {
		return true
	}
	return false
}

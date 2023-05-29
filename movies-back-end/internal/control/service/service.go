package service

import (
	"context"
	"movies-service/internal/control"
	"movies-service/internal/model"
	"movies-service/internal/user"
)

type managementCtrl struct {
	userRepository user.UserRepository
}

func NewManagementCtrl(userRepository user.UserRepository) control.Service {
	return &managementCtrl{
		userRepository: userRepository,
	}
}

func (mc *managementCtrl) CheckPrivilege(username string) bool {
	theUser, err := mc.userRepository.FindUserByUsername(context.Background(), &model.User{
		Username: username,
		IsNew:    false,
	})
	if err != nil {
		return false
	}
	if theUser.Role.RoleCode == "ADMIN" || theUser.Role.RoleCode == "MOD" {
		return true
	}
	return false
}

func (mc *managementCtrl) CheckAdminPrivilege(username string) bool {
	theUser, err := mc.userRepository.FindUserByUsername(context.Background(), &model.User{
		Username: username,
		IsNew:    false,
	})
	if err != nil {
		return false
	}
	if theUser.Role.RoleCode == "ADMIN" {
		return true
	}
	return false
}

func (mc *managementCtrl) CheckUser(username string) bool {
	theUser, err := mc.userRepository.FindUserByUsername(context.Background(), &model.User{
		Username: username,
		IsNew:    false,
	})
	if err != nil {
		return false
	}
	if theUser.Username == username && theUser.Role.RoleCode != "BANNED" {
		return true
	}
	return false
}

package service

import (
	"context"
	"fmt"
	"log"
	"movies-service/internal/control"
	"movies-service/internal/dto"
	"movies-service/internal/errors"
	"movies-service/internal/middlewares"
	"movies-service/internal/models"
	"movies-service/internal/roles"
	"movies-service/internal/users"
	"movies-service/pkg/pagination"
	"strconv"
	"strings"
	"time"
)

type userService struct {
	mgmtCtrl       control.Service
	roleRepository roles.Repository
	userRepository users.UserRepository
}

func NewUserService(mgmtCtrl control.Service, roleRepository roles.Repository, userRepository users.UserRepository) users.Service {
	return &userService{
		mgmtCtrl:       mgmtCtrl,
		roleRepository: roleRepository,
		userRepository: userRepository,
	}
}

func (u *userService) GetUsers(ctx context.Context, pageRequest *pagination.PageRequest, key string, isNew string) (*pagination.Page[*dto.UserDto], error) {
	log.Println("checking admin privilege...")
	username := fmt.Sprintf("%s", ctx.Value(middlewares.CtxUserKey))
	if !u.mgmtCtrl.CheckAdminPrivilege(username) {
		return nil, errors.ErrUnAuthorized
	}

	isNewBool, _ := strconv.ParseBool(isNew)

	page := &pagination.Page[*models.User]{}

	userResults, err := u.userRepository.FindAllUsers(ctx, pageRequest, page, key, isNewBool)
	if err != nil {
		log.Println(err)
		return nil, errors.ErrResourceNotFound
	}

	var userDtos []*dto.UserDto
	for _, u := range userResults.Content {
		userDtos = append(userDtos, &dto.UserDto{
			ID:        u.ID,
			Username:  u.Username,
			Email:     u.Email,
			FirstName: u.FirstName,
			LastName:  u.LastName,
			IsNew:     u.IsNew,
			Role: dto.RoleDto{
				ID:       u.Role.ID,
				RoleName: u.Role.RoleName,
				RoleCode: u.Role.RoleCode,
			},
			CreatedAt: u.CreatedAt,
			UpdatedAt: u.UpdatedAt,
		})
	}
	return &pagination.Page[*dto.UserDto]{
		PageSize:      pageRequest.PageSize,
		PageNumber:    pageRequest.PageNumber,
		Sort:          pageRequest.Sort,
		TotalElements: userResults.TotalElements,
		TotalPages:    userResults.TotalPages,
		Content:       userDtos,
	}, nil
}

func (u *userService) UpdateUserRole(ctx context.Context, userDto *dto.UserDto) error {
	log.Println("checking admin privilege...")
	username := fmt.Sprintf("%s", ctx.Value(middlewares.CtxUserKey))
	if !u.mgmtCtrl.CheckAdminPrivilege(username) {
		return errors.ErrUnAuthorized
	}

	user, err := u.userRepository.FindUserById(ctx, userDto.ID)
	if err != nil {
		return err
	}

	role, err := u.roleRepository.FindRoleByRoleCode(ctx, userDto.Role.RoleCode)
	if err != nil {
		return err
	}

	err = u.userRepository.UpdateUserRole(ctx, user.ID, role.ID)
	if err != nil {
		return err
	}
	return nil
}

func (u *userService) AddOidcUser(ctx context.Context, userDto *dto.UserDto) (*dto.UserDto, error) {
	log.Println("Checking user...")
	fmtUsername := strings.ToLower(userDto.Username)
	euser, _ := u.userRepository.FindUserByUsername(ctx, &models.User{
		Username: fmtUsername,
		IsNew:    false,
	})
	if euser != nil && !userDto.IsNew {
		return nil, errors.ErrUserExisted
	}

	log.Println("checking admin privilege...")
	author := fmt.Sprintf("%s", ctx.Value(middlewares.CtxUserKey))
	if !u.mgmtCtrl.CheckAdminPrivilege(author) {
		return nil, errors.ErrUnAuthorized
	}

	getRole, err := u.roleRepository.FindRoleByRoleCode(ctx, userDto.Role.RoleCode)

	err = u.userRepository.InsertUser(ctx, &models.User{
		Username:  userDto.Username,
		Email:     userDto.Email,
		FirstName: userDto.FirstName,
		LastName:  userDto.LastName,
		Role:      getRole,
		IsNew:     false,
		CreatedAt: time.Now(),
		CreatedBy: author,
		UpdatedAt: time.Now(),
		UpdatedBy: author,
	})
	if err != nil {
		return nil, err
	}
	return userDto, nil
}

package service

import (
	"context"
	"errors"
	"github.com/Nerzal/gocloak/v13"
	"github.com/golang-jwt/jwt"
	"log"
	"movies-service/config"
	"movies-service/internal/dto"
	"movies-service/internal/models"
	"movies-service/internal/roles"
	"movies-service/internal/users"
	"movies-service/pkg/pagination"
	"strconv"
)

type AuthClaims struct {
	jwt.StandardClaims
	Username string `json:"username"`
	UserId   string `json:"userId"`
}

type userService struct {
	keycload       config.KeycloakConfig
	cloak          *gocloak.GoCloak
	roleRepository roles.Repository
	userRepository users.UserRepository
}

func NewUserService(keycloak config.KeycloakConfig, cloak *gocloak.GoCloak, roleRepository roles.Repository, userRepository users.UserRepository) users.Service {
	return &userService{
		keycload:       keycloak,
		cloak:          cloak,
		roleRepository: roleRepository,
		userRepository: userRepository,
	}
}

func (u *userService) GetUsers(ctx context.Context, pageRequest *pagination.PageRequest, key string, isNew string) (*pagination.Page[*dto.UserDto], error) {
	isNewBool, _ := strconv.ParseBool(isNew)

	page := &pagination.Page[*models.User]{}

	userResults, err := u.userRepository.FindAllUsers(ctx, pageRequest, page, key, isNewBool)
	if err != nil {
		log.Println(err)
		return nil, errors.New("not found")
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

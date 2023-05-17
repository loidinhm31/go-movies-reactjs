package service

import (
	"context"
	"github.com/Nerzal/gocloak/v13"
	"github.com/golang-jwt/jwt"
	"movies-service/config"
	"movies-service/internal/auth"
	"movies-service/internal/dto"
	"movies-service/internal/errors"
	"movies-service/internal/models"
	"movies-service/internal/roles"
	"movies-service/internal/users"
	"strings"
	"time"
)

type AuthClaims struct {
	jwt.StandardClaims
	Username string `json:"username"`
	UserId   string `json:"userId"`
}

type authService struct {
	keycloak       config.KeycloakConfig
	cloak          *gocloak.GoCloak
	roleRepository roles.Repository
	userRepository users.UserRepository
}

func NewAuthService(keycloak config.KeycloakConfig, cloak *gocloak.GoCloak, roleRepository roles.Repository, userRepository users.UserRepository) auth.Service {
	return &authService{
		keycloak:       keycloak,
		cloak:          cloak,
		roleRepository: roleRepository,
		userRepository: userRepository,
	}
}

func (a *authService) SignUp(ctx context.Context, userDto *dto.UserDto) (*dto.UserDto, error) {
	fmtUsername := strings.ToLower(userDto.Username)
	euser, _ := a.userRepository.FindUserByUsername(ctx, &models.User{
		Username: fmtUsername,
		IsNew:    false,
	})
	if euser != nil && !userDto.IsNew {
		return nil, errors.ErrUserExisted
	}

	role, err := a.roleRepository.FindRoleByRoleCode(ctx, "BANNED")
	if err != nil {
		return nil, err
	}

	user := &models.User{
		Username:  fmtUsername,
		FirstName: userDto.FirstName,
		LastName:  userDto.LastName,
		Email:     userDto.Email,
		IsNew:     userDto.IsNew,
		CreatedAt: time.Now(),
		CreatedBy: "keycloak",
		UpdatedAt: time.Now(),
		UpdatedBy: "keycloak",
		Role:      role,
	}
	err = a.userRepository.InsertUser(ctx, user)
	if err != nil {
		return nil, err
	}
	return userDto, nil
}

func (a *authService) SignIn(ctx context.Context, username string) (*dto.UserDto, error) {
	user, _ := a.userRepository.FindUserByUsername(ctx, &models.User{
		Username: username,
		IsNew:    false,
	})
	if user == nil {
		return nil, errors.ErrNotFound
	}

	userDto := &dto.UserDto{
		ID:        user.ID,
		Username:  user.Username,
		Email:     user.Email,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Role: dto.RoleDto{
			ID:        user.Role.ID,
			RoleName:  user.Role.RoleName,
			RoleCode:  user.Role.RoleCode,
			CreatedAt: user.Role.CreatedAt,
			UpdatedAt: user.Role.UpdatedAt,
		},
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
	return userDto, nil
}

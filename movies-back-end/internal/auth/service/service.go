package service

import (
	"context"
	"fmt"
	"github.com/Nerzal/gocloak/v13"
	"github.com/golang-jwt/jwt"
	"log"
	"movies-service/config"
	"movies-service/internal/auth"
	"movies-service/internal/control"
	"movies-service/internal/dto"
	"movies-service/internal/errors"
	"movies-service/internal/middlewares"
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
	mgmtCtrl       control.Service
	roleRepository roles.Repository
	userRepository users.UserRepository
}

func NewAuthService(keycloak config.KeycloakConfig, cloak *gocloak.GoCloak, mgmtCtrl control.Service, roleRepository roles.Repository, userRepository users.UserRepository) auth.Service {
	return &authService{
		keycloak:       keycloak,
		cloak:          cloak,
		mgmtCtrl:       mgmtCtrl,
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
		return nil, errors.ErrResourceNotFound
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

func (a *authService) FindUserFromODIC(ctx context.Context, username *string) (*dto.UserDto, error) {
	log.Println("checking admin privilege...")
	author := fmt.Sprintf("%s", ctx.Value(middlewares.CtxUserKey))
	if !a.mgmtCtrl.CheckPrivilege(author) {
		return nil, errors.ErrResourceNotFound
	}

	accessToken := fmt.Sprintf("%s", ctx.Value(middlewares.CtxAccessToken))
	getUsers, err := a.cloak.GetUsers(ctx, accessToken, a.keycloak.Realm, gocloak.GetUsersParams{
		Username: username,
	})
	if err != nil {
		return nil, err
	}
	if len(getUsers) == 0 {
		return nil, nil
	}

	oidcUser := getUsers[0]
	if !*oidcUser.Enabled {
		return nil, errors.ErrInvalidClient
	}

	return &dto.UserDto{
		Username:  *oidcUser.Username,
		Email:     *oidcUser.Email,
		FirstName: *oidcUser.FirstName,
		LastName:  *oidcUser.LastName,
	}, nil
}

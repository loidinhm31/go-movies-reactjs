package service

import (
	"context"
	"fmt"
	"github.com/Nerzal/gocloak/v13"
	"log"
	"movies-service/config"
	"movies-service/internal/auth"
	"movies-service/internal/cloak"
	"movies-service/internal/common/dto"
	"movies-service/internal/common/entity"
	"movies-service/internal/control"
	"movies-service/internal/errors"
	"movies-service/internal/middlewares"
	"movies-service/internal/role"
	"movies-service/internal/user"
	"strings"
	"time"
)

type authService struct {
	keycloak       config.KeycloakConfig
	cloak          cloak.GoCloakClientInterface
	mgmtCtrl       control.Service
	roleRepository role.Repository
	userRepository user.UserRepository
}

func NewAuthService(keycloak config.KeycloakConfig, cloak cloak.GoCloakClientInterface, mgmtCtrl control.Service, roleRepository role.Repository, userRepository user.UserRepository) auth.Service {
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
	euser, _ := a.userRepository.FindUserByUsername(ctx, fmtUsername)
	if euser != nil {
		return nil, errors.ErrUserExisted
	}

	theRole, err := a.roleRepository.FindRoleByRoleCode(ctx, "BANNED")
	if err != nil {
		return nil, err
	}

	theUser := &entity.User{
		Username:  fmtUsername,
		FirstName: userDto.FirstName,
		LastName:  userDto.LastName,
		Email:     userDto.Email,
		IsNew:     userDto.IsNew,
		CreatedAt: time.Now(),
		CreatedBy: "keycloak",
		UpdatedAt: time.Now(),
		UpdatedBy: "keycloak",
		Role:      theRole,
	}
	err = a.userRepository.InsertUser(ctx, theUser)
	if err != nil {
		return nil, err
	}
	return userDto, nil
}

func (a *authService) SignIn(ctx context.Context, username string) (*dto.UserDto, error) {
	theUser, _ := a.userRepository.FindUserByUsernameAndIsNew(ctx, username, false)
	if theUser == nil {
		return nil, errors.ErrResourceNotFound
	}

	userDto := &dto.UserDto{
		ID:        theUser.ID,
		Username:  theUser.Username,
		Email:     theUser.Email,
		FirstName: theUser.FirstName,
		LastName:  theUser.LastName,
		Role: dto.RoleDto{
			ID:        theUser.Role.ID,
			RoleName:  theUser.Role.RoleName,
			RoleCode:  theUser.Role.RoleCode,
			CreatedAt: theUser.Role.CreatedAt,
			UpdatedAt: theUser.Role.UpdatedAt,
		},
		CreatedAt: theUser.CreatedAt,
		UpdatedAt: theUser.UpdatedAt,
	}
	return userDto, nil
}

func (a *authService) FindUserFromODIC(ctx context.Context, username *string) (*dto.UserDto, error) {
	log.Println("checking admin privilege...")
	author := fmt.Sprintf("%s", ctx.Value(middlewares.CtxUserKey))
	if !a.mgmtCtrl.CheckPrivilege(author) {
		return nil, errors.ErrUnAuthorized
	}

	accessToken := fmt.Sprintf("%s", ctx.Value(middlewares.CtxAccessToken))
	getUsers, err := a.cloak.GetUsers(ctx, accessToken, a.keycloak.Realm, gocloak.GetUsersParams{
		Username: username,
	})
	if err != nil {
		return nil, err
	}
	if len(getUsers) == 0 {
		return nil, errors.ErrUserNotFound
	}

	oidcUser := getUsers[0]
	if !*oidcUser.Enabled {
		return nil, errors.ErrUserNotFound
	}

	return &dto.UserDto{
		Username:  *oidcUser.Username,
		Email:     *oidcUser.Email,
		FirstName: *oidcUser.FirstName,
		LastName:  *oidcUser.LastName,
	}, nil
}

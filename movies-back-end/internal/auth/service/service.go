package service

import (
	"context"
	"github.com/golang-jwt/jwt"
	"movies-service/internal/auth"
	"movies-service/internal/dto"
	"movies-service/internal/errors"
	"movies-service/internal/models"
	"strings"
)

type AuthClaims struct {
	jwt.StandardClaims
	Username string `json:"username"`
	UserId   string `json:"userId"`
}

type authService struct {
	userRepository auth.UserRepository
}

func NewAuthService(
	userRepository auth.UserRepository,
) auth.Service {
	return &authService{
		userRepository: userRepository,
	}
}

func (a *authService) SignUp(ctx context.Context, userDto *dto.UserDto) error {
	fmtUsername := strings.ToLower(userDto.Username)
	euser, _ := a.userRepository.FindUserByUsername(ctx, fmtUsername)

	if euser != nil {
		return errors.ErrUserExisted
	}
	user := &models.User{
		Username:  fmtUsername,
		FirstName: userDto.FirstName,
		LastName:  userDto.LastName,
	}
	err := a.userRepository.InsertUser(ctx, user)
	if err != nil {
		return err
	}

	return nil
}

func (a *authService) SignIn(ctx context.Context, username string) (*dto.UserDto, error) {
	user, _ := a.userRepository.FindUserByUsername(ctx, username)
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

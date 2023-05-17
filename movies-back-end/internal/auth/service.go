package auth

import (
	"context"
	"movies-service/internal/dto"
)

type Service interface {
	SignUp(ctx context.Context, userDto *dto.UserDto) (*dto.UserDto, error)
	SignIn(ctx context.Context, username string) (*dto.UserDto, error)
}

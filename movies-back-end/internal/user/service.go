package user

import (
	"context"
	"movies-service/internal/dto"
	"movies-service/pkg/pagination"
)

type Service interface {
	GetUsers(ctx context.Context, pageRequest *pagination.PageRequest, key string, isNew bool) (*pagination.Page[*dto.UserDto], error)
	UpdateUserRole(ctx context.Context, userDto *dto.UserDto) error
	AddOidcUser(ctx context.Context, userDto *dto.UserDto) (*dto.UserDto, error)
}

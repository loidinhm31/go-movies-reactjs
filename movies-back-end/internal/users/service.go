package users

import (
	"context"
	"movies-service/internal/dto"
	"movies-service/pkg/pagination"
)

type Service interface {
	GetUsers(ctx context.Context, pageRequest *pagination.PageRequest, key string, isNew string) (*pagination.Page[*dto.UserDto], error)
	UpdateUserRole(ctx context.Context, userDto *dto.UserDto) error
}

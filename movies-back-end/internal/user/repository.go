package user

import (
	"context"
	"movies-service/internal/common/entity"
	"movies-service/pkg/pagination"
)

type UserRepository interface {
	InsertUser(ctx context.Context, user *entity.User) error
	FindUserByUsername(ctx context.Context, username string) (*entity.User, error)
	FindUserByUsernameAndIsNew(ctx context.Context, username string, isNew bool) (*entity.User, error)
	FindUserByID(ctx context.Context, userID uint) (*entity.User, error)
	FindAllUsers(ctx context.Context, pageRequest *pagination.PageRequest, page *pagination.Page[*entity.User], key string, isNew bool) (*pagination.Page[*entity.User], error)
	UpdateUserRole(ctx context.Context, userId uint, roleId uint) error
}

package user

import (
	"context"
	"movies-service/internal/common/model"
	"movies-service/pkg/pagination"
)

type UserRepository interface {
	InsertUser(ctx context.Context, user *model.User) error
	FindUserByUsername(ctx context.Context, username string) (*model.User, error)
	FindUserByUsernameAndIsNew(ctx context.Context, username string, isNew bool) (*model.User, error)
	FindUserByID(ctx context.Context, userID uint) (*model.User, error)
	FindAllUsers(ctx context.Context, pageRequest *pagination.PageRequest, page *pagination.Page[*model.User], key string, isNew bool) (*pagination.Page[*model.User], error)
	UpdateUserRole(ctx context.Context, userId uint, roleId uint) error
}

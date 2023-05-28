package user

import (
	"context"
	"movies-service/internal/model"
	"movies-service/pkg/pagination"
)

type UserRepository interface {
	InsertUser(ctx context.Context, user *model.User) error
	FindUserByUsername(ctx context.Context, user *model.User) (*model.User, error)
	FindUserByID(ctx context.Context, userID int) (*model.User, error)
	FindAllUsers(ctx context.Context, pageRequest *pagination.PageRequest, page *pagination.Page[*model.User], key string, isNew bool) (*pagination.Page[*model.User], error)
	UpdateUserRole(ctx context.Context, userId int, roleId int) error
}

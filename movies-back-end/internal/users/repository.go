package users

import (
	"context"
	"movies-service/internal/models"
	"movies-service/pkg/pagination"
)

type UserRepository interface {
	InsertUser(ctx context.Context, user *models.User) error
	FindUserByUsername(ctx context.Context, user *models.User) (*models.User, error)
	FindUserById(ctx context.Context, userId int) (*models.User, error)
	FindAllUsers(ctx context.Context, pageRequest *pagination.PageRequest, page *pagination.Page[*models.User], key string, isNew bool) (*pagination.Page[*models.User], error)
	UpdateUserRole(ctx context.Context, userId int, roleId int) error
}

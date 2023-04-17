package auth

import (
	"context"
	"movies-service/internal/models"
)

type UserRepository interface {
	InsertUser(ctx context.Context, user *models.User) error
	FindUserByUsername(ctx context.Context, username string) (*models.User, error)
	FindUserById(ctx context.Context, userId string) (*models.User, error)
}

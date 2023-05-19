package roles

import (
	"context"
	"movies-service/internal/models"
)

type Repository interface {
	FindRoleByRoleCode(ctx context.Context, username string) (*models.Role, error)
	FindRoles(ctx context.Context) ([]*models.Role, error)
}

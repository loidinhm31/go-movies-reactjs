package role

import (
	"context"
	"movies-service/internal/model"
)

type Repository interface {
	FindRoleByRoleCode(ctx context.Context, username string) (*model.Role, error)
	FindRoles(ctx context.Context) ([]*model.Role, error)
}

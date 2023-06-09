package role

import (
	"context"
	"movies-service/internal/common/entity"
)

type Repository interface {
	FindRoleByRoleCode(ctx context.Context, username string) (*entity.Role, error)
	FindRoles(ctx context.Context) ([]*entity.Role, error)
}

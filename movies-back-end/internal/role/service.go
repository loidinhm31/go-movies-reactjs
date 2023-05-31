package role

import (
	"context"
	"movies-service/internal/dto"
)

type Service interface {
	GetAllRoles(ctx context.Context) ([]*dto.RoleDto, error)
}

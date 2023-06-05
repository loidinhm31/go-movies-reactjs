package role

import (
	"context"
	"movies-service/internal/common/dto"
)

type Service interface {
	GetAllRoles(ctx context.Context) ([]*dto.RoleDto, error)
}

package service

import (
	"context"
	"movies-service/internal/dto"
	"movies-service/internal/role"
)

type roleService struct {
	roleRepository role.Repository
}

func NewRoleService(roleRepository role.Repository) role.Service {
	return &roleService{
		roleRepository: roleRepository,
	}
}

func (r roleService) GetAllRoles(ctx context.Context) ([]*dto.RoleDto, error) {
	allRoles, err := r.roleRepository.FindRoles(ctx)
	if err != nil {
		return nil, err
	}

	var roleDtos []*dto.RoleDto
	for _, r := range allRoles {
		roleDtos = append(roleDtos, &dto.RoleDto{
			ID:       r.ID,
			RoleName: r.RoleName,
			RoleCode: r.RoleCode,
		})
	}
	return roleDtos, nil
}

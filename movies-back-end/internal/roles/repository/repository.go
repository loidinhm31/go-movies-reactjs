package repository

import (
	"context"
	"gorm.io/gorm"
	"movies-service/config"
	"movies-service/internal/models"
	"movies-service/internal/roles"
)

type roleRepository struct {
	cfg *config.Config
	db  *gorm.DB
}

func NewRoleRepository(cfg *config.Config, db *gorm.DB) roles.Repository {
	return &roleRepository{cfg: cfg, db: db}
}

func (rr roleRepository) FindRoleByRoleCode(ctx context.Context, roleCode string) (*models.Role, error) {
	var role models.Role
	err := rr.db.WithContext(ctx).Where(&models.Role{
		RoleCode: roleCode,
	}).Find(&role).Error

	if err != nil {
		return nil, err
	}

	return &role, nil
}

func (rr roleRepository) FindRoles(ctx context.Context) ([]*models.Role, error) {
	var role []*models.Role
	err := rr.db.WithContext(ctx).Find(&role).Error

	if err != nil {
		return nil, err
	}

	return role, nil
}

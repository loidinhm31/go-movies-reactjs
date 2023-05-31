package repository

import (
	"context"
	"gorm.io/gorm"
	"movies-service/config"
	"movies-service/internal/model"
	"movies-service/internal/role"
)

type roleRepository struct {
	cfg *config.Config
	db  *gorm.DB
}

func NewRoleRepository(cfg *config.Config, db *gorm.DB) role.Repository {
	return &roleRepository{cfg: cfg, db: db}
}

func (rr roleRepository) FindRoleByRoleCode(ctx context.Context, roleCode string) (*model.Role, error) {
	var result model.Role
	err := rr.db.WithContext(ctx).Where(&model.Role{
		RoleCode: roleCode,
	}).Find(&result).Error

	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (rr roleRepository) FindRoles(ctx context.Context) ([]*model.Role, error) {
	var results []*model.Role
	err := rr.db.WithContext(ctx).Find(&results).Error

	if err != nil {
		return nil, err
	}

	return results, nil
}

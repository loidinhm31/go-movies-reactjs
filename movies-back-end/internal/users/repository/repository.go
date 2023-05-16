package repository

import (
	"context"
	"gorm.io/gorm"
	"movies-service/config"
	"movies-service/internal/models"
	"movies-service/internal/users"
	"movies-service/pkg/pagination"
	"strings"
)

type userRepository struct {
	cfg *config.Config
	db  *gorm.DB
}

func NewUserRepository(cfg *config.Config, db *gorm.DB) users.UserRepository {
	return &userRepository{cfg: cfg, db: db}
}

func (ur *userRepository) InsertUser(ctx context.Context, user *models.User) error {
	result := ur.db.WithContext(ctx).Create(&user)

	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (ur *userRepository) FindUserByUsername(ctx context.Context, user *models.User) (*models.User, error) {
	err := ur.db.WithContext(ctx).Where(user).
		Preload("Role").First(&user).Error

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (ur *userRepository) FindUserById(ctx context.Context, userId int) (*models.User, error) {
	var user models.User
	err := ur.db.WithContext(ctx).Where(&models.User{
		ID: userId,
	}).First(&user).Error

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (ur *userRepository) FindAllUsers(ctx context.Context, pageRequest *pagination.PageRequest, page *pagination.Page[*models.User], key string, isNew bool) (*pagination.Page[*models.User], error) {
	var allUsers []*models.User
	var totalRows int64

	tx := ur.db.WithContext(ctx)
	if ur.cfg.Server.Debug {
		tx = tx.Debug()
	}
	tx = tx.Model(&models.User{}).
		Where("is_new = ?", isNew)

	if key != "" {
		wildCardKey := "%" + strings.ToLower(key) + "%"
		tx = tx.Where("(LOWER(username) LIKE ? OR LOWER(first_name) LIKE ? OR LOWER(last_name) LIKE ? OR LOWER(email) LIKE ?)",
			wildCardKey, wildCardKey, wildCardKey, wildCardKey)
	}

	err := tx.Preload("Role").
		Count(&totalRows).
		Scopes(pagination.PageImplCountCriteria[*models.User](totalRows, pageRequest, page)).
		Find(&allUsers).Error
	if err != nil {
		return nil, err
	}
	page.Content = allUsers
	return page, nil
}

func (ur *userRepository) UpdateUserRole(ctx context.Context, userId int, roleId int) error {
	tx := ur.db.WithContext(ctx)
	if ur.cfg.Server.Debug {
		tx = tx.Debug()
	}

	err := tx.Model(&models.User{}).Where("id = ?", userId).Updates(map[string]interface{}{"role_id": roleId, "is_new": false}).Error
	if err != nil {
		return err
	}
	return nil
}

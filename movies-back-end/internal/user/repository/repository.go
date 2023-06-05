package repository

import (
	"context"
	"gorm.io/gorm"
	"movies-service/config"
	"movies-service/internal/common/model"
	"movies-service/internal/user"
	"movies-service/pkg/pagination"
	"strings"
)

type userRepository struct {
	cfg *config.Config
	db  *gorm.DB
}

func NewUserRepository(cfg *config.Config, db *gorm.DB) user.UserRepository {
	return &userRepository{cfg: cfg, db: db}
}

func (ur *userRepository) InsertUser(ctx context.Context, user *model.User) error {
	result := ur.db.WithContext(ctx).Create(&user)

	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (ur *userRepository) FindUserByUsername(ctx context.Context, username string) (*model.User, error) {
	var theUser *model.User
	err := ur.db.WithContext(ctx).Where("username = ?", username).
		Preload("Role").First(&theUser).Error

	if err != nil {
		return nil, err
	}

	return theUser, nil
}

func (ur *userRepository) FindUserByUsernameAndIsNew(ctx context.Context, username string, isNew bool) (*model.User, error) {
	var theUser *model.User
	err := ur.db.WithContext(ctx).Where("username = ? AND is_new = ?", username, isNew).
		Preload("Role").First(&theUser).Error

	if err != nil {
		return nil, err
	}

	return theUser, nil
}

func (ur *userRepository) FindUserByID(ctx context.Context, userID uint) (*model.User, error) {
	var theUser model.User
	err := ur.db.WithContext(ctx).Where("id = ?", userID).First(&theUser).Error

	if err != nil {
		return nil, err
	}

	return &theUser, nil
}

func (ur *userRepository) FindAllUsers(ctx context.Context, pageRequest *pagination.PageRequest, page *pagination.Page[*model.User], key string, isNew bool) (*pagination.Page[*model.User], error) {
	var allUsers []*model.User
	var totalRows int64

	tx := ur.db.WithContext(ctx)
	if ur.cfg.Server.Debug {
		tx = tx.Debug()
	}
	tx = tx.Model(&model.User{}).
		Where("is_new = ?", isNew)

	if key != "" {
		wildCardKey := "%" + strings.ToLower(key) + "%"
		tx = tx.Where("(LOWER(username) LIKE ? OR LOWER(first_name) LIKE ? OR LOWER(last_name) LIKE ? OR LOWER(email) LIKE ?)",
			wildCardKey, wildCardKey, wildCardKey, wildCardKey)
	}

	err := tx.Preload("Role").
		Count(&totalRows).
		Scopes(pagination.PageImplCountCriteria[*model.User](totalRows, pageRequest, page)).
		Find(&allUsers).Error
	if err != nil {
		return nil, err
	}
	page.Content = allUsers
	return page, nil
}

func (ur *userRepository) UpdateUserRole(ctx context.Context, userID uint, roleID uint) error {
	tx := ur.db.WithContext(ctx)
	if ur.cfg.Server.Debug {
		tx = tx.Debug()
	}

	err := tx.Model(&model.User{}).Where("id = ?", userID).Updates(map[string]interface{}{"role_id": roleID, "is_new": false}).Error
	if err != nil {
		return err
	}
	return nil
}

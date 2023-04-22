package repository

import (
	"context"
	"gorm.io/gorm"
	"movies-service/config"
	"movies-service/internal/auth"
	"movies-service/internal/models"
	"strings"
)

type userRepository struct {
	cfg *config.Config
	db  *gorm.DB
}

func NewUserRepository(cfg *config.Config, db *gorm.DB) auth.UserRepository {
	return &userRepository{cfg: cfg, db: db}
}

func (ur *userRepository) InsertUser(ctx context.Context, user *models.User) error {
	result := ur.db.WithContext(ctx).Create(&user)

	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (ur *userRepository) FindUserByUsername(ctx context.Context, username string) (*models.User, error) {
	var user models.User
	err := ur.db.WithContext(ctx).Where(&models.User{
		Username: strings.ToLower(username),
	}).First(&user).Error

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (ur *userRepository) FindUserById(ctx context.Context, userId string) (*models.User, error) {
	var user models.User
	err := ur.db.WithContext(ctx).Where(&models.User{
		UserID: userId,
	}).First(&user).Error

	if err != nil {
		return nil, err
	}

	return &user, nil
}

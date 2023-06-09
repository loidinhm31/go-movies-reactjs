package repository

import (
	"context"
	"gorm.io/gorm"
	"movies-service/config"
	"movies-service/internal/common/entity"
	"movies-service/internal/rating"
)

type ratingRepository struct {
	cfg *config.Config
	db  *gorm.DB
}

func NewRatingRepository(cfg *config.Config, db *gorm.DB) rating.Repository {
	return &ratingRepository{cfg: cfg, db: db}
}

func (mr *ratingRepository) FindRatings(ctx context.Context) ([]*entity.Rating, error) {
	var results []*entity.Rating

	tx := mr.db.WithContext(ctx)
	if mr.cfg.Server.Debug {
		tx = tx.Debug()
	}

	err := tx.Find(&results).Error
	if err != nil {
		return nil, err
	}
	return results, nil
}

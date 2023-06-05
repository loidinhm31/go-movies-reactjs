package repository

import (
	"context"
	"gorm.io/gorm"
	"movies-service/config"
	"movies-service/internal/common/model"
	"movies-service/internal/rating"
)

type ratingRepository struct {
	cfg *config.Config
	db  *gorm.DB
}

func NewRatingRepository(cfg *config.Config, db *gorm.DB) rating.Repository {
	return &ratingRepository{cfg: cfg, db: db}
}

func (mr *ratingRepository) FindRatings(ctx context.Context) ([]*model.Rating, error) {
	var results []*model.Rating

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

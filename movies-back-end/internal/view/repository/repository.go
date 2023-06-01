package repository

import (
	"context"
	"gorm.io/gorm"
	"movies-service/config"
	"movies-service/internal/dto"
	"movies-service/internal/model"
	"movies-service/internal/view"
	"time"
)

type viewRepository struct {
	cfg *config.Config
	db  *gorm.DB
}

func NewViewRepository(cfg *config.Config, db *gorm.DB) view.Repository {
	return &viewRepository{cfg: cfg, db: db}
}

func (vr *viewRepository) InsertView(ctx context.Context, viewer *dto.Viewer) error {
	tx := vr.db.WithContext(ctx)
	if vr.cfg.Server.Debug {
		tx = tx.Debug()
	}
	result := tx.Create(&model.View{
		ViewedBy: viewer.Viewer,
		ViewedAt: time.Now(),
		MovieId:  viewer.MovieId,
	})

	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (vr *viewRepository) CountViewsByMovieId(ctx context.Context, movieId uint) (int64, error) {
	var totalViews int64

	tx := vr.db.WithContext(ctx)
	if vr.cfg.Server.Debug {
		tx = tx.Debug()
	}
	err := tx.Table("views").
		Where("views.movie_id = ?", movieId).
		Count(&totalViews).Error
	if err != nil {
		return 0, err
	}
	return totalViews, nil
}

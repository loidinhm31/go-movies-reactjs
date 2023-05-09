package repository

import (
	"context"
	"gorm.io/gorm"
	"movies-service/config"
	"movies-service/internal/dto"
	"movies-service/internal/models"
	"movies-service/internal/views"
	"strconv"
	"time"
)

type viewRepository struct {
	cfg *config.Config
	db  *gorm.DB
}

func NewViewRepository(cfg *config.Config, db *gorm.DB) views.Repository {
	return &viewRepository{cfg: cfg, db: db}
}

func (vr *viewRepository) InsertView(ctx context.Context, viewer *dto.Viewer) error {
	movieId, err := strconv.Atoi(viewer.MovieId)
	if err != nil {
		return err
	}

	tx := vr.db.WithContext(ctx)
	if vr.cfg.Server.Debug {
		tx = tx.Debug()
	}
	result := tx.Create(&models.View{
		ViewedBy: viewer.Viewer,
		ViewedAt: time.Now(),
		MovieId:  movieId,
	})

	if result.Error != nil {
		return result.Error
	}
	return nil
}
package repository

import (
	"context"
	"gorm.io/gorm"
	"movies-service/config"
	"movies-service/internal/models"
	"movies-service/internal/seasons"
)

type seasonRepository struct {
	cfg *config.Config
	db  *gorm.DB
}

func NewSeasonRepository(cfg *config.Config, db *gorm.DB) seasons.Repository {
	return &seasonRepository{cfg: cfg, db: db}
}

func (s seasonRepository) FindSeasonByID(ctx context.Context, id int) (*models.Season, error) {
	var seasonObject *models.Season

	tx := s.db.WithContext(ctx)
	if s.cfg.Server.Debug {
		tx = tx.Debug()
	}
	err := tx.Where("id = ?", id).First(&seasonObject).Error
	if err != nil {
		return nil, err
	}

	return seasonObject, nil
}

func (s seasonRepository) FindSeasonsByMovieID(ctx context.Context, movieID int) ([]*models.Season, error) {
	var seasonObjects []*models.Season

	tx := s.db.WithContext(ctx)
	if s.cfg.Server.Debug {
		tx = tx.Debug()
	}
	err := tx.Where("movie_id = ?", movieID).Order("air_date").Find(&seasonObjects).Error
	if err != nil {
		return nil, err
	}

	return seasonObjects, nil
}

func (s seasonRepository) InsertSeason(ctx context.Context, season *models.Season) error {
	tx := s.db.WithContext(ctx)
	if s.cfg.Server.Debug {
		tx = tx.Debug()
	}
	err := tx.Create(&season).Error
	if err != nil {
		return err
	}
	return nil
}

func (s seasonRepository) UpdateSeason(ctx context.Context, season *models.Season) error {
	tx := s.db.WithContext(ctx)

	if s.cfg.Server.Debug {
		tx = tx.Debug()
	}
	err := tx.Model(&models.Season{}).Where("id = ?", season.ID).
		Updates(season).Error
	if err != nil {
		return err
	}
	return nil
}

func (s seasonRepository) DeleteSeasonByID(ctx context.Context, id int) error {
	tx := s.db.WithContext(ctx)
	if s.cfg.Server.Debug {
		tx = tx.Debug()
	}
	err := tx.Where("id = ?", id).Delete(&models.Season{}).Error
	if err != nil {
		return err
	}
	return nil
}

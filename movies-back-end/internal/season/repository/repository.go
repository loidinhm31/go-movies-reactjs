package repository

import (
	"context"
	"gorm.io/gorm"
	"movies-service/config"
	"movies-service/internal/common/entity"
	"movies-service/internal/season"
)

type seasonRepository struct {
	cfg *config.Config
	db  *gorm.DB
}

func NewSeasonRepository(cfg *config.Config, db *gorm.DB) season.Repository {
	return &seasonRepository{cfg: cfg, db: db}
}

func (s seasonRepository) FindSeasonByID(ctx context.Context, id uint) (*entity.Season, error) {
	var seasonObject *entity.Season

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

func (s seasonRepository) FindSeasonsByMovieID(ctx context.Context, movieID uint) ([]*entity.Season, error) {
	var seasonObjects []*entity.Season

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

func (s seasonRepository) InsertSeason(ctx context.Context, season *entity.Season) error {
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

func (s seasonRepository) UpdateSeason(ctx context.Context, season *entity.Season) error {
	tx := s.db.WithContext(ctx)

	if s.cfg.Server.Debug {
		tx = tx.Debug()
	}
	err := tx.Model(&entity.Season{}).Where("id = ?", season.ID).
		Save(season).Error
	if err != nil {
		return err
	}
	return nil
}

func (s seasonRepository) DeleteSeasonByID(ctx context.Context, id uint) error {
	tx := s.db.WithContext(ctx)
	if s.cfg.Server.Debug {
		tx = tx.Debug()
	}
	err := tx.Where("id = ?", id).Delete(&entity.Season{}).Error
	if err != nil {
		return err
	}
	return nil
}

package repository

import (
	"context"
	"gorm.io/gorm"
	"movies-service/config"
	"movies-service/internal/common/entity"
	"movies-service/internal/episode"
)

type episodeRepository struct {
	cfg *config.Config
	db  *gorm.DB
}

func NewEpisodeRepository(cfg *config.Config, db *gorm.DB) episode.Repository {
	return &episodeRepository{cfg: cfg, db: db}
}

func (e episodeRepository) FindEpisodeByID(ctx context.Context, id uint) (*entity.Episode, error) {
	var episodeObject *entity.Episode

	tx := e.db.WithContext(ctx)
	if e.cfg.Server.Debug {
		tx = tx.Debug()
	}
	err := tx.Where("id = ?", id).Preload("Season").First(&episodeObject).Error
	if err != nil {
		return nil, err
	}

	return episodeObject, nil
}

func (e episodeRepository) FindEpisodesBySeasonID(ctx context.Context, seasonID uint) ([]*entity.Episode, error) {
	var episodeObjects []*entity.Episode

	tx := e.db.WithContext(ctx)
	if e.cfg.Server.Debug {
		tx = tx.Debug()
	}
	err := tx.Where("season_id = ?", seasonID).Order("air_date, name").
		Find(&episodeObjects).Error
	if err != nil {
		return nil, err
	}

	return episodeObjects, nil
}

func (e episodeRepository) InsertEpisode(ctx context.Context, episode *entity.Episode) error {
	tx := e.db.WithContext(ctx)
	if e.cfg.Server.Debug {
		tx = tx.Debug()
	}
	err := tx.Create(&episode).Error
	if err != nil {
		return err
	}
	return nil
}

func (e episodeRepository) UpdateEpisode(ctx context.Context, episode *entity.Episode) error {
	tx := e.db.WithContext(ctx)

	if e.cfg.Server.Debug {
		tx = tx.Debug()
	}
	err := tx.Model(&entity.Episode{}).Where("id = ?", episode.ID).
		Save(episode).Error
	if err != nil {
		return err
	}
	return nil
}

func (e episodeRepository) DeleteEpisodeByID(ctx context.Context, id uint) error {
	tx := e.db.WithContext(ctx)
	if e.cfg.Server.Debug {
		tx = tx.Debug()
	}
	err := tx.Where("id = ?", id).Delete(&entity.Episode{}).Error
	if err != nil {
		return err
	}
	return nil
}

func (e episodeRepository) DeleteEpisodeBySeasonID(ctx context.Context, seasonID uint) error {
	tx := e.db.WithContext(ctx)
	if e.cfg.Server.Debug {
		tx = tx.Debug()
	}

	err := tx.Where("season_id = ?", seasonID).Delete(&entity.Episode{}).Error
	if err != nil {
		return err
	}
	return nil
}

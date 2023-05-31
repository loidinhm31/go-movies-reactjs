package repository

import (
	"context"
	"gorm.io/gorm"
	"movies-service/config"
	"movies-service/internal/episode"
	"movies-service/internal/model"
)

type episodeRepository struct {
	cfg *config.Config
	db  *gorm.DB
}

func NewEpisodeRepository(cfg *config.Config, db *gorm.DB) episode.Repository {
	return &episodeRepository{cfg: cfg, db: db}
}

func (e episodeRepository) FindEpisodeByID(ctx context.Context, id int) (*model.Episode, error) {
	var episodeObject *model.Episode

	tx := e.db.WithContext(ctx)
	if e.cfg.Server.Debug {
		tx = tx.Debug()
	}
	err := tx.Where("id = ?", id).First(&episodeObject).Error
	if err != nil {
		return nil, err
	}

	return episodeObject, nil
}

func (e episodeRepository) FindEpisodesBySeasonID(ctx context.Context, seasonID int) ([]*model.Episode, error) {
	var episodeObjects []*model.Episode

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

func (e episodeRepository) InsertEpisode(ctx context.Context, episode *model.Episode) error {
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

func (e episodeRepository) UpdateEpisode(ctx context.Context, episode *model.Episode) error {
	tx := e.db.WithContext(ctx)

	if e.cfg.Server.Debug {
		tx = tx.Debug()
	}
	err := tx.Model(&model.Episode{}).Where("id = ?", episode.ID).
		Updates(episode).Error
	if err != nil {
		return err
	}
	return nil
}

func (e episodeRepository) DeleteEpisodeByID(ctx context.Context, id int) error {
	tx := e.db.WithContext(ctx)
	if e.cfg.Server.Debug {
		tx = tx.Debug()
	}
	err := tx.Where("id = ?", id).Delete(&model.Episode{}).Error
	if err != nil {
		return err
	}
	return nil
}

func (e episodeRepository) DeleteEpisodeBySeasonID(ctx context.Context, seasonID int) error {
	tx := e.db.WithContext(ctx)
	if e.cfg.Server.Debug {
		tx = tx.Debug()
	}

	err := tx.Where("season_id = ?", seasonID).Delete(&model.Episode{}).Error
	if err != nil {
		return err
	}
	return nil
}
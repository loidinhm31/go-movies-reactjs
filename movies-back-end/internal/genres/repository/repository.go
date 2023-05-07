package repository

import (
	"context"
	"gorm.io/gorm"
	"movies-service/config"
	"movies-service/internal/genres"
	"movies-service/internal/models"
)

type genreRepository struct {
	cfg *config.Config
	db  *gorm.DB
}

func NewGenreRepository(cfg *config.Config, db *gorm.DB) genres.GenreRepository {
	return &genreRepository{cfg: cfg, db: db}
}

func (gr *genreRepository) FindAllGenres(ctx context.Context) ([]*models.Genre, error) {
	var allGenres []*models.Genre

	tx := gr.db.WithContext(ctx)
	if gr.cfg.Server.Debug {
		tx = tx.Debug()
	}
	err := tx.Order("genre").
		Find(&allGenres).
		Error
	if err != nil {
		return nil, err
	}

	return allGenres, nil
}

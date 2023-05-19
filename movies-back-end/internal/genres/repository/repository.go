package repository

import (
	"context"
	"github.com/gin-gonic/gin"
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
	err := tx.Order("name").
		Find(&allGenres).
		Error
	if err != nil {
		return nil, err
	}

	return allGenres, nil
}

func (gr *genreRepository) FindGenreByName(ctx *gin.Context, genre *models.Genre) (*models.Genre, error) {
	var result models.Genre

	tx := gr.db.WithContext(ctx)
	if gr.cfg.Server.Debug {
		tx = tx.Debug()
	}
	err := tx.Model(&models.Genre{}).Where(genre).Find(result).Error
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (gr *genreRepository) InsertGenres(ctx *gin.Context, genres []*models.Genre) error {
	tx := gr.db.WithContext(ctx)
	if gr.cfg.Server.Debug {
		tx = tx.Debug()
	}
	err := tx.Create(genres).Error
	if err != nil {
		return err
	}
	return nil
}

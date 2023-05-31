package repository

import (
	"context"
	"gorm.io/gorm"
	"movies-service/config"
	"movies-service/internal/genre"
	"movies-service/internal/model"
)

type genreRepository struct {
	cfg *config.Config
	db  *gorm.DB
}

func NewGenreRepository(cfg *config.Config, db *gorm.DB) genre.Repository {
	return &genreRepository{cfg: cfg, db: db}
}

func (gr *genreRepository) FindAllGenres(ctx context.Context) ([]*model.Genre, error) {
	var allGenres []*model.Genre

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

func (gr *genreRepository) FindAllGenresByTypeCode(ctx context.Context, movieType string) ([]*model.Genre, error) {
	var allGenres []*model.Genre

	tx := gr.db.WithContext(ctx)
	if gr.cfg.Server.Debug {
		tx = tx.Debug()
	}
	err := tx.Order("name").
		Where("type_code = ?", movieType).
		Find(&allGenres).
		Error
	if err != nil {
		return nil, err
	}

	return allGenres, nil
}

func (gr *genreRepository) FindGenreByNameAndTypeCode(ctx context.Context, genre *model.Genre) (*model.Genre, error) {
	var result model.Genre

	tx := gr.db.WithContext(ctx)
	if gr.cfg.Server.Debug {
		tx = tx.Debug()
	}
	err := tx.Model(&model.Genre{}).Where("name = ? AND type_code = ?", genre.Name, genre.TypeCode).Find(result).Error
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (gr *genreRepository) InsertGenres(ctx context.Context, genres []*model.Genre) error {
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

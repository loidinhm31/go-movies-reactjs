package repository

import (
	"context"
	"gorm.io/gorm"
	"movies-service/internal/genres"
	"movies-service/internal/models"
)

type genreRepository struct {
	db *gorm.DB
}

func NewGenreRepository(db *gorm.DB) genres.GenreRepository {
	return &genreRepository{db: db}
}

func (gr *genreRepository) FindAllGenres(ctx context.Context) ([]*models.Genre, error) {
	var allGenres []*models.Genre

	err := gr.db.WithContext(ctx).Find(&allGenres).Error
	if err != nil {
		return nil, err
	}

	return allGenres, nil
}

package repository

import (
	"context"
	"gorm.io/gorm"
	"movies-service/config"
	"movies-service/internal/analysis"
	"movies-service/internal/models"
)

type analysisRepository struct {
	cfg *config.Config
	db  *gorm.DB
}

func NewAnalysisRepository(cfg *config.Config, db *gorm.DB) analysis.Repository {
	return &analysisRepository{cfg: cfg, db: db}
}

func (ar *analysisRepository) CountMoviesByGenre(ctx context.Context) ([]*models.GenreCount, error) {
	var result []*models.GenreCount

	tx := ar.db.WithContext(ctx)
	if ar.cfg.Server.Debug {
		tx = tx.Debug()
	}
	err := tx.Raw("SELECT g.genre, COUNT(mg.movie_id) AS num_movies " +
		"FROM genres g " +
		"INNER JOIN movies_genres mg on g.id = mg.genre_id " +
		"GROUP BY g.genre").
		Scan(&result).Error
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (ar *analysisRepository) CountMoviesByReleaseDate(ctx context.Context, year string, months []string) ([]*models.MovieCount, error) {
	var result []*models.MovieCount

	tx := ar.db.WithContext(ctx)
	if ar.cfg.Server.Debug {
		tx = tx.Debug()
	}
	err := tx.Raw("SELECT EXTRACT(YEAR FROM m.release_date) AS year, EXTRACT(MONTH FROM m.release_date) AS month, COUNT(m.id) AS num_movies "+
		"FROM movies m "+
		"WHERE EXTRACT(YEAR FROM m.release_date) = ? "+
		"AND EXTRACT(MONTH FROM m.release_date) IN ? "+
		"GROUP BY EXTRACT(YEAR FROM m.release_date), EXTRACT(MONTH FROM m.release_date);", year, months).
		Scan(&result).Error
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (ar *analysisRepository) CountMoviesByCreatedDate(ctx context.Context, year string, months []string) ([]*models.MovieCount, error) {
	var result []*models.MovieCount

	tx := ar.db.WithContext(ctx)
	if ar.cfg.Server.Debug {
		tx = tx.Debug()
	}
	err := tx.Raw("SELECT EXTRACT(YEAR FROM m.created_at) AS year, EXTRACT(MONTH FROM m.created_at) AS month, COUNT(m.id) AS num_movies "+
		"FROM movies m "+
		"WHERE EXTRACT(YEAR FROM m.created_at) = ? "+
		"AND EXTRACT(MONTH FROM m.created_at) IN ? "+
		"GROUP BY EXTRACT(YEAR FROM m.created_at), EXTRACT(MONTH FROM m.created_at);", year, months).
		Scan(&result).Error
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (ar *analysisRepository) CountViewsByGenreAndViewedDate(ctx context.Context, genre, year string, months []string) ([]*models.ViewCount, error) {
	var result []*models.ViewCount

	tx := ar.db.WithContext(ctx)
	if ar.cfg.Server.Debug {
		tx = tx.Debug()
	}
	err := tx.Raw("SELECT EXTRACT(YEAR FROM v.viewed_at) AS year, EXTRACT(MONTH FROM v.viewed_at) AS month, COUNT(v.id) AS num_viewers "+
		"FROM view v "+
		"INNER JOIN movies m on m.id = v.movie_id "+
		"INNER JOIN movies_genres mg on m.id = mg.movie_id "+
		"INNER JOIN genres g on g.id = mg.genre_id "+
		"WHERE LOWER(g.genre) = LOWER(?) "+
		"AND EXTRACT(YEAR FROM m.created_at) = ? "+
		"AND EXTRACT(MONTH FROM m.created_at) IN ? "+
		"GROUP BY EXTRACT(YEAR FROM v.viewed_at), EXTRACT(MONTH FROM v.viewed_at);", genre, year, months).
		Scan(&result).Error
	if err != nil {
		return nil, err
	}

	return result, nil
}

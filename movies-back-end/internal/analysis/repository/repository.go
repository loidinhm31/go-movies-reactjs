package repository

import (
	"context"
	"gorm.io/gorm"
	"movies-service/config"
	"movies-service/internal/analysis"
	"movies-service/internal/dto"
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
		"GROUP BY g.genre " +
		"ORDER BY g.genre;").
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

func (ar *analysisRepository) CountViewsByGenreAndViewedDate(ctx context.Context, request *dto.RequestData) ([]*models.ViewCount, error) {
	var result []*models.ViewCount

	tx := ar.db.WithContext(ctx)
	if ar.cfg.Server.Debug {
		tx = tx.Debug()
	}

	tx = tx.Table("views").
		Select("EXTRACT(YEAR FROM views.viewed_at) AS year, EXTRACT(MONTH FROM views.viewed_at) AS month, COUNT(views.id) AS num_viewers").
		InnerJoins("INNER JOIN movies on movies.id = views.movie_id").
		InnerJoins("INNER JOIN movies_genres on movies.id = movies_genres.movie_id").
		InnerJoins("INNER JOIN genres on genres.id = movies_genres.genre_id").
		Where("LOWER(genres.genre) = LOWER(?)", request.Genre)

	orBuild := ar.db
	for _, a := range request.Analysis {
		orBuild = orBuild.Or("EXTRACT(YEAR FROM views.viewed_at) = ? AND EXTRACT(MONTH FROM views.viewed_at) IN ?", a.Year, a.Months)
	}

	err := tx.Where(orBuild).Group("EXTRACT(YEAR FROM views.viewed_at), EXTRACT(MONTH FROM views.viewed_at)").
		Order("EXTRACT(YEAR FROM views.viewed_at), EXTRACT(MONTH FROM views.viewed_at)").
		Find(&result).Error

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (ar *analysisRepository) CountCumulativeViewsByGenreAndViewedDate(ctx context.Context, request *dto.RequestData) ([]*models.ViewCount, error) {
	var result []*models.ViewCount

	tx := ar.db.WithContext(ctx)
	if ar.cfg.Server.Debug {
		tx = tx.Debug()
	}

	tx = tx.Table("views").
		Select("EXTRACT(YEAR FROM views.viewed_at) AS year, "+
			"EXTRACT(MONTH FROM views.viewed_at) AS month, "+
			"COUNT(views.id) AS num_viewers, "+
			"SUM(COUNT(views.id)) OVER (ORDER BY EXTRACT(YEAR FROM views.viewed_at), EXTRACT(MONTH FROM views.viewed_at)) AS cumulative").
		InnerJoins("INNER JOIN movies on movies.id = views.movie_id").
		InnerJoins("INNER JOIN movies_genres on movies.id = movies_genres.movie_id").
		InnerJoins("INNER JOIN genres on genres.id = movies_genres.genre_id").
		Where("LOWER(genres.genre) = LOWER(?)", request.Genre)

	orBuild := ar.db
	for _, a := range request.Analysis {
		orBuild = orBuild.Or("EXTRACT(YEAR FROM views.viewed_at) = ? AND EXTRACT(MONTH FROM views.viewed_at) IN ?", a.Year, a.Months)
	}

	err := tx.Group("EXTRACT(YEAR FROM views.viewed_at), EXTRACT(MONTH FROM views.viewed_at)").
		Order("EXTRACT(YEAR FROM views.viewed_at), EXTRACT(MONTH FROM views.viewed_at)").
		Find(&result).Error
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (ar *analysisRepository) CountViewsByViewedDate(ctx context.Context, request *dto.RequestData) ([]*models.ViewCount, error) {
	var result []*models.ViewCount

	tx := ar.db.WithContext(ctx)
	if ar.cfg.Server.Debug {
		tx = tx.Debug()
	}

	tx = tx.Table("views").
		Select("EXTRACT(YEAR FROM views.viewed_at) AS year, EXTRACT(MONTH FROM views.viewed_at) AS month, COUNT(views.id) AS num_viewers")

	for _, a := range request.Analysis {
		tx.Or("EXTRACT(YEAR FROM views.viewed_at) = ? AND EXTRACT(MONTH FROM views.viewed_at) IN ?", a.Year, a.Months)
	}

	err := tx.Group("EXTRACT(YEAR FROM views.viewed_at), EXTRACT(MONTH FROM views.viewed_at)").
		Find(&result).Error
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (ar *analysisRepository) CountMoviesByGenreAndReleasedDate(ctx context.Context, request *dto.RequestData) ([]*models.MovieCount, error) {
	var result []*models.MovieCount

	tx := ar.db.WithContext(ctx)
	if ar.cfg.Server.Debug {
		tx = tx.Debug()
	}

	tx = tx.Table("movies").
		Select("EXTRACT(YEAR FROM movies.release_date) AS year, "+
			"EXTRACT(MONTH FROM movies.release_date) AS month, "+
			"COUNT(movies.id) AS num_movies, "+
			"SUM(COUNT(movies.id)) OVER (ORDER BY EXTRACT(YEAR FROM movies.release_date), EXTRACT(MONTH FROM movies.release_date)) AS cumulative").
		InnerJoins("INNER JOIN movies_genres on movies.id = movies_genres.movie_id").
		InnerJoins("INNER JOIN genres on genres.id = movies_genres.genre_id").
		Where("LOWER(genres.genre) = LOWER(?)", request.Genre)

	orBuild := ar.db
	for _, a := range request.Analysis {
		orBuild = orBuild.Or("EXTRACT(YEAR FROM movies.release_date) = ? AND EXTRACT(MONTH FROM movies.release_date) IN ?", a.Year, a.Months)
	}

	err := tx.Where(orBuild).
		Group("EXTRACT(YEAR FROM movies.release_date), EXTRACT(MONTH FROM movies.release_date)").
		Order("EXTRACT(YEAR FROM movies.release_date), EXTRACT(MONTH FROM movies.release_date)").
		Find(&result).Error

	if err != nil {
		return nil, err
	}

	return result, nil
}

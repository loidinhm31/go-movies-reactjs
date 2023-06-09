package repository

import (
	"context"
	"gorm.io/gorm"
	"movies-service/config"
	"movies-service/internal/analysis"
	"movies-service/internal/common/dto"
	"movies-service/internal/common/entity"
)

type analysisRepository struct {
	cfg *config.Config
	db  *gorm.DB
}

func NewAnalysisRepository(cfg *config.Config, db *gorm.DB) analysis.Repository {
	return &analysisRepository{cfg: cfg, db: db}
}

func (ar *analysisRepository) CountMoviesByGenre(ctx context.Context, movieType string) ([]*entity.GenreCount, error) {
	var result []*entity.GenreCount

	tx := ar.db.WithContext(ctx)
	if ar.cfg.Server.Debug {
		tx = tx.Debug()
	}

	if movieType != "" {
		err := tx.Raw("SELECT g.name, g.type_code, COUNT(mg.movie_id) AS num_movies "+
			"FROM genres g "+
			"INNER JOIN movies_genres mg on g.id = mg.genre_id "+
			"GROUP BY g.name, g.type_code "+
			"HAVING g.type_code = ? "+
			"ORDER BY g.name;", movieType).
			Scan(&result).Error
		if err != nil {
			return nil, err
		}
	} else {
		err := tx.Raw("SELECT g.name, g.type_code, COUNT(mg.movie_id) AS num_movies " +
			"FROM genres g " +
			"INNER JOIN movies_genres mg on g.id = mg.genre_id " +
			"GROUP BY g.name, g.type_code " +
			"ORDER BY g.name;").
			Scan(&result).Error
		if err != nil {
			return nil, err
		}
	}
	return result, nil
}

func (ar *analysisRepository) CountMoviesByReleaseDate(ctx context.Context, year string, months []string) ([]*entity.MovieCount, error) {
	var result []*entity.MovieCount

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

func (ar *analysisRepository) CountMoviesByCreatedDate(ctx context.Context, year string, months []string) ([]*entity.MovieCount, error) {
	var result []*entity.MovieCount

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

func (ar *analysisRepository) CountViewsByGenreAndViewedDate(ctx context.Context, request *dto.RequestData) ([]*entity.ViewCount, error) {
	var result []*entity.ViewCount

	tx := ar.db.WithContext(ctx)
	if ar.cfg.Server.Debug {
		tx = tx.Debug()
	}

	tx = tx.Table("views").
		Select("EXTRACT(YEAR FROM views.viewed_at) AS year, EXTRACT(MONTH FROM views.viewed_at) AS month, COUNT(views.id) AS num_viewers").
		InnerJoins("INNER JOIN movies on movies.id = views.movie_id AND movies.type_code = ? ", request.TypeCode).
		InnerJoins("INNER JOIN movies_genres on movies.id = movies_genres.movie_id").
		InnerJoins("INNER JOIN genres on genres.id = movies_genres.genre_id").
		Where("genres.name = ?", request.Name)

	orBuild := ar.db
	for _, a := range request.Analysis {
		orBuild = orBuild.Or("EXTRACT(YEAR FROM views.viewed_at) = ? AND EXTRACT(MONTH FROM views.viewed_at) IN ?", a.Year, a.Months)
	}

	err := tx.Group("EXTRACT(YEAR FROM views.viewed_at), EXTRACT(MONTH FROM views.viewed_at)").
		Having(orBuild).
		Order("EXTRACT(YEAR FROM views.viewed_at), EXTRACT(MONTH FROM views.viewed_at)").
		Find(&result).Error

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (ar *analysisRepository) CountCumulativeViewsByGenreAndViewedDate(ctx context.Context, request *dto.RequestData) ([]*entity.ViewCount, error) {
	var result []*entity.ViewCount

	tx := ar.db.WithContext(ctx)
	if ar.cfg.Server.Debug {
		tx = tx.Debug()
	}

	tx = tx.Table("views").
		Select("EXTRACT(YEAR FROM views.viewed_at) AS year, "+
			"EXTRACT(MONTH FROM views.viewed_at) AS month, "+
			"COUNT(views.id) AS num_viewers, "+
			"SUM(COUNT(views.id)) OVER (ORDER BY EXTRACT(YEAR FROM views.viewed_at), EXTRACT(MONTH FROM views.viewed_at)) AS cumulative").
		InnerJoins("INNER JOIN movies ON movies.id = views.movie_id AND movies.type_code = ?", request.TypeCode).
		InnerJoins("INNER JOIN movies_genres ON movies.id = movies_genres.movie_id").
		InnerJoins("INNER JOIN genres ON genres.id = movies_genres.genre_id AND genres.type_code = ?", request.TypeCode).
		Where("genres.name = ?", request.Name)

	orBuild := ar.db
	for _, a := range request.Analysis {
		orBuild = orBuild.Or("EXTRACT(YEAR FROM views.viewed_at) = ? AND EXTRACT(MONTH FROM views.viewed_at) IN ?", a.Year, a.Months)
	}

	err := tx.Group("EXTRACT(YEAR FROM views.viewed_at), EXTRACT(MONTH FROM views.viewed_at)").
		Having(orBuild).
		Order("EXTRACT(YEAR FROM views.viewed_at), EXTRACT(MONTH FROM views.viewed_at)").
		Find(&result).Error
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (ar *analysisRepository) CountViewsByViewedDate(ctx context.Context, request *dto.RequestData) ([]*entity.ViewCount, error) {
	var result []*entity.ViewCount

	tx := ar.db.WithContext(ctx)
	if ar.cfg.Server.Debug {
		tx = tx.Debug()
	}

	tx = tx.Table("views")

	if request.TypeCode == "" {
		tx = tx.Select("EXTRACT(YEAR FROM views.viewed_at) AS year, EXTRACT(MONTH FROM views.viewed_at) AS month, COUNT(views.id) AS num_viewers").
			Group("EXTRACT(YEAR FROM views.viewed_at), EXTRACT(MONTH FROM views.viewed_at)")
	} else {
		tx = tx.Joins("INNER JOIN movies ON movies.id = views.movie_id AND movies.type_code = ?", request.TypeCode).
			Select("EXTRACT(YEAR FROM views.viewed_at) AS year, EXTRACT(MONTH FROM views.viewed_at) AS month, COUNT(views.id) AS num_viewers").
			Group("EXTRACT(YEAR FROM views.viewed_at), EXTRACT(MONTH FROM views.viewed_at)")
	}

	buildHaving := ar.db
	for _, a := range request.Analysis {
		buildHaving = tx.Or("EXTRACT(YEAR FROM views.viewed_at) = ? AND EXTRACT(MONTH FROM views.viewed_at) IN ?", a.Year, a.Months)
	}

	err := tx.Having(buildHaving).
		Find(&result).Error

	if err != nil {
		return nil, err
	}
	return result, nil
}

func (ar *analysisRepository) CountMoviesByGenreAndReleasedDate(ctx context.Context, request *dto.RequestData) ([]*entity.MovieCount, error) {
	var result []*entity.MovieCount

	tx := ar.db.WithContext(ctx)
	if ar.cfg.Server.Debug {
		tx = tx.Debug()
	}

	tx = tx.Table("movies").
		Select("EXTRACT(YEAR FROM movies.release_date) AS year, "+
			"EXTRACT(MONTH FROM movies.release_date) AS month, "+
			"COUNT(movies.id) AS num_movies, "+
			"SUM(COUNT(movies.id)) OVER (ORDER BY EXTRACT(YEAR FROM movies.release_date), EXTRACT(MONTH FROM movies.release_date)) AS cumulative").
		InnerJoins("INNER JOIN movies_genres ON movies.id = movies_genres.movie_id").
		InnerJoins("INNER JOIN genres ON genres.id = movies_genres.genre_id").
		Where("genres.name = ? AND movies.type_code = ?", request.Name, request.TypeCode)

	orBuild := ar.db
	for _, a := range request.Analysis {
		orBuild = orBuild.Or("EXTRACT(YEAR FROM movies.release_date) = ? AND EXTRACT(MONTH FROM movies.release_date) IN ?", a.Year, a.Months)
	}

	err := tx.Group("EXTRACT(YEAR FROM movies.release_date), EXTRACT(MONTH FROM movies.release_date)").
		Having(orBuild).
		Order("EXTRACT(YEAR FROM movies.release_date), EXTRACT(MONTH FROM movies.release_date)").
		Find(&result).Error

	if err != nil {
		return nil, err
	}

	return result, nil
}

package repository

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"movies-service/config"
	"movies-service/internal/model"
	"movies-service/internal/movie"
	"movies-service/pkg/pagination"
	"strings"
)

type movieRepository struct {
	cfg *config.Config
	db  *gorm.DB
}

func NewMovieRepository(cfg *config.Config, db *gorm.DB) movie.Repository {
	return &movieRepository{cfg: cfg, db: db}
}

func (mr *movieRepository) InsertMovie(ctx context.Context, movie *model.Movie) error {
	tx := mr.db.WithContext(ctx)
	if mr.cfg.Server.Debug {
		tx = tx.Debug()
	}
	err := tx.Create(&movie).Error
	if err != nil {
		return err
	}
	return nil
}

func (mr *movieRepository) FindAllMovies(ctx context.Context, keyword string,
	pageRequest *pagination.PageRequest,
	page *pagination.Page[*model.Movie]) (*pagination.Page[*model.Movie], error) {
	var allMovies []*model.Movie

	tx := mr.db.WithContext(ctx)
	if mr.cfg.Server.Debug {
		tx = tx.Debug()
	}

	tx = tx.Model(allMovies)
	if keyword != "" {
		lowerWord := fmt.Sprintf("%%%s%%", strings.ToLower(keyword))
		tx = tx.Where("LOWER(title) LIKE ? OR LOWER(description) = ?", lowerWord, lowerWord)
	}
	err := tx.Preload("Genres").
		Scopes(pagination.PageImpl[*model.Movie](allMovies, pageRequest, page, mr.db)).
		Find(&allMovies).Error
	if err != nil {
		return nil, err
	}
	page.Content = allMovies
	return page, nil
}

func (mr *movieRepository) FindAllMoviesByType(ctx context.Context, keyword, movieType string, pageRequest *pagination.PageRequest, page *pagination.Page[*model.Movie]) (*pagination.Page[*model.Movie], error) {
	var allMovies []*model.Movie
	var totalRows int64

	tx := mr.db.WithContext(ctx)
	if mr.cfg.Server.Debug {
		tx = tx.Debug()
	}

	tx = tx.Model(allMovies)
	if keyword != "" {
		lowerWord := fmt.Sprintf("%%%s%%", strings.ToLower(keyword))
		tx = tx.Where("LOWER(title) LIKE ? OR LOWER(description) = ?", lowerWord, lowerWord)
	}

	err := tx.Where("type_code = ?", movieType).
		Count(&totalRows).
		Preload("Genres").
		Scopes(pagination.PageImplCountCriteria[*model.Movie](totalRows, pageRequest, page)).
		Find(&allMovies).Error
	if err != nil {
		return nil, err
	}
	page.Content = allMovies
	return page, nil
}

func (mr *movieRepository) FindMovieById(ctx context.Context, id uint) (*model.Movie, error) {
	var result model.Movie

	tx := mr.db.WithContext(ctx)
	if mr.cfg.Server.Debug {
		tx = tx.Debug()
	}
	err := tx.Preload("Genres").Where("id = ?", id).First(&result).Error
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (mr *movieRepository) FindMoviesByGenre(ctx context.Context,
	pageRequest *pagination.PageRequest,
	page *pagination.Page[*model.Movie],
	genreId uint) (*pagination.Page[*model.Movie], error) {
	var movieResults []*model.Movie
	var totalRows int64

	tx := mr.db.WithContext(ctx)
	if mr.cfg.Server.Debug {
		tx = tx.Debug()
	}

	err := tx.Model(movieResults).Where("movies.id IN (SELECT movie_id FROM movies_genres WHERE genre_id = ?)", genreId).
		Count(&totalRows).
		Scopes(pagination.PageImplCountCriteria[*model.Movie](totalRows, pageRequest, page)).
		Find(&movieResults).Error
	if err != nil {
		return nil, err
	}
	page.Content = movieResults
	return page, nil
}

func (mr *movieRepository) UpdateMovie(ctx context.Context, movie *model.Movie) error {
	tx := mr.db.WithContext(ctx)

	if mr.cfg.Server.Debug {
		tx = tx.Debug()
	}
	err := tx.Model(&model.Movie{}).Where("id = ?", movie.ID).Updates(movie).Error
	if err != nil {
		return err
	}
	return nil
}

func (mr *movieRepository) UpdateMovieGenres(ctx context.Context, movie *model.Movie, genres []*model.Genre) error {
	tx := mr.db.WithContext(ctx)
	if mr.cfg.Server.Debug {
		tx = tx.Debug()
	}
	err := tx.Model(&movie).Omit("Genres.*").
		Association("Genres").
		Replace(genres)
	if err != nil {
		return err
	}
	return nil
}

func (mr *movieRepository) DeleteMovieById(ctx context.Context, id uint) error {
	tx := mr.db.WithContext(ctx)
	if mr.cfg.Server.Debug {
		tx = tx.Debug()
	}
	err := tx.Select(clause.Associations).Delete(&model.Movie{
		ID: id,
	}).Error
	if err != nil {
		return err
	}
	return nil
}

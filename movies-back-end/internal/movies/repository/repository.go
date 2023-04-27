package repository

import (
	"context"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"movies-service/config"
	"movies-service/internal/models"
	"movies-service/internal/movies"
	"movies-service/pkg/pagination"
)

type movieRepository struct {
	cfg *config.Config
	db  *gorm.DB
}

func NewMovieRepository(cfg *config.Config, db *gorm.DB) movies.MovieRepository {
	return &movieRepository{cfg: cfg, db: db}
}

func (mr *movieRepository) InsertMovie(ctx context.Context, movie *models.Movie) error {
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

func (mr *movieRepository) FindAllMovies(ctx context.Context,
	pageRequest *pagination.PageRequest,
	page *pagination.Page[*models.Movie]) (*pagination.Page[*models.Movie], error) {
	var allMovies []*models.Movie

	tx := mr.db.WithContext(ctx)
	if mr.cfg.Server.Debug {
		tx = tx.Debug()
	}
	err := tx.Model(allMovies).Preload("Genres").
		Scopes(pagination.PageImpl[*models.Movie](allMovies, pageRequest, page, mr.db)).
		Find(&allMovies).Error
	if err != nil {
		return nil, err
	}
	page.Data = allMovies
	return page, nil
}

func (mr *movieRepository) FindMovieById(ctx context.Context, id int) (*models.Movie, error) {
	var movie models.Movie

	tx := mr.db.WithContext(ctx)
	if mr.cfg.Server.Debug {
		tx = tx.Debug()
	}
	err := tx.Preload("Genres").Where(&models.Movie{
		ID: id,
	}).First(&movie).Error
	if err != nil {
		return nil, err
	}

	return &movie, nil
}

func (mr *movieRepository) FindMoviesByGenre(ctx context.Context,
	pageRequest *pagination.PageRequest,
	page *pagination.Page[*models.Movie],
	genreId int) (*pagination.Page[*models.Movie], error) {
	var movieResults []*models.Movie
	var totalRows int64

	tx := mr.db.WithContext(ctx)
	if mr.cfg.Server.Debug {
		tx = tx.Debug()
	}

	err := tx.Model(movieResults).Where("movies.id IN (SELECT movie_id FROM movies_genres WHERE genre_id = ?)", genreId).
		Count(&totalRows).
		Scopes(pagination.PageImplCountCriteria[*models.Movie](totalRows, pageRequest, page, mr.db)).
		Find(&movieResults).Error
	if err != nil {
		return nil, err
	}
	page.Data = movieResults
	return page, nil
}

func (mr *movieRepository) UpdateMovie(ctx context.Context, movie *models.Movie) error {
	tx := mr.db.WithContext(ctx)

	if mr.cfg.Server.Debug {
		tx = tx.Debug()
	}
	err := tx.Model(&models.Movie{}).Where("id = ?", movie.ID).Updates(movie).Error
	if err != nil {
		return err
	}
	return nil
}

func (mr *movieRepository) UpdateMovieGenres(ctx context.Context, movie *models.Movie, genres []*models.Genre) error {
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

func (mr *movieRepository) DeleteMovieById(ctx context.Context, id int) error {
	tx := mr.db.WithContext(ctx)
	if mr.cfg.Server.Debug {
		tx = tx.Debug()
	}
	err := tx.Select(clause.Associations).Delete(&models.Movie{
		ID: id,
	}).Error
	if err != nil {
		return err
	}
	return nil
}

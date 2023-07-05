package repository

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"movies-service/config"
	"movies-service/internal/common/entity"
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

func (mr *movieRepository) InsertMovie(ctx context.Context, movie *entity.Movie) error {
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
	page *pagination.Page[*entity.Movie]) (*pagination.Page[*entity.Movie], error) {
	var allMovies []*entity.Movie

	tx := mr.db.WithContext(ctx)
	if mr.cfg.Server.Debug {
		tx = tx.Debug()
	}

	tx = tx.Model(allMovies)
	if keyword != "" {
		lowerWord := fmt.Sprintf("%%%s%%", strings.ToLower(keyword))
		tx = tx.Where("LOWER(title) LIKE ? OR LOWER(description) LIKE ?", lowerWord, lowerWord)
	}
	err := tx.Preload("Genres").
		Scopes(pagination.PageImpl[*entity.Movie](allMovies, pageRequest, page, mr.db)).
		Find(&allMovies).Error
	if err != nil {
		return nil, err
	}
	page.Content = allMovies
	return page, nil
}

func (mr *movieRepository) FindAllMoviesByType(ctx context.Context, keyword, movieType string, pageRequest *pagination.PageRequest, page *pagination.Page[*entity.Movie]) (*pagination.Page[*entity.Movie], error) {
	var allMovies []*entity.Movie
	var totalRows int64

	tx := mr.db.WithContext(ctx)
	if mr.cfg.Server.Debug {
		tx = tx.Debug()
	}

	tx = tx.Model(allMovies)
	if keyword != "" {
		lowerWord := fmt.Sprintf("%%%s%%", strings.ToLower(keyword))
		tx = tx.Where("LOWER(title) LIKE ? OR LOWER(description) LIKE ?", lowerWord, lowerWord)
	}

	err := tx.Where("type_code = ?", movieType).
		Count(&totalRows).
		Preload("Genres").
		Scopes(pagination.PageImplCountCriteria[*entity.Movie](totalRows, pageRequest, page)).
		Find(&allMovies).Error
	if err != nil {
		return nil, err
	}
	page.Content = allMovies
	return page, nil
}

func (mr *movieRepository) FindMovieByID(ctx context.Context, id uint) (*entity.Movie, error) {
	var result entity.Movie

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
	page *pagination.Page[*entity.Movie],
	genreId uint) (*pagination.Page[*entity.Movie], error) {
	var movieResults []*entity.Movie
	var totalRows int64

	tx := mr.db.WithContext(ctx)
	if mr.cfg.Server.Debug {
		tx = tx.Debug()
	}

	err := tx.Model(movieResults).Where("movies.id IN (SELECT movie_id FROM movies_genres WHERE genre_id = ?)", genreId).
		Count(&totalRows).
		Scopes(pagination.PageImplCountCriteria[*entity.Movie](totalRows, pageRequest, page)).
		Find(&movieResults).Error
	if err != nil {
		return nil, err
	}
	page.Content = movieResults
	return page, nil
}

func (mr *movieRepository) UpdateMovie(ctx context.Context, movie *entity.Movie) error {
	tx := mr.db.WithContext(ctx)

	if mr.cfg.Server.Debug {
		tx = tx.Debug()
	}
	err := tx.Model(&entity.Movie{}).Where("id = ?", movie.ID).Save(movie).Error
	if err != nil {
		return err
	}
	return nil
}

func (mr *movieRepository) UpdateMovieGenres(ctx context.Context, movie *entity.Movie, genres []*entity.Genre) error {
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

func (mr *movieRepository) DeleteMovieByID(ctx context.Context, id uint) error {
	tx := mr.db.WithContext(ctx)
	if mr.cfg.Server.Debug {
		tx = tx.Debug()
	}
	err := tx.Select(clause.Associations).Delete(&entity.Movie{
		ID: id,
	}).Error
	if err != nil {
		return err
	}
	return nil
}

func (mr *movieRepository) FindMovieByEpisodeID(ctx context.Context, episodeID uint) (*entity.Movie, error) {
	var result *entity.Movie

	tx := mr.db.WithContext(ctx)
	if mr.cfg.Server.Debug {
		tx = tx.Debug()
	}
	err := tx.Table("episodes e").
		Joins("JOIN seasons s ON s.id = e.season_id").
		Joins("JOIN movies m ON m.id = s.movie_id").
		Where("e.id = ?", episodeID).
		Select("m.*").
		Find(&result).Error
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (mr *movieRepository) UpdatePriceWithAverageEpisodePrice(ctx context.Context, movieID uint) error {
	tx := mr.db.WithContext(ctx)
	if mr.cfg.Server.Debug {
		tx = tx.Debug()
	}
	err := tx.Model(&entity.Movie{}).
		Where("id = ? AND type_code = 'TV'", movieID).
		Update("price", mr.db.Raw("SELECT AVG(e.price) "+
			"FROM seasons s "+
			"INNER JOIN episodes e ON s.id = e.season_id "+
			"WHERE s.movie_id = ?", movieID)).Error
	if err != nil {
		return err
	}

	return nil
}

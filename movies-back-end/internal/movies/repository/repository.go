package repository

import (
	"context"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"movies-service/internal/models"
	"movies-service/internal/movies"
)

type movieRepository struct {
	db *gorm.DB
}

func NewMovieRepository(db *gorm.DB) movies.MovieRepository {
	return &movieRepository{db: db}
}

func (mr *movieRepository) InsertMovie(ctx context.Context, movie *models.Movie) error {
	result := mr.db.WithContext(ctx).Create(&movie)

	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (mr *movieRepository) FindAllMovies(ctx context.Context) ([]*models.Movie, error) {
	var m []*models.Movie
	err := mr.db.WithContext(ctx).Preload("Genres").Find(&m).Error

	if err != nil {
		return nil, err
	}

	return m, nil
}

func (mr *movieRepository) FindMovieById(ctx context.Context, id int) (*models.Movie, error) {
	var movie models.Movie
	err := mr.db.WithContext(ctx).Preload("Genres").Where(&models.Movie{
		ID: id,
	}).First(&movie).Error

	if err != nil {
		return nil, err
	}

	return &movie, nil
}

func (mr *movieRepository) FindMoviesByGenre(ctx context.Context, genreId int) ([]*models.Movie, error) {
	var m []*models.Movie

	err := mr.db.WithContext(ctx).Where("movies.id IN (SELECT movie_id FROM movies_genres WHERE genre_id = ?)", genreId).
		Find(&m).Error
	if err != nil {
		return nil, err
	}

	return m, nil
}

func (mr *movieRepository) UpdateMovie(ctx context.Context, movie *models.Movie) error {
	err := mr.db.WithContext(ctx).Model(&models.Movie{}).Where("id = ?", movie.ID).Updates(movie).Error

	if err != nil {
		return err
	}
	return nil
}

func (mr *movieRepository) UpdateMovieGenres(ctx context.Context, movie *models.Movie, genres []*models.Genre) error {
	err := mr.db.WithContext(ctx).Model(&movie).Omit("Genres.*").Association("Genres").Replace(genres)
	if err != nil {
		return err
	}
	return nil
}

func (mr *movieRepository) DeleteMovieById(ctx context.Context, id int) error {
	err := mr.db.WithContext(ctx).Select(clause.Associations).Delete(&models.Movie{
		ID: id,
	}).Error
	if err != nil {
		return err
	}
	return nil
}

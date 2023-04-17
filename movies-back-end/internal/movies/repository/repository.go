package repository

import (
	"context"
	"gorm.io/gorm"
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
	var movies []*models.Movie
	err := mr.db.WithContext(ctx).Preload("Genres").Find(&movies).Error

	if err != nil {
		return nil, err
	}

	return movies, nil
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
	var movies []*models.Movie

	err := mr.db.WithContext(ctx).Where("movies.id IN (SELECT movie_id FROM movies_genres WHERE genre_id = ?)", genreId).
		Find(&movies).Error
	if err != nil {
		return nil, err
	}

	return movies, nil
}

func (mr *movieRepository) UpdateMovie(ctx context.Context, movie models.Movie) error {
	err := mr.db.WithContext(ctx).Updates(movie).Error

	if err != nil {
		return err
	}
	return nil
}

//func (mr *movieRepository) UpdateMovieGenres(id int, genreIDs []int) error {
//
//
//	stmt := `DELETE FROM movies_genres WHERE movie_id = $1`
//
//	_, err := gr.DB.ExecContext(ctx, stmt, id)
//	if err != nil {
//		return err
//	}
//
//	for _, n := range genreIDs {
//		stmt := `INSERT INTO movies_genres (movie_id, genre_id) VALUES ($1, $2)`
//		_, err := gr.DB.ExecContext(ctx, stmt, id, n)
//		if err != nil {
//			return err
//		}
//	}
//	return nil
//}

func (mr *movieRepository) DeleteMovieById(ctx context.Context, id int) error {
	err := mr.db.WithContext(ctx).Delete(&models.Movie{}, id).Error
	if err != nil {
		return err
	}
	return nil
}

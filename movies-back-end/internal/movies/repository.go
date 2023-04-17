package movies

import (
	"context"
	"movies-service/internal/models"
)

type MovieRepository interface {
	InsertMovie(ctx context.Context, user *models.Movie) error
	FindAllMovies(ctx context.Context) ([]*models.Movie, error)
	FindMovieById(ctx context.Context, id int) (*models.Movie, error)
	FindMoviesByGenre(ctx context.Context, genreId int) ([]*models.Movie, error)
	UpdateMovie(ctx context.Context, movie models.Movie) error
	DeleteMovieById(ctx context.Context, id int) error
}

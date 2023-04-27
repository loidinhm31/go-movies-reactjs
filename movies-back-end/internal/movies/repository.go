package movies

import (
	"context"
	"movies-service/internal/models"
	"movies-service/pkg/pagination"
)

type MovieRepository interface {
	InsertMovie(ctx context.Context, user *models.Movie) error
	FindAllMovies(ctx context.Context, pageRequest *pagination.PageRequest,
		page *pagination.Page[*models.Movie]) (*pagination.Page[*models.Movie], error)
	FindMovieById(ctx context.Context, id int) (*models.Movie, error)
	FindMoviesByGenre(ctx context.Context, pageRequest *pagination.PageRequest, page *pagination.Page[*models.Movie], genreId int) (*pagination.Page[*models.Movie], error)
	UpdateMovie(ctx context.Context, movie *models.Movie) error
	DeleteMovieById(ctx context.Context, id int) error
	UpdateMovieGenres(ctx context.Context, movie *models.Movie, genres []*models.Genre) error
}

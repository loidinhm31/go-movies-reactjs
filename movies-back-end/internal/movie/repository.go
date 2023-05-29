package movie

import (
	"context"
	"movies-service/internal/model"
	"movies-service/pkg/pagination"
)

type Repository interface {
	InsertMovie(ctx context.Context, movie *model.Movie) error

	FindAllMovies(ctx context.Context, keyword string, pageRequest *pagination.PageRequest,
		page *pagination.Page[*model.Movie]) (*pagination.Page[*model.Movie], error)

	FindAllMoviesByType(ctx context.Context, keyword, movieType string, pageRequest *pagination.PageRequest, page *pagination.Page[*model.Movie]) (*pagination.Page[*model.Movie], error)

	FindMovieById(ctx context.Context, id int) (*model.Movie, error)

	FindMoviesByGenre(ctx context.Context, pageRequest *pagination.PageRequest, page *pagination.Page[*model.Movie], genreId int) (*pagination.Page[*model.Movie], error)

	UpdateMovie(ctx context.Context, movie *model.Movie) error

	DeleteMovieById(ctx context.Context, id int) error

	UpdateMovieGenres(ctx context.Context, movie *model.Movie, genres []*model.Genre) error
}

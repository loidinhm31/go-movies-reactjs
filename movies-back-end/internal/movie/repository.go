package movie

import (
	"context"
	model2 "movies-service/internal/common/model"
	"movies-service/pkg/pagination"
)

type Repository interface {
	InsertMovie(ctx context.Context, movie *model2.Movie) error

	FindAllMovies(ctx context.Context, keyword string, pageRequest *pagination.PageRequest,
		page *pagination.Page[*model2.Movie]) (*pagination.Page[*model2.Movie], error)

	FindAllMoviesByType(ctx context.Context, keyword, movieType string, pageRequest *pagination.PageRequest, page *pagination.Page[*model2.Movie]) (*pagination.Page[*model2.Movie], error)

	FindMovieByID(ctx context.Context, id uint) (*model2.Movie, error)

	FindMoviesByGenre(ctx context.Context, pageRequest *pagination.PageRequest, page *pagination.Page[*model2.Movie], genreId uint) (*pagination.Page[*model2.Movie], error)

	UpdateMovie(ctx context.Context, movie *model2.Movie) error

	DeleteMovieByID(ctx context.Context, id uint) error

	UpdateMovieGenres(ctx context.Context, movie *model2.Movie, genres []*model2.Genre) error

	FindMovieByEpisodeID(ctx context.Context, episdoeID uint) (*model2.Movie, error)
}

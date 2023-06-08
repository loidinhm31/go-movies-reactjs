package movie

import (
	"context"
	"movies-service/internal/common/entity"
	"movies-service/pkg/pagination"
)

type Repository interface {
	InsertMovie(ctx context.Context, movie *entity.Movie) error

	FindAllMovies(ctx context.Context, keyword string, pageRequest *pagination.PageRequest,
		page *pagination.Page[*entity.Movie]) (*pagination.Page[*entity.Movie], error)

	FindAllMoviesByType(ctx context.Context, keyword, movieType string, pageRequest *pagination.PageRequest, page *pagination.Page[*entity.Movie]) (*pagination.Page[*entity.Movie], error)

	FindMovieByID(ctx context.Context, id uint) (*entity.Movie, error)

	FindMoviesByGenre(ctx context.Context, pageRequest *pagination.PageRequest, page *pagination.Page[*entity.Movie], genreId uint) (*pagination.Page[*entity.Movie], error)

	UpdateMovie(ctx context.Context, movie *entity.Movie) error

	DeleteMovieByID(ctx context.Context, id uint) error

	UpdateMovieGenres(ctx context.Context, movie *entity.Movie, genres []*entity.Genre) error

	FindMovieByEpisodeID(ctx context.Context, episdoeID uint) (*entity.Movie, error)
}

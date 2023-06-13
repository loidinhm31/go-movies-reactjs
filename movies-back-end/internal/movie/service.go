package movie

import (
	"context"
	"movies-service/internal/common/dto"
	"movies-service/pkg/pagination"
)

type Service interface {
	GetAllMoviesByType(ctx context.Context, keyword, movieType string, pageRequest *pagination.PageRequest) (*pagination.Page[*dto.MovieDto], error)
	GetMovieByID(ctx context.Context, id uint) (*dto.MovieDto, error)
	GetMoviesByGenre(ctx context.Context, pageRequest *pagination.PageRequest, genreId uint) (*pagination.Page[*dto.MovieDto], error)
	AddMovie(ctx context.Context, movie *dto.MovieDto) error
	UpdateMovie(ctx context.Context, movie *dto.MovieDto) error
	RemoveMovieByID(ctx context.Context, id uint) error
	GetMovieByEpisodeID(ctx context.Context, id uint) (*dto.MovieDto, error)
	UpdatePriceWithAverageEpisodePrice(ctx context.Context, movieID uint) error
}

package movies

import (
	"context"
	"movies-service/internal/dto"
	"movies-service/pkg/pagination"
)

type Service interface {
	GetAllMovies(ctx context.Context, pageRequest *pagination.PageRequest) (*pagination.Page[*dto.MovieDto], error)
	GetMovieById(ctx context.Context, id int) (*dto.MovieDto, error)
	GetMoviesByGenre(ctx context.Context, pageRequest *pagination.PageRequest, genreId int) (*pagination.Page[*dto.MovieDto], error)
	AddMovie(ctx context.Context, movie *dto.MovieDto) error
	UpdateMovie(ctx context.Context, movie *dto.MovieDto) error
	DeleteMovieById(ctx context.Context, id int) error
}

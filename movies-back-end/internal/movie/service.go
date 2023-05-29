package movie

import (
	"context"
	"movies-service/internal/dto"
	"movies-service/pkg/pagination"
)

type Service interface {
	GetAllMoviesByType(ctx context.Context, keyword, movieType string, pageRequest *pagination.PageRequest) (*pagination.Page[*dto.MovieDto], error)
	GetMovieById(ctx context.Context, id int) (*dto.MovieDto, error)
	GetMoviesByGenre(ctx context.Context, pageRequest *pagination.PageRequest, genreId int) (*pagination.Page[*dto.MovieDto], error)
	AddMovie(ctx context.Context, movie *dto.MovieDto) error
	UpdateMovie(ctx context.Context, movie *dto.MovieDto) error
	DeleteMovieById(ctx context.Context, id int) error
}
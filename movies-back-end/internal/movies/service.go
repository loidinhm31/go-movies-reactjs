package movies

import (
	"context"
	"movies-service/internal/dto"
)

type Service interface {
	GetAllMovies(ctx context.Context) ([]*dto.MovieDto, error)
	GetMovieById(ctx context.Context, id int) (*dto.MovieDto, error)
	GetMoviesByGenre(ctx context.Context, genreId int) ([]*dto.MovieDto, error)
	UpdateMovie(ctx context.Context, movie *dto.MovieDto) error
	DeleteMovieById(ctx context.Context, id int) error
}

package reference

import (
	"context"
	"movies-service/internal/common/dto"
)

type Service interface {
	GetMoviesByType(ctx context.Context, movie *dto.MovieDto) ([]*dto.MovieDto, error)
	GetMovieById(ctx context.Context, movieId int64, movieType string) (*dto.MovieDto, error)
}

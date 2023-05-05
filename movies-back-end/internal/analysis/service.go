package analysis

import (
	"context"
	"movies-service/internal/dto"
)

type Service interface {
	GetNumberOfMoviesByGenre(ctx context.Context) (*dto.ResultDto, error)
	GetNumberOfMoviesByReleaseDate(ctx context.Context, year string, months []string) (*dto.ResultDto, error)
	GetNumberOfMoviesByCreatedDate(ctx context.Context, year string, months []string) (*dto.ResultDto, error)
	GetNumberOfViewsByGenreAndViewedDate(ctx context.Context, request *dto.RequestData) (*dto.ResultDto, error)
	GetNumberOfViewsByViewedDate(ctx context.Context, request *dto.RequestData) (*dto.ResultDto, error)
}

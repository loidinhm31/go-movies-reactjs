package analysis

import (
	"context"
	"movies-service/internal/common/dto"
	"movies-service/internal/common/model"
)

type Repository interface {
	CountMoviesByGenre(ctx context.Context, movieType string) ([]*model.GenreCount, error)
	CountMoviesByReleaseDate(ctx context.Context, year string, months []string) ([]*model.MovieCount, error)
	CountMoviesByCreatedDate(ctx context.Context, year string, months []string) ([]*model.MovieCount, error)
	CountViewsByGenreAndViewedDate(ctx context.Context, request *dto.RequestData) ([]*model.ViewCount, error)
	CountCumulativeViewsByGenreAndViewedDate(ctx context.Context, request *dto.RequestData) ([]*model.ViewCount, error)
	CountViewsByViewedDate(ctx context.Context, request *dto.RequestData) ([]*model.ViewCount, error)
	CountMoviesByGenreAndReleasedDate(ctx context.Context, request *dto.RequestData) ([]*model.MovieCount, error)
}

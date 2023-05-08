package analysis

import (
	"context"
	"movies-service/internal/dto"
	"movies-service/internal/models"
)

type Repository interface {
	CountMoviesByGenre(ctx context.Context) ([]*models.GenreCount, error)
	CountMoviesByReleaseDate(ctx context.Context, year string, months []string) ([]*models.MovieCount, error)
	CountMoviesByCreatedDate(ctx context.Context, year string, months []string) ([]*models.MovieCount, error)
	CountViewsByGenreAndViewedDate(ctx context.Context, request *dto.RequestData) ([]*models.ViewCount, error)
	CountCumulativeViewsByGenreAndViewedDate(ctx context.Context, request *dto.RequestData) ([]*models.ViewCount, error)
	CountViewsByViewedDate(ctx context.Context, request *dto.RequestData) ([]*models.ViewCount, error)
	CountMoviesByGenreAndReleasedDate(ctx context.Context, request *dto.RequestData) ([]*models.MovieCount, error)
}

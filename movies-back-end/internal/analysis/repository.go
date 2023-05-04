package analysis

import (
	"context"
	"movies-service/internal/models"
)

type Repository interface {
	CountMoviesByGenre(ctx context.Context) ([]*models.GenreCount, error)
	CountMoviesByReleaseDate(ctx context.Context, year string, months []string) ([]*models.MovieCount, error)
	CountMoviesByCreatedDate(ctx context.Context, year string, months []string) ([]*models.MovieCount, error)
	CountViewsByGenreAndViewedDate(ctx context.Context, genre, year string, months []string) ([]*models.ViewCount, error)
}

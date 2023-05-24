package genres

import (
	"context"
	"movies-service/internal/models"
)

type GenreRepository interface {
	FindAllGenres(ctx context.Context) ([]*models.Genre, error)
	FindAllGenresByTypeCode(ctx context.Context, movieType string) ([]*models.Genre, error)
	FindGenreByNameAndTypeCode(ctx context.Context, genre *models.Genre) (*models.Genre, error)
	InsertGenres(ctx context.Context, genres []*models.Genre) error
}

package genres

import (
	"context"
	"movies-service/internal/models"
)

type GenreRepository interface {
	FindAllGenres(ctx context.Context) ([]*models.Genre, error)
}

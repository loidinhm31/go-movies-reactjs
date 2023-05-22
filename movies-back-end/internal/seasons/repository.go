package seasons

import (
	"context"
	"movies-service/internal/models"
)

type Repository interface {
	FindSeasonByID(ctx context.Context, id int) (*models.Season, error)

	FindSeasonsByMovieID(ctx context.Context, movieID int) ([]*models.Season, error)

	InsertSeason(ctx context.Context, season *models.Season) error

	UpdateSeason(ctx context.Context, season *models.Season) error

	DeleteSeasonByID(ctx context.Context, id int) error
}

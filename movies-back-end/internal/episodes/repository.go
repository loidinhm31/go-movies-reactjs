package episodes

import (
	"context"
	"movies-service/internal/models"
)

type Repository interface {
	FindEpisodeByID(ctx context.Context, id int) (*models.Episode, error)

	FindEpisodesBySeasonID(ctx context.Context, movieID int) ([]*models.Episode, error)

	InsertEpisode(ctx context.Context, episode *models.Episode) error

	UpdateEpisode(ctx context.Context, episode *models.Episode) error

	DeleteEpisodeById(ctx context.Context, id int) error

	DeleteEpisodeBySeasonID(ctx context.Context, seasonID int) error
}

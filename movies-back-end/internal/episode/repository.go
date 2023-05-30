package episode

import (
	"context"
	"movies-service/internal/model"
)

type Repository interface {
	FindEpisodeByID(ctx context.Context, id int) (*model.Episode, error)

	FindEpisodesBySeasonID(ctx context.Context, movieID int) ([]*model.Episode, error)

	InsertEpisode(ctx context.Context, episode *model.Episode) error

	UpdateEpisode(ctx context.Context, episode *model.Episode) error

	DeleteEpisodeByID(ctx context.Context, id int) error

	DeleteEpisodeBySeasonID(ctx context.Context, seasonID int) error
}

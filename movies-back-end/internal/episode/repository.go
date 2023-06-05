package episode

import (
	"context"
	"movies-service/internal/common/model"
)

type Repository interface {
	FindEpisodeByID(ctx context.Context, id uint) (*model.Episode, error)

	FindEpisodesBySeasonID(ctx context.Context, movieID uint) ([]*model.Episode, error)

	InsertEpisode(ctx context.Context, episode *model.Episode) error

	UpdateEpisode(ctx context.Context, episode *model.Episode) error

	DeleteEpisodeByID(ctx context.Context, id uint) error

	DeleteEpisodeBySeasonID(ctx context.Context, seasonID uint) error
}

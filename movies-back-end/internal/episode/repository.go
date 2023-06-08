package episode

import (
	"context"
	"movies-service/internal/common/entity"
)

type Repository interface {
	FindEpisodeByID(ctx context.Context, id uint) (*entity.Episode, error)

	FindEpisodesBySeasonID(ctx context.Context, movieID uint) ([]*entity.Episode, error)

	InsertEpisode(ctx context.Context, episode *entity.Episode) error

	UpdateEpisode(ctx context.Context, episode *entity.Episode) error

	DeleteEpisodeByID(ctx context.Context, id uint) error

	DeleteEpisodeBySeasonID(ctx context.Context, seasonID uint) error
}

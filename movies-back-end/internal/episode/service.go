package episode

import (
	"context"
	"movies-service/internal/dto"
)

type Service interface {
	GetEpisodesByID(ctx context.Context, id int) (*dto.EpisodeDto, error)
	GetEpisodesBySeasonID(ctx context.Context, seasonID int) ([]*dto.EpisodeDto, error)
	AddEpisode(ctx context.Context, episode *dto.EpisodeDto) error
	UpdateEpisode(ctx context.Context, episode *dto.EpisodeDto) error
	RemoveEpisodeByID(ctx context.Context, id int) error
}

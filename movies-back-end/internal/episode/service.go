package episode

import (
	"context"
	"movies-service/internal/common/dto"
)

type Service interface {
	GetEpisodeByID(ctx context.Context, id uint) (*dto.EpisodeDto, error)
	GetEpisodesBySeasonID(ctx context.Context, seasonID uint) ([]*dto.EpisodeDto, error)
	AddEpisode(ctx context.Context, episode *dto.EpisodeDto) error
	UpdateEpisode(ctx context.Context, episode *dto.EpisodeDto) error
	RemoveEpisodeByID(ctx context.Context, id uint) error
}

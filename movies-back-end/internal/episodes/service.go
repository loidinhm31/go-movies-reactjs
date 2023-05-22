package episodes

import (
	"context"
	"github.com/gin-gonic/gin"
	"movies-service/internal/dto"
)

type Service interface {
	GetEpisodesByID(ctx *gin.Context, id int) (*dto.EpisodeDto, error)
	GetEpisodesBySeasonID(ctx context.Context, seasonID int) ([]*dto.EpisodeDto, error)
	AddEpisode(ctx context.Context, episode *dto.EpisodeDto) error
	UpdateEpisode(ctx context.Context, episode *dto.EpisodeDto) error
	DeleteEpisodeById(ctx context.Context, id int) error
}

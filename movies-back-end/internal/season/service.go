package season

import (
	"context"
	"github.com/gin-gonic/gin"
	"movies-service/internal/dto"
)

type Service interface {
	GetSeasonsByID(ctx *gin.Context, id uint) (*dto.SeasonDto, error)
	GetSeasonsByMovieID(ctx context.Context, movieID uint) ([]*dto.SeasonDto, error)
	AddSeason(ctx context.Context, season *dto.SeasonDto) error
	UpdateSeason(ctx context.Context, season *dto.SeasonDto) error
	RemoveSeasonByID(ctx context.Context, id uint) error
}

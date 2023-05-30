package season

import (
	"context"
	"github.com/gin-gonic/gin"
	"movies-service/internal/dto"
)

type Service interface {
	GetSeasonsByID(ctx *gin.Context, id int) (*dto.SeasonDto, error)
	GetSeasonsByMovieID(ctx context.Context, movieID int) ([]*dto.SeasonDto, error)
	AddSeason(ctx context.Context, season *dto.SeasonDto) error
	UpdateSeason(ctx context.Context, season *dto.SeasonDto) error
	RemoveSeasonByID(ctx context.Context, id int) error
}

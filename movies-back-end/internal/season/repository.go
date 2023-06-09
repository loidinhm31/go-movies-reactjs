package season

import (
	"context"
	"movies-service/internal/common/entity"
)

type Repository interface {
	FindSeasonByID(ctx context.Context, id uint) (*entity.Season, error)

	FindSeasonsByMovieID(ctx context.Context, movieID uint) ([]*entity.Season, error)

	InsertSeason(ctx context.Context, season *entity.Season) error

	UpdateSeason(ctx context.Context, season *entity.Season) error

	DeleteSeasonByID(ctx context.Context, id uint) error
}

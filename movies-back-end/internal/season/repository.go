package season

import (
	"context"
	"movies-service/internal/model"
)

type Repository interface {
	FindSeasonByID(ctx context.Context, id int) (*model.Season, error)

	FindSeasonsByMovieID(ctx context.Context, movieID int) ([]*model.Season, error)

	InsertSeason(ctx context.Context, season *model.Season) error

	UpdateSeason(ctx context.Context, season *model.Season) error

	DeleteSeasonByID(ctx context.Context, id int) error
}

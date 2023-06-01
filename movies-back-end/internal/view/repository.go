package view

import (
	"context"
	"movies-service/internal/dto"
)

type Repository interface {
	InsertView(ctx context.Context, view *dto.Viewer) error
	CountViewsByMovieId(ctx context.Context, movieId uint) (int64, error)
}

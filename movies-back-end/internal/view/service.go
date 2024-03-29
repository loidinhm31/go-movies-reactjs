package view

import (
	"context"
	"movies-service/internal/common/dto"
)

type Service interface {
	RecognizeViewForMovie(ctx context.Context, viewer *dto.Viewer) error
	GetNumberOfViewsByMovieId(ctx context.Context, movieId uint) (int64, error)
}

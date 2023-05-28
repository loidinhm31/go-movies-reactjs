package view

import (
	"context"
	"movies-service/internal/dto"
)

type Service interface {
	RecognizeViewForMovie(ctx context.Context, viewer *dto.Viewer) error
	GetNumberOfViewsByMovieId(ctx context.Context, movieId int) (int64, error)
}

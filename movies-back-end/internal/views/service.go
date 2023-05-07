package views

import (
	"context"
	"movies-service/internal/dto"
)

type Service interface {
	RecognizeViewForMovie(ctx context.Context, viewer *dto.Viewer) error
}

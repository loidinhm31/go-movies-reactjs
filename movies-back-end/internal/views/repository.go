package views

import (
	"context"
	"movies-service/internal/dto"
)

type Repository interface {
	InsertView(ctx context.Context, view *dto.Viewer) error
}

package search

import (
	"context"
	"movies-service/internal/models"
	"movies-service/pkg/pagination"
)

type Repository interface {
	Search(ctx context.Context, searchParams *models.SearchParams) (*pagination.Page[*models.Movie], error)
}

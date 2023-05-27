package search

import (
	"context"
	"movies-service/internal/model"
	"movies-service/pkg/pagination"
)

type Repository interface {
	SearchMovie(ctx context.Context, searchParams *model.SearchParams) (*pagination.Page[*model.Movie], error)
}

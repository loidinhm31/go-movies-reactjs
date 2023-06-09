package search

import (
	"context"
	"movies-service/internal/common/entity"
	model2 "movies-service/internal/common/model"
	"movies-service/pkg/pagination"
)

type Repository interface {
	SearchMovie(ctx context.Context, searchParams *model2.SearchParams) (*pagination.Page[*entity.Movie], error)
}

package search

import (
	"context"
	"movies-service/internal/dto"
	"movies-service/internal/model"
	"movies-service/pkg/pagination"
)

type Service interface {
	SearchMovie(ctx context.Context, searchParams *model.SearchParams) (*pagination.Page[*dto.MovieDto], error)
}

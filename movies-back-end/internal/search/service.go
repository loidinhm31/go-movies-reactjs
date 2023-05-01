package search

import (
	"context"
	"movies-service/internal/dto"
	"movies-service/internal/models"
	"movies-service/pkg/pagination"
)

type Service interface {
	Search(ctx context.Context, searchParams *models.SearchParams) (*pagination.Page[*dto.MovieDto], error)
}

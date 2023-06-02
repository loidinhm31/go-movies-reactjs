package collection

import (
	"context"
	"movies-service/internal/dto"
	"movies-service/pkg/pagination"
)

type Service interface {
	AddCollection(ctx context.Context, movieID uint) error
	GetCollectionsByUsernameAndType(ctx context.Context, movieType string, keyword string, pageRequest *pagination.PageRequest) (*pagination.Page[*dto.CollectionDto], error)
	GetCollectionByUsernameAndMovieID(ctx context.Context, username string, movieID uint) (*dto.CollectionDto, error)
}

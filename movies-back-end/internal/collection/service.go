package collection

import (
	"context"
	"movies-service/internal/dto"
	"movies-service/pkg/pagination"
)

type Service interface {
	AddCollection(ctx context.Context, collection *dto.CollectionDto) error
	GetCollectionsByUsernameAndType(ctx context.Context, movieType string, keyword string, pageRequest *pagination.PageRequest) (*pagination.Page[*dto.CollectionDetailDto], error)
	GetCollectionByUsernameAndRefID(ctx context.Context, username string, typeCode string, refID uint) (*dto.CollectionDto, error)
}

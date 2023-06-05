package collection

import (
	"context"
	"movies-service/internal/common/dto"
	"movies-service/pkg/pagination"
)

type Service interface {
	AddCollection(ctx context.Context, collection *dto.CollectionDto) error
	GetCollectionsByUserAndType(ctx context.Context, movieType string, keyword string, pageRequest *pagination.PageRequest) (*pagination.Page[*dto.CollectionDetailDto], error)
	GetCollectionByUserAndRefID(ctx context.Context, typeCode string, refID uint) (*dto.CollectionDto, error)
	RemoveCollectionByRefID(ctx context.Context, typeCode string, refID uint) error
}

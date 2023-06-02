package collection

import (
	"context"
	"movies-service/internal/model"
	"movies-service/pkg/pagination"
)

type Repository interface {
	InsertCollection(ctx context.Context, collection *model.Collection) error
	FindCollectionsByUsernameAndType(ctx context.Context, username string, movieType string, keyword string, pageRequest *pagination.PageRequest, page *pagination.Page[*model.CollectionDetail]) (*pagination.Page[*model.CollectionDetail], error)
	FindCollectionByUsernameAndMovieID(ctx context.Context, username string, movieID uint) (*model.Collection, error)
	FindCollectionByPaymentID(ctx context.Context, paymentID uint) (*model.Collection, error)
	FindCollectionsByMovieID(ctx context.Context, movieID uint) ([]*model.Collection, error)
}

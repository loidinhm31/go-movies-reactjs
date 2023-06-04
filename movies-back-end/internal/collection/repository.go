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
	FindCollectionByUsernameAndEpisodeID(ctx context.Context, username string, episodeID uint) (*model.Collection, error)
	FindCollectionByPaymentID(ctx context.Context, paymentID uint) (*model.Collection, error)
	FindCollectionsByMovieID(ctx context.Context, movieID uint) ([]*model.Collection, error)
	FindCollectionByMovieID(ctx context.Context, movieID uint) (*model.Collection, error)
	FindCollectionsByEpisodeID(ctx context.Context, episodeID uint) ([]*model.Collection, error)
	FindCollectionByEpisodeID(ctx context.Context, episodeID uint) (*model.Collection, error)
	FindCollectionsByID(ctx context.Context, id uint) (*model.Collection, error)
}

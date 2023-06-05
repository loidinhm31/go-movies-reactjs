package collection

import (
	"context"
	"movies-service/internal/common/model"
	"movies-service/pkg/pagination"
)

type Repository interface {
	InsertCollection(ctx context.Context, collection *model.Collection) error
	FindCollectionsByUserIDAndType(ctx context.Context, userID uint, movieType string, keyword string, pageRequest *pagination.PageRequest, page *pagination.Page[*model.CollectionDetail]) (*pagination.Page[*model.CollectionDetail], error)
	FindCollectionByUserIDAndMovieID(ctx context.Context, userID uint, movieID uint) (*model.Collection, error)
	FindCollectionByUserIDAndEpisodeID(ctx context.Context, userID uint, episodeID uint) (*model.Collection, error)
	FindCollectionByPaymentID(ctx context.Context, paymentID uint) (*model.Collection, error)
	FindCollectionsByMovieID(ctx context.Context, movieID uint) ([]*model.Collection, error)
	FindCollectionByMovieID(ctx context.Context, movieID uint) (*model.Collection, error)
	FindCollectionsByEpisodeID(ctx context.Context, episodeID uint) ([]*model.Collection, error)
	FindCollectionByEpisodeID(ctx context.Context, episodeID uint) (*model.Collection, error)
	FindCollectionsByID(ctx context.Context, id uint) (*model.Collection, error)
	DeleteCollectionByTypeCodeAndMovieID(ctx context.Context, typeCode string, movieID uint) error
	DeleteCollectionByTypeCodeAndEpisodeID(ctx context.Context, typeCode string, episodeID uint) error
}

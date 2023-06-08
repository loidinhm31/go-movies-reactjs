package collection

import (
	"context"
	"movies-service/internal/common/entity"
	"movies-service/pkg/pagination"
)

type Repository interface {
	InsertCollection(ctx context.Context, collection *entity.Collection) error
	FindCollectionsByUserIDAndType(ctx context.Context, userID uint, movieType string, keyword string, pageRequest *pagination.PageRequest, page *pagination.Page[*entity.CollectionDetail]) (*pagination.Page[*entity.CollectionDetail], error)
	FindCollectionByUserIDAndMovieID(ctx context.Context, userID uint, movieID uint) (*entity.Collection, error)
	FindCollectionByUserIDAndEpisodeID(ctx context.Context, userID uint, episodeID uint) (*entity.Collection, error)
	FindCollectionByPaymentID(ctx context.Context, paymentID uint) (*entity.Collection, error)
	FindCollectionsByMovieID(ctx context.Context, movieID uint) ([]*entity.Collection, error)
	FindCollectionByMovieID(ctx context.Context, movieID uint) (*entity.Collection, error)
	FindCollectionsByEpisodeID(ctx context.Context, episodeID uint) ([]*entity.Collection, error)
	FindCollectionByEpisodeID(ctx context.Context, episodeID uint) (*entity.Collection, error)
	FindCollectionsByID(ctx context.Context, id uint) (*entity.Collection, error)
	DeleteCollectionByTypeCodeAndMovieID(ctx context.Context, typeCode string, movieID uint) error
	DeleteCollectionByTypeCodeAndEpisodeID(ctx context.Context, typeCode string, episodeID uint) error
}

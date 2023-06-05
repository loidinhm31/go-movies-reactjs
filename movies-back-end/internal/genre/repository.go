package genre

import (
	"context"
	"movies-service/internal/common/model"
)

type Repository interface {
	FindAllGenres(ctx context.Context) ([]*model.Genre, error)
	FindAllGenresByTypeCode(ctx context.Context, movieType string) ([]*model.Genre, error)
	FindGenreByNameAndTypeCode(ctx context.Context, genre *model.Genre) (*model.Genre, error)
	InsertGenres(ctx context.Context, genres []*model.Genre) error
}

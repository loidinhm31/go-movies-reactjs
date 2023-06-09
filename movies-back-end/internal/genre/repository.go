package genre

import (
	"context"
	"movies-service/internal/common/entity"
)

type Repository interface {
	FindAllGenres(ctx context.Context) ([]*entity.Genre, error)
	FindAllGenresByTypeCode(ctx context.Context, movieType string) ([]*entity.Genre, error)
	FindGenreByNameAndTypeCode(ctx context.Context, genre *entity.Genre) (*entity.Genre, error)
	InsertGenres(ctx context.Context, genres []*entity.Genre) error
}

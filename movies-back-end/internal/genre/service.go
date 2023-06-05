package genre

import (
	"context"
	"movies-service/internal/common/dto"
)

type Service interface {
	GetAllGenresByTypeCode(ctx context.Context, movieType string) ([]*dto.GenreDto, error)
	AddGenres(ctx context.Context, genreNames []dto.GenreDto) error
}

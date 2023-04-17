package genres

import (
	"context"
	"movies-service/internal/dto"
)

type Service interface {
	GetAllGenres(ctx context.Context) ([]*dto.GenreDto, error)
}

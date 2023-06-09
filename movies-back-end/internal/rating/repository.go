package rating

import (
	"context"
	"movies-service/internal/common/entity"
)

type Repository interface {
	FindRatings(ctx context.Context) ([]*entity.Rating, error)
}

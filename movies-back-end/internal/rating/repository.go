package rating

import (
	"context"
	"movies-service/internal/common/model"
)

type Repository interface {
	FindRatings(ctx context.Context) ([]*model.Rating, error)
}

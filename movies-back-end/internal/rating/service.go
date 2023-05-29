package rating

import (
	"context"
	"movies-service/internal/dto"
)

type Service interface {
	GetAllRatings(ctx context.Context) ([]*dto.RatingDto, error)
}

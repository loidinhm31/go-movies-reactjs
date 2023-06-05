package rating

import (
	"context"
	"movies-service/internal/common/dto"
)

type Service interface {
	GetAllRatings(ctx context.Context) ([]*dto.RatingDto, error)
}

package service

import (
	"context"
	"movies-service/internal/common/dto"
	"movies-service/internal/common/mapper"
	"movies-service/internal/rating"
)

type ratingService struct {
	ratingRepository rating.Repository
}

func NewRatingService(ratingRepository rating.Repository) rating.Service {
	return &ratingService{
		ratingRepository: ratingRepository,
	}
}

func (rs *ratingService) GetAllRatings(ctx context.Context) ([]*dto.RatingDto, error) {
	results, err := rs.ratingRepository.FindRatings(ctx)
	if err != nil {
		return nil, err
	}
	return mapper.MapToRatingDtoSlice(results), nil
}

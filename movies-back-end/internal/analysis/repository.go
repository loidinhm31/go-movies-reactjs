package analysis

import (
	"context"
	"movies-service/internal/common/dto"
	"movies-service/internal/common/entity"
)

type Repository interface {
	CountMoviesByGenre(ctx context.Context, movieType string) ([]*entity.GenreCount, error)
	CountMoviesByReleaseDate(ctx context.Context, year string, months []string) ([]*entity.MovieCount, error)
	CountMoviesByCreatedDate(ctx context.Context, year string, months []string) ([]*entity.MovieCount, error)
	CountViewsByGenreAndViewedDate(ctx context.Context, request *dto.RequestData) ([]*entity.ViewCount, error)
	CountCumulativeViewsByGenreAndViewedDate(ctx context.Context, request *dto.RequestData) ([]*entity.ViewCount, error)
	CountViewsByViewedDate(ctx context.Context, request *dto.RequestData) ([]*entity.ViewCount, error)
	CountMoviesByGenreAndReleasedDate(ctx context.Context, request *dto.RequestData) ([]*entity.MovieCount, error)
	SumTotalAmountAndTotalReceivedPayment(ctx context.Context, typeCode string) (*entity.TotalPayment, error)
}

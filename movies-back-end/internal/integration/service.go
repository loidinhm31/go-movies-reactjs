package integration

import (
	"context"
	"mime/multipart"
	"movies-service/internal/dto"
)

type Service interface {
	UploadVideo(ctx context.Context, file multipart.File) (string, error)
	DeleteVideo(ctx context.Context, fileId string) (string, error)
	GetMovies(ctx context.Context, movie *dto.MovieDto) ([]*dto.MovieDto, error)
	GetMovieById(ctx context.Context, movieId int64) (*dto.MovieDto, error)
}

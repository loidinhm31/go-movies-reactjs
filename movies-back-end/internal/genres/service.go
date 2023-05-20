package genres

import (
	"context"
	"github.com/gin-gonic/gin"
	"movies-service/internal/dto"
)

type Service interface {
	GetAllGenresByTypeCode(ctx context.Context, movieType string) ([]*dto.GenreDto, error)
	AddGenres(ctx *gin.Context, genreNames []dto.Genres) error
}

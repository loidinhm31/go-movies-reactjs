package genres

import (
	"context"
	"github.com/gin-gonic/gin"
	"movies-service/internal/dto"
)

type Service interface {
	GetAllGenres(ctx context.Context) ([]*dto.GenreDto, error)
	AddGenres(ctx *gin.Context, genreNames []string) error
}

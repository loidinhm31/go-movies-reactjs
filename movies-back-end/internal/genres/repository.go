package genres

import (
	"context"
	"github.com/gin-gonic/gin"
	"movies-service/internal/models"
)

type GenreRepository interface {
	FindAllGenres(ctx context.Context) ([]*models.Genre, error)
	FindGenreByName(ctx *gin.Context, genre *models.Genre) (*models.Genre, error)
	InsertGenres(ctx *gin.Context, genres []*models.Genre) error
}

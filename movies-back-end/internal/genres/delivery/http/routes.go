package http

import (
	"github.com/gin-gonic/gin"
	"movies-service/internal/genres"
)

// MapGenreRoutes Map auth routes
func MapGenreRoutes(genreGroup *gin.RouterGroup, h genres.GenreHandler) {
	genreGroup.GET("/", h.FetchGenres())
}

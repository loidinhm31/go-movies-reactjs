package http

import (
	"github.com/gin-gonic/gin"
	"movies-service/internal/genres"
)

// MapGenreRoutes Map genre routes
func MapGenreRoutes(genreGroup *gin.RouterGroup, h genres.GenreHandler) {
	genreGroup.GET("/", h.FetchGenres())
}

func MapAuthGenreRoutes(genreGroup *gin.RouterGroup, h genres.GenreHandler) {
	genreGroup.POST("/batch", h.PostGenres())
}

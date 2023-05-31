package http

import (
	"github.com/gin-gonic/gin"
	"movies-service/internal/genre"
)

// MapGenreRoutes Map genre routes
func MapGenreRoutes(genreGroup *gin.RouterGroup, h genre.Handler) {
	genreGroup.GET("/", h.FetchGenres())
}

func MapAuthGenreRoutes(genreGroup *gin.RouterGroup, h genre.Handler) {
	genreGroup.POST("/batch", h.PostGenres())
}

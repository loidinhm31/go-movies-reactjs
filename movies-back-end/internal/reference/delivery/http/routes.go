package http

import (
	"github.com/gin-gonic/gin"
	"movies-service/internal/reference"
)

// MapReferenceRoutes Map reference routes
func MapReferenceRoutes(refGroup *gin.RouterGroup, h reference.Handler) {
	refGroup.POST("/tmdb", h.FindMovies())
	refGroup.GET("/tmdb/:id", h.FindMovieById())
}

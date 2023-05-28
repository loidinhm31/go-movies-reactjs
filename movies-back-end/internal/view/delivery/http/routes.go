package http

import (
	"github.com/gin-gonic/gin"
	"movies-service/internal/view"
)

// MapViewRoutes Map view routes
func MapViewRoutes(viewGroup *gin.RouterGroup, h view.Handler) {
	viewGroup.POST("/", h.RecognizeViewForMovie())
	viewGroup.GET("/:movieId", h.FetchNumberOfViewsByMovieId())
}

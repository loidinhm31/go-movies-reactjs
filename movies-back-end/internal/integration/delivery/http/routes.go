package http

import (
	"github.com/gin-gonic/gin"
	"movies-service/internal/integration"
)

// MapIntegrationRoutes Map integration routes
func MapIntegrationRoutes(integrationGroup *gin.RouterGroup, h integration.Handler) {
	integrationGroup.POST("/videos", h.UploadVideo())
	integrationGroup.DELETE("/videos/:key", h.DeleteVideo())
	integrationGroup.POST("/tmdb", h.FindMovies())
	integrationGroup.GET("/tmdb/:id", h.FindMovieById())
}

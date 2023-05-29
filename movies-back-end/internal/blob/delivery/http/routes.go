package http

import (
	"github.com/gin-gonic/gin"
	"movies-service/internal/blob"
)

// MapIntegrationRoutes Map reference routes
func MapIntegrationRoutes(integrationGroup *gin.RouterGroup, h blob.Handler) {
	integrationGroup.POST("/videos", h.UploadVideo())
	integrationGroup.DELETE("/videos/:key", h.DeleteVideo())
}

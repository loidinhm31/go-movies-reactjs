package http

import (
	"github.com/gin-gonic/gin"
	"movies-service/internal/views"
)

// MapViewRoutes Map auth routes
func MapViewRoutes(viewGroup *gin.RouterGroup, h views.Handler) {
	viewGroup.POST("/", h.RecognizeViewForMovie())
}

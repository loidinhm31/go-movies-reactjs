package http

import (
	"github.com/gin-gonic/gin"
	"movies-service/internal/search"
)

// MapGraphRoutes Map auth routes
func MapGraphRoutes(graphGroup *gin.RouterGroup, h search.Handler) {
	graphGroup.POST("/", h.Search())
}

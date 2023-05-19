package http

import (
	"github.com/gin-gonic/gin"
	"movies-service/internal/roles"
)

// MapRoleRoutes Map role routes
func MapRoleRoutes(movieGroup *gin.RouterGroup, h roles.Handler) {
	movieGroup.GET("/", h.FetchRoles())
}

package http

import (
	"github.com/gin-gonic/gin"
	"movies-service/internal/role"
)

// MapRoleRoutes Map role routes
func MapRoleRoutes(movieGroup *gin.RouterGroup, h role.Handler) {
	movieGroup.GET("/", h.FetchRoles())
}

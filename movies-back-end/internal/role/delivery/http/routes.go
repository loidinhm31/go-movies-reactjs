package http

import (
	"github.com/gin-gonic/gin"
	"movies-service/internal/role"
)

// MapRoleRoutes Map role routes
func MapRoleRoutes(roleGroup *gin.RouterGroup, h role.Handler) {
	roleGroup.GET("/", h.FetchRoles())
}

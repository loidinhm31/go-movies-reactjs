package http

import (
	"github.com/gin-gonic/gin"
	"movies-service/internal/users"
)

// MapUserRoutes Map user routes
func MapUserRoutes(userGroup *gin.RouterGroup, h users.Handler) {
	userGroup.POST("/", h.FetchUsers())
	userGroup.PATCH("/role", h.PatchUserRole())
}

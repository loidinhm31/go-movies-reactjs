package http

import (
	"github.com/gin-gonic/gin"
	"movies-service/internal/user"
)

// MapUserRoutes Map user routes
func MapUserRoutes(userGroup *gin.RouterGroup, h user.Handler) {
	userGroup.POST("/", h.FetchUsers())
	userGroup.PATCH("/role", h.PatchUserRole())
	userGroup.PUT("/oidc", h.PutOidcUSer())
}

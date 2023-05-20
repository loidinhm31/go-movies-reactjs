package http

import (
	"github.com/gin-gonic/gin"
	"movies-service/internal/auth"
)

// MapAuthRoutes Map auth routes
func MapAuthRoutes(authGroup *gin.RouterGroup, h auth.Handler) {
	authGroup.GET("/ping", h.Ping())
	authGroup.POST("/login", h.Login())
	authGroup.POST("/signup", h.SignUp())
	authGroup.GET("/oidc", h.FetchUserFromOIDC())
}

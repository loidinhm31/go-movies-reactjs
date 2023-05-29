package http

import (
	"github.com/Nerzal/gocloak/v13"
	"github.com/gin-gonic/gin"
	"log"
	"movies-service/config"
	"movies-service/internal/auth"
	"movies-service/internal/dto"
	"movies-service/pkg/util"
	"net/http"
)

type authHandler struct {
	authService auth.Service
	keycloak    config.KeycloakConfig
	gocloak     *gocloak.GoCloak
}

func NewAuthHandler(authService auth.Service, keycloak config.KeycloakConfig, gocloak *gocloak.GoCloak) auth.Handler {
	return &authHandler{
		authService: authService,
		keycloak:    keycloak,
		gocloak:     gocloak,
	}
}

func (h *authHandler) Ping() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	}
}

func (h *authHandler) Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		user := &dto.UserDto{}

		if err := util.ReadRequest(c, user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "error",
			})
			c.Abort()
			return
		}

		userDto, err := h.authService.SignIn(c, user.Username)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": err.Error(),
			})
			c.Abort()
			return
		}
		c.JSON(http.StatusOK, userDto)
	}
}

func (h *authHandler) SignUp() gin.HandlerFunc {
	return func(c *gin.Context) {
		user := &dto.UserDto{}
		if err := util.ReadRequest(c, user); err != nil {
			log.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "error",
			})
			c.Abort()
			return
		}

		user, err := h.authService.SignUp(c, user)
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "cannot create user",
			})
			c.Abort()
			return
		}
		c.JSON(http.StatusOK, user)
	}
}

func (h *authHandler) FetchUserFromOIDC() gin.HandlerFunc {
	return func(c *gin.Context) {
		username := c.Query("username")

		user, err := h.authService.FindUserFromODIC(c, &username)
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err,
			})
			c.Abort()
			return
		}
		c.JSON(http.StatusOK, user)
	}
}

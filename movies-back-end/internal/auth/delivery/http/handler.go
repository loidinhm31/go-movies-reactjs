package http

import (
	"encoding/base64"
	"fmt"
	"github.com/Nerzal/gocloak/v13"
	"github.com/gin-gonic/gin"
	"movies-service/config"
	"movies-service/internal/auth"
	"movies-service/internal/dto"
	"movies-service/internal/errors"
	"movies-service/pkg/utils"
	"net/http"
	"strings"
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

		if err := utils.ReadRequest(c, user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "error",
			})
			c.Abort()
			return
		}

		userDto, err := h.authService.SignIn(c.Request.Context(), user.Username)
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

func (h *authHandler) VerifyToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		clientId, clientSecret, err := clientValidation(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			c.Abort()
			return
		}

		introspectForm := &dto.Introspect{
			ClientId:     clientId,
			ClientSecret: clientSecret,
		}

		if err := utils.ReadRequest(c, introspectForm); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "error",
			})
			c.Abort()
			return
		}

		//user, err := h.authService.VerifyToken(c.Request.Context(), c, introspectForm)
		//if err != nil {
		//	c.JSON(http.StatusBadRequest, gin.H{
		//		"message": err.Error(),
		//	})
		//	c.Abort()
		//	return
		//}
		//
		//c.JSON(http.StatusOK, gin.H{
		//	"client_id": introspectForm.ClientId,
		//	"username":  user.Username,
		//	"active":    true,
		//})
	}
}

func clientValidation(c *gin.Context) (string, string, error) {
	authHeader := c.Request.Header["Authorization"]

	if len(authHeader) == 0 {
		return "", "", errors.ErrInvalidClient
	}

	extractHeader := authHeader[0]
	headerParts := strings.Split(extractHeader, " ")
	if len(headerParts) != 2 {
		return "", "", errors.ErrInvalidClient
	}

	if headerParts[0] != "Basic" {
		return "", "", errors.ErrInvalidClient
	}

	rawDecodedAuthen, err := base64.StdEncoding.DecodeString(headerParts[1])
	if err != nil {
		return "", "", errors.ErrInvalidClient
	}

	authenParts := strings.Split(fmt.Sprintf("%s", rawDecodedAuthen), ":")
	if len(authenParts) != 2 {
		return "", "", errors.ErrInvalidClient
	}

	return authenParts[0], authenParts[1], nil
}

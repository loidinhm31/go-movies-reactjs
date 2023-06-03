package middlewares

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"movies-service/internal/model"
	"net/http"
	"strings"
)

const CtxUserKey = "username"
const CtxAccessToken = "access_token"

func (mw *MiddlewareManager) JWTValidation() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header["Authorization"]

		if len(authHeader) == 0 {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Unauthorized",
			})
			c.Abort()
			return
		}

		extractHeader := authHeader[0]
		headerParts := strings.Split(extractHeader, " ")
		if len(headerParts) != 2 {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Unauthorized",
			})
			c.Abort()
			return
		}

		if headerParts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Unauthorized",
			})
			c.Abort()
			return
		}

		// Call Keycloak API to verify the access token
		result, err := mw.gocloak.RetrospectToken(
			c.Request.Context(),
			headerParts[1],
			mw.keycloak.ClientId,
			mw.keycloak.ClientSecret,
			mw.keycloak.Realm,
		)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": err.Error(),
			})
			c.Abort()
			return
		}

		jwt, _, _ := mw.gocloak.DecodeAccessToken(
			c.Request.Context(),
			headerParts[1],
			mw.keycloak.Realm,
		)

		jwtj, _ := json.Marshal(jwt)

		var userToken model.UserToken
		err = json.Unmarshal(jwtj, &userToken)
		if err != nil {
			fmt.Printf(err.Error())
			return
		}

		// Check if the token isn't expired and valid
		if !*result.Active {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "invalid token",
			})
			c.Abort()
			return
		}
		c.Set(CtxUserKey, userToken.Claims.Username)
		c.Set(CtxAccessToken, headerParts[1])

		c.Next()
	}
}

func (mw *MiddlewareManager) JWTValidationOptional() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header["Authorization"]

		if len(authHeader) == 0 {
			return
		}

		extractHeader := authHeader[0]
		headerParts := strings.Split(extractHeader, " ")
		if len(headerParts) != 2 {
			return
		}

		if headerParts[0] != "Bearer" {
			return
		}

		// Call Keycloak API to verify the access token
		result, err := mw.gocloak.RetrospectToken(
			c.Request.Context(),
			headerParts[1],
			mw.keycloak.ClientId,
			mw.keycloak.ClientSecret,
			mw.keycloak.Realm,
		)
		if err != nil {
			return
		}

		jwt, _, _ := mw.gocloak.DecodeAccessToken(
			c.Request.Context(),
			headerParts[1],
			mw.keycloak.Realm,
		)

		jwtj, _ := json.Marshal(jwt)

		var userToken model.UserToken
		err = json.Unmarshal(jwtj, &userToken)
		if err != nil {
			fmt.Printf(err.Error())
			return
		}

		// Check if the token isn't expired and valid
		if !*result.Active {
			return
		}
		c.Set(CtxUserKey, userToken.Claims.Username)
		c.Set(CtxAccessToken, headerParts[1])

		c.Next()
	}
}

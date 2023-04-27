package middlewares

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

const CtxUserKey = "userId"

func (mw *MiddlewareManager) JWTValidation() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header["Authorization"]

		if len(authHeader) == 0 {
			c.JSON(http.StatusInternalServerError, gin.H{
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

		jwt, _, err := mw.gocloak.DecodeAccessToken(
			c.Request.Context(),
			headerParts[1],
			mw.keycloak.Realm,
		)

		jwtj, _ := json.Marshal(jwt)
		fmt.Printf("token: %v\n", string(jwtj))

		// Check if the token isn't expired and valid
		if !*result.Active {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "invalid token",
			})
			c.Abort()
			return
		}

		c.Next()
	}
}

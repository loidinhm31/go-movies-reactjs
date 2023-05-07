package auth

import "github.com/gin-gonic/gin"

type Handler interface {
	Ping() gin.HandlerFunc
	Login() gin.HandlerFunc
}

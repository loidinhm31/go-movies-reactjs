package auth

import "github.com/gin-gonic/gin"

type Handler interface {
	Ping() gin.HandlerFunc
	FetchGenres() gin.HandlerFunc
	Login() gin.HandlerFunc
	VerifyToken() gin.HandlerFunc
}

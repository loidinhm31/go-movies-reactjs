package users

import "github.com/gin-gonic/gin"

type Handler interface {
	FetchUsers() gin.HandlerFunc
	PatchUserRole() gin.HandlerFunc
	PutOidcUSer() gin.HandlerFunc
}

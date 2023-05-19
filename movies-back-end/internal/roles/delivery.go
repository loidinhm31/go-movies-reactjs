package roles

import "github.com/gin-gonic/gin"

type Handler interface {
	FetchRoles() gin.HandlerFunc
}

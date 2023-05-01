package search

import "github.com/gin-gonic/gin"

type Handler interface {
	Search() gin.HandlerFunc
}

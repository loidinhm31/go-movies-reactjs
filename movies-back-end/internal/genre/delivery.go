package genre

import "github.com/gin-gonic/gin"

type Handler interface {
	FetchGenres() gin.HandlerFunc
	PostGenres() gin.HandlerFunc
}

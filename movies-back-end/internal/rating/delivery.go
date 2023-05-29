package rating

import "github.com/gin-gonic/gin"

type Handler interface {
	FetchRatings() gin.HandlerFunc
}

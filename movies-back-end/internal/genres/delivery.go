package genres

import "github.com/gin-gonic/gin"

type GenreHandler interface {
	FetchGenres() gin.HandlerFunc
	PostGenres() gin.HandlerFunc
}

package reference

import "github.com/gin-gonic/gin"

type Handler interface {
	FindMovies() gin.HandlerFunc
	FindMovieById() gin.HandlerFunc
}

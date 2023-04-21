package http

import (
	"github.com/gin-gonic/gin"
	"movies-service/internal/movies"
)

// MapMovieRoutes Map auth routes
func MapMovieRoutes(movieGroup *gin.RouterGroup, h movies.MovieHandler) {
	movieGroup.GET("/", h.FetchMovies())
	movieGroup.GET("/:id", h.FetchMovieById())
	movieGroup.GET("/genres/:id", h.FetchMovieByGenre())
	movieGroup.PUT("/", h.PutMovie())
	movieGroup.DELETE("/:id", h.DeleteMovie())
}

package http

import (
	"movies-service/internal/movie"

	"github.com/gin-gonic/gin"
)

// MapMovieRoutes Map movie routes
func MapMovieRoutes(movieGroup *gin.RouterGroup, h movie.Handler) {
	movieGroup.POST("/page", h.FetchMovies())
	movieGroup.GET("/page", h.FetchMovies())
	movieGroup.GET("/:id", h.FetchMovieById())
	movieGroup.POST("/genres/:id", h.FetchMovieByGenre())
}

func MapAuthMovieRoutes(movieGroup *gin.RouterGroup, h movie.Handler) {
	movieGroup.POST("/", h.PostMovie())
	movieGroup.DELETE("/:id", h.DeleteMovie())
	movieGroup.PUT("/", h.PutMovie())
	movieGroup.GET("/episodes/:id", h.FetchMovieByEpisode())
	movieGroup.PUT("/:id/price", h.PatchMoviePrice())
}

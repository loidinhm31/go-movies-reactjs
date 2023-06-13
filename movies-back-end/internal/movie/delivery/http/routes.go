package http

import (
	"movies-service/internal/movie"

	"github.com/gin-gonic/gin"
)

// MapMovieRoutes Map auth routes
func MapMovieRoutes(movieGroup *gin.RouterGroup, h movie.Handler) {
	movieGroup.POST("/", h.FetchMoviesByType())
	movieGroup.GET("/", h.FetchMoviesByType())
	movieGroup.GET("/:id", h.FetchMovieById())
	movieGroup.POST("/genres/:id", h.FetchMovieByGenre())
}

func MapAuthMovieRoutes(movieGroup *gin.RouterGroup, h movie.Handler) {
	movieGroup.PUT("/", h.PutMovie())
	movieGroup.DELETE("/:id", h.DeleteMovie())
	movieGroup.PATCH("/", h.PatchMovie())
	movieGroup.GET("/", h.FetchMovies())
	movieGroup.PATCH("/:id/price", h.PatchMoviePrice())
}

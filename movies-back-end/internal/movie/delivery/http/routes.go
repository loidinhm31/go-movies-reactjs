package http

import (
	"github.com/gin-gonic/gin"
	"movies-service/internal/movie"
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
}

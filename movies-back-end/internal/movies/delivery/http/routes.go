package http

import (
	"github.com/gin-gonic/gin"
	"movies-service/internal/movies"
)

// MapMovieRoutes Map auth routes
func MapMovieRoutes(movieGroup *gin.RouterGroup, h movies.MovieHandler) {
	movieGroup.POST("/", h.FetchMoviesByType())
	movieGroup.GET("/:id", h.FetchMovieById())
	movieGroup.POST("/genres/:id", h.FetchMovieByGenre())
}

func MapAuthMovieRoutes(movieGroup *gin.RouterGroup, h movies.MovieHandler) {
	movieGroup.PUT("/", h.PutMovie())
	movieGroup.DELETE("/:id", h.DeleteMovie())
	movieGroup.PATCH("/", h.PatchMovie())
}

package http

import (
	"github.com/gin-gonic/gin"
	"movies-service/internal/movies"
)

// MapMovieRoutes Map auth routes
func MapMovieRoutes(movieGroup *gin.RouterGroup, h movies.MovieHandler) {
	movieGroup.POST("/", h.FetchMovies())
	movieGroup.GET("/:id", h.FetchMovieById())
	movieGroup.POST("/genres/:id", h.FetchMovieByGenre())
}

func MapRoleMovieRoutes(movieGroup *gin.RouterGroup, h movies.MovieHandler) {
	movieGroup.PUT("/", h.PutMovie())
	movieGroup.DELETE("/:id", h.DeleteMovie())
	movieGroup.PATCH("/", h.PatchMovie())
}

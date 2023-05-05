package http

import (
	"github.com/gin-gonic/gin"
	"movies-service/internal/analysis"
)

// MapAnalysisRoutes Map auth routes
func MapAnalysisRoutes(analysisGroup *gin.RouterGroup, h analysis.Handler) {
	analysisGroup.GET("/genres/movies", h.FetchNumberOfMoviesByGenre())
	analysisGroup.POST("/movies/release-date", h.FetchNumberOfMoviesByReleaseDate())
	analysisGroup.POST("/movies/created-date", h.FetchNumberOfMoviesByCreatedDate())
	analysisGroup.POST("/genres/views", h.FetchNumberOfViewsByGenreAndViewedDate())
	analysisGroup.POST("/views", h.FetchNumberOfViewsByViewedDate())
}

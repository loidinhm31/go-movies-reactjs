package http

import (
	"github.com/gin-gonic/gin"
	"movies-service/internal/analysis"
)

// MapAnalysisRoutes Map auth routes
func MapAnalysisRoutes(analysisGroup *gin.RouterGroup, h analysis.Handler) {
	analysisGroup.GET("/movies/genres", h.FetchNumberOfMoviesByGenre())
	analysisGroup.POST("/movies/genres/release-date", h.FetchNumberOfMoviesByGenreAndReleaseDate())
	analysisGroup.POST("/movies/release-date", h.FetchNumberOfMoviesByReleaseDate())
	analysisGroup.POST("/movies/created-date", h.FetchNumberOfMoviesByCreatedDate())
	analysisGroup.POST("/views/genres", h.FetchViewsByGenreAndViewedDate())
	analysisGroup.POST("/views", h.FetchNumberOfViewsByViewedDate())
	analysisGroup.GET("/payments", h.FetchTotalAmountAndTotalReceivedPayment())
}

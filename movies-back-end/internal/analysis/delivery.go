package analysis

import (
	"github.com/gin-gonic/gin"
)

type Handler interface {
	FetchNumberOfMoviesByGenre() gin.HandlerFunc
	FetchNumberOfMoviesByReleaseDate() gin.HandlerFunc
	FetchNumberOfMoviesByCreatedDate() gin.HandlerFunc
	FetchViewsByGenreAndViewedDate() gin.HandlerFunc
	FetchNumberOfViewsByViewedDate() gin.HandlerFunc
	FetchNumberOfMoviesByGenreAndReleaseDate() gin.HandlerFunc
	FetchTotalAmountAndTotalReceivedPayment() gin.HandlerFunc
}

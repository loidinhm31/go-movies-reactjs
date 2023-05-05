package analysis

import "github.com/gin-gonic/gin"

type Handler interface {
	FetchNumberOfMoviesByGenre() gin.HandlerFunc
	FetchNumberOfMoviesByReleaseDate() gin.HandlerFunc
	FetchNumberOfMoviesByCreatedDate() gin.HandlerFunc
	FetchNumberOfViewsByGenreAndViewedDate() gin.HandlerFunc
	FetchNumberOfViewsByViewedDate() gin.HandlerFunc
}

package movies

import "github.com/gin-gonic/gin"

type MovieHandler interface {
	FetchMoviesByType() gin.HandlerFunc
	FetchMovieById() gin.HandlerFunc
	FetchMovieByGenre() gin.HandlerFunc
	PutMovie() gin.HandlerFunc
	DeleteMovie() gin.HandlerFunc
	PatchMovie() gin.HandlerFunc
}

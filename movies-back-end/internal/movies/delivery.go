package movies

import "github.com/gin-gonic/gin"

type MovieHandler interface {
	FetchMovies() gin.HandlerFunc
	FetchMovieById() gin.HandlerFunc
	FetchMovieByGenre() gin.HandlerFunc
	PutMovie() gin.HandlerFunc
	DeleteMovie() gin.HandlerFunc
	PatchMovie() gin.HandlerFunc
}

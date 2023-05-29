package movie

import "github.com/gin-gonic/gin"

type Handler interface {
	FetchMoviesByType() gin.HandlerFunc
	FetchMovieById() gin.HandlerFunc
	FetchMovieByGenre() gin.HandlerFunc
	PutMovie() gin.HandlerFunc
	DeleteMovie() gin.HandlerFunc
	PatchMovie() gin.HandlerFunc
}

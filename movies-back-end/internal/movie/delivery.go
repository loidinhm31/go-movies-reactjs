package movie

import "github.com/gin-gonic/gin"

type Handler interface {
	FetchMovies() gin.HandlerFunc
	FetchMovieById() gin.HandlerFunc
	FetchMovieByGenre() gin.HandlerFunc
	PutMovie() gin.HandlerFunc
	DeleteMovie() gin.HandlerFunc
	PatchMovie() gin.HandlerFunc
	FetchMovieByEpisode() gin.HandlerFunc
	PatchMoviePrice() gin.HandlerFunc
}

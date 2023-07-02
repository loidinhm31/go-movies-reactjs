package movie

import "github.com/gin-gonic/gin"

type Handler interface {
	FetchMovies() gin.HandlerFunc
	FetchMovieById() gin.HandlerFunc
	FetchMovieByGenre() gin.HandlerFunc
	PostMovie() gin.HandlerFunc
	DeleteMovie() gin.HandlerFunc
	PutMovie() gin.HandlerFunc
	FetchMovieByEpisode() gin.HandlerFunc
	PatchMoviePrice() gin.HandlerFunc
}

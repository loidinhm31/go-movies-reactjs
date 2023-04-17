package http

import (
	"github.com/gin-gonic/gin"
	"log"
	"movies-service/internal/movies"
	"net/http"
	"strconv"
)

type movieHandler struct {
	movieService movies.Service
}

func NewMovieHandler(movieService movies.Service) movies.MovieHandler {
	return &movieHandler{
		movieService: movieService,
	}
}

func (mh *movieHandler) FetchMovies() gin.HandlerFunc {
	return func(c *gin.Context) {
		movies, err := mh.movieService.GetAllMovies(c.Request.Context())
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			c.Abort()
			return
		}
		c.JSON(http.StatusOK, movies)
	}
}

func (mh *movieHandler) FetchMovieById() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		movieId, err := strconv.Atoi(id)
		if err != nil {
			log.Println("Error during conversion")
			return
		}

		movie, err := mh.movieService.GetMovieById(c.Request.Context(), movieId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			c.Abort()
			return
		}
		c.JSON(http.StatusOK, movie)
	}
}

func (mh *movieHandler) FetchMovieByGenre() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		genreId, err := strconv.Atoi(id)
		if err != nil {
			log.Println("Error during conversion")
			return
		}

		movies, err := mh.movieService.GetMoviesByGenre(c.Request.Context(), genreId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			c.Abort()
			return
		}
		c.JSON(http.StatusOK, movies)
	}
}

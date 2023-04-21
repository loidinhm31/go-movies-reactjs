package http

import (
	"github.com/gin-gonic/gin"
	"log"
	"movies-service/internal/dto"
	"movies-service/internal/movies"
	"movies-service/pkg/utils"
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
		allMovies, err := mh.movieService.GetAllMovies(c.Request.Context())
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			c.Abort()
			return
		}
		c.JSON(http.StatusOK, allMovies)
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

		movieDtos, err := mh.movieService.GetMoviesByGenre(c.Request.Context(), genreId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			c.Abort()
			return
		}
		c.JSON(http.StatusOK, movieDtos)
	}
}

func (mh *movieHandler) PutMovie() gin.HandlerFunc {
	return func(c *gin.Context) {
		movie := &dto.MovieDto{}

		if err := utils.ReadRequest(c, movie); err != nil {
			log.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "error",
			})
			c.Abort()
			return
		}

		err := mh.movieService.AddMovie(c, movie)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			c.Abort()
			return
		}
	}
}

func (mh *movieHandler) DeleteMovie() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		movieId, err := strconv.Atoi(id)
		if err != nil {
			log.Println("Error during conversion")
			return
		}

		err = mh.movieService.DeleteMovieById(c.Request.Context(), movieId)
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			c.Abort()
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "deleted"})
	}
}

func (mh *movieHandler) PatchMovie() gin.HandlerFunc {
	return func(c *gin.Context) {
		movie := &dto.MovieDto{}

		if err := utils.ReadRequest(c, movie); err != nil {
			log.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "error",
			})
			c.Abort()
			return
		}

		err := mh.movieService.UpdateMovie(c, movie)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			c.Abort()
			return
		}
	}
}

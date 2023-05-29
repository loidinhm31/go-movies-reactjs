package http

import (
	"github.com/gin-gonic/gin"
	"log"
	"movies-service/internal/dto"
	"movies-service/internal/movie"
	"movies-service/pkg/pagination"
	"movies-service/pkg/util"
	"net/http"
	"strconv"
)

type movieHandler struct {
	movieService movie.Service
}

func NewMovieHandler(movieService movie.Service) movie.Handler {
	return &movieHandler{
		movieService: movieService,
	}
}

func (mh *movieHandler) FetchMoviesByType() gin.HandlerFunc {
	return func(c *gin.Context) {
		keyword := c.Query("q")
		movieType := c.Query("type")

		pageable, _ := pagination.ReadPageRequest(c)

		if err := util.ReadRequest(c, pageable); err != nil {
			log.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "error",
			})
			c.Abort()
			return
		}

		allMovies, err := mh.movieService.GetAllMoviesByType(c, keyword, movieType, pageable)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
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

		result, err := mh.movieService.GetMovieById(c, movieId)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			c.Abort()
			return
		}
		c.JSON(http.StatusOK, result)
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

		pageable, _ := pagination.ReadPageRequest(c)

		if err := util.ReadRequest(c, pageable); err != nil {
			log.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "error",
			})
			c.Abort()
			return
		}

		movieDtos, err := mh.movieService.GetMoviesByGenre(c, pageable, genreId)
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
		theMovie := &dto.MovieDto{}

		if err := util.ReadRequest(c, theMovie); err != nil {
			log.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "error",
			})
			c.Abort()
			return
		}

		err := mh.movieService.AddMovie(c, theMovie)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			c.Abort()
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "ok"})
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

		err = mh.movieService.DeleteMovieById(c, movieId)
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			c.Abort()
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "ok"})
	}
}

func (mh *movieHandler) PatchMovie() gin.HandlerFunc {
	return func(c *gin.Context) {
		theMovie := &dto.MovieDto{}

		if err := util.ReadRequest(c, theMovie); err != nil {
			log.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "error",
			})
			c.Abort()
			return
		}

		err := mh.movieService.UpdateMovie(c, theMovie)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			c.Abort()
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "ok"})
	}
}

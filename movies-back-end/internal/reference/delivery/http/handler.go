package http

import (
	"log"
	"movies-service/internal/common/dto"
	"movies-service/internal/reference"
	"movies-service/pkg/util"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type referenceHandler struct {
	referenceService reference.Service
}

func NewReferenceHandler(referebceService reference.Service) reference.Handler {
	return &referenceHandler{
		referenceService: referebceService,
	}
}

func (ih *referenceHandler) FindMovies() gin.HandlerFunc {
	return func(c *gin.Context) {
		movie := &dto.MovieDto{}
		if err := util.ReadRequest(c, movie); err != nil {
			log.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "error",
			})
			c.Abort()
			return
		}

		results, err := ih.referenceService.GetMoviesByType(c, movie)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "cannot access resource",
			})
			c.Abort()
			return
		}

		c.JSON(http.StatusOK, results)
	}
}

func (ih *referenceHandler) FindMovieById() gin.HandlerFunc {
	return func(c *gin.Context) {
		ID := c.Param("id")
		movieType := c.Query("type")

		movieId, _ := strconv.ParseInt(ID, 10, 64)

		result, err := ih.referenceService.GetMovieById(c, movieId, movieType)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "cannot access resource",
			})
			c.Abort()
			return
		}

		c.JSON(http.StatusOK, result)
	}
}

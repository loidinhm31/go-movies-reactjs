package http

import (
	"log"
	"movies-service/internal/dto"
	"movies-service/internal/integration"
	"movies-service/pkg/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type integrationHandler struct {
	integrationService integration.Service
}

func NewIntegrationHandler(integrationService integration.Service) integration.Handler {
	return &integrationHandler{
		integrationService: integrationService,
	}
}

func (ih *integrationHandler) UploadVideo() gin.HandlerFunc {
	return func(c *gin.Context) {
		file, _, err := c.Request.FormFile("file")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":   err,
				"message": "failed to upload",
			})
			c.Abort()
			return
		}

		resp, err := ih.integrationService.UploadVideo(c, file)
		if err != nil || resp == "" {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			c.Abort()
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"fileName": resp,
		})
	}
}

func (ih *integrationHandler) DeleteVideo() gin.HandlerFunc {
	return func(c *gin.Context) {
		videoKey := c.Param("key")

		res, err := ih.integrationService.DeleteVideo(c, videoKey)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			c.Abort()
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"result": res,
		})
	}
}

func (ih *integrationHandler) FindMovies() gin.HandlerFunc {
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

		results, err := ih.integrationService.GetMovies(c, movie)
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

func (ih *integrationHandler) FindMovieById() gin.HandlerFunc {
	return func(c *gin.Context) {
		ID := c.Param("id")
		movieId, _ := strconv.ParseInt(ID, 10, 64)

		result, err := ih.integrationService.GetMovieById(c, movieId)
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

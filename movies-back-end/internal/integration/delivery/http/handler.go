package http

import (
	"github.com/gin-gonic/gin"
	"movies-service/internal/integration"
	"net/http"
)

type integrationHandler struct {
	integrationService integration.Service
}

func NewIntegrationHandler(integrationService integration.Service) integration.Handler {
	return &integrationHandler{
		integrationService: integrationService,
	}
}

func (mh *integrationHandler) UploadVideo() gin.HandlerFunc {
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

		resp, err := mh.integrationService.UploadVideo(c, file)
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

func (mh *integrationHandler) DeleteVideo() gin.HandlerFunc {
	return func(c *gin.Context) {
		videoKey := c.Param("key")

		res, err := mh.integrationService.DeleteVideo(c, videoKey)
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

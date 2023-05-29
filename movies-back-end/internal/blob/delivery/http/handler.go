package http

import (
	"github.com/gin-gonic/gin"
	"log"
	"movies-service/internal/blob"
	"net/http"
)

type blobHandler struct {
	blobService blob.Service
}

func NewBlobHandler(blobService blob.Service) blob.Handler {
	return &blobHandler{
		blobService: blobService,
	}
}

func (ih *blobHandler) UploadVideo() gin.HandlerFunc {
	return func(c *gin.Context) {
		file, _, err := c.Request.FormFile("file")
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{
				"error":   err,
				"message": "failed to upload",
			})
			c.Abort()
			return
		}

		resp, err := ih.blobService.UploadVideo(c, file)
		if err != nil || resp == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			c.Abort()
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"message":  "ok",
			"fileName": resp,
		})
	}
}

func (ih *blobHandler) DeleteVideo() gin.HandlerFunc {
	return func(c *gin.Context) {
		videoKey := c.Param("key")

		res, err := ih.blobService.DeleteVideo(c, videoKey)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			c.Abort()
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "ok",
			"result":  res,
		})
	}
}

package integration

import "github.com/gin-gonic/gin"

type Handler interface {
	UploadVideo() gin.HandlerFunc
	DeleteVideo() gin.HandlerFunc
}

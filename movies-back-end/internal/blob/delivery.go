package blob

import "github.com/gin-gonic/gin"

type Handler interface {
	UploadFile() gin.HandlerFunc
	DeleteFile() gin.HandlerFunc
}

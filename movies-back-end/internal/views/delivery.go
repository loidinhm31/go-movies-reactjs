package views

import "github.com/gin-gonic/gin"

type Handler interface {
	RecognizeViewForMovie() gin.HandlerFunc
}

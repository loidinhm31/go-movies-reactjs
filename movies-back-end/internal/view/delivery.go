package view

import "github.com/gin-gonic/gin"

type Handler interface {
	RecognizeViewForMovie() gin.HandlerFunc
	FetchNumberOfViewsByMovieId() gin.HandlerFunc
}

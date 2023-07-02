package season

import "github.com/gin-gonic/gin"

type Handler interface {
	FetchSeasonsByID() gin.HandlerFunc
	FetchSeasonsByMovieID() gin.HandlerFunc
	PostSeason() gin.HandlerFunc
	PutSeason() gin.HandlerFunc
	DeleteSeason() gin.HandlerFunc
}

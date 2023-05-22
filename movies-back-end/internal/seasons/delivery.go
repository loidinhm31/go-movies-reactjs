package seasons

import "github.com/gin-gonic/gin"

type Handler interface {
	FetchSeasonsByID() gin.HandlerFunc
	FetchSeasonsByMovieID() gin.HandlerFunc
	PutSeason() gin.HandlerFunc
	PatchSeason() gin.HandlerFunc
	DeleteSeason() gin.HandlerFunc
}

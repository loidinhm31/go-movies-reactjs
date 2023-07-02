package http

import (
	"github.com/gin-gonic/gin"
	"movies-service/internal/season"
)

func MapSeasonRoutes(seasonGroup *gin.RouterGroup, h season.Handler) {
	seasonGroup.GET("/:id", h.FetchSeasonsByID())
	seasonGroup.GET("/", h.FetchSeasonsByMovieID())
}

func MapAuthSeasonRoutes(seasonGroup *gin.RouterGroup, h season.Handler) {
	seasonGroup.GET("/", h.FetchSeasonsByMovieID())
	seasonGroup.POST("/", h.PostSeason())
	seasonGroup.PUT("/", h.PutSeason())
	seasonGroup.DELETE("/:id", h.DeleteSeason())
}

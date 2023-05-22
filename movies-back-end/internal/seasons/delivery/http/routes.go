package http

import (
	"github.com/gin-gonic/gin"
	"movies-service/internal/seasons"
)

func MapSeasonRoutes(seasonGroup *gin.RouterGroup, h seasons.Handler) {
	seasonGroup.GET("/:id", h.FetchSeasonsByID())
	seasonGroup.GET("/", h.FetchSeasonsByMovieID())
}

func MapAuthSeasonRoutes(seasonGroup *gin.RouterGroup, h seasons.Handler) {
	seasonGroup.GET("/", h.FetchSeasonsByMovieID())
	seasonGroup.PUT("/", h.PutSeason())
	seasonGroup.PATCH("/", h.PatchSeason())
	seasonGroup.DELETE("/:id", h.DeleteSeason())
}

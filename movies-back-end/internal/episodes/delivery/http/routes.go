package http

import (
	"github.com/gin-gonic/gin"
	"movies-service/internal/episodes"
)

func MapEpisodeRoutes(episodeGroup *gin.RouterGroup, h episodes.Handler) {
	episodeGroup.GET("/:id", h.FetchEpisodesByID())
	episodeGroup.GET("/", h.FetchEpisodesBySeasonID())
}

func MapAuthEpisodeRoutes(episodeGroup *gin.RouterGroup, h episodes.Handler) {
	episodeGroup.PUT("/", h.PutEpisode())
	episodeGroup.PATCH("/", h.PatchEpisode())
	episodeGroup.DELETE("/:id", h.DeleteEpisode())
}

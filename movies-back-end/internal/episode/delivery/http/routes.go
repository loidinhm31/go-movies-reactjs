package http

import (
	"github.com/gin-gonic/gin"
	"movies-service/internal/episode"
)

func MapEpisodeRoutes(episodeGroup *gin.RouterGroup, h episode.Handler) {
	episodeGroup.GET("/:id", h.FetchEpisodeByID())
	episodeGroup.GET("/", h.FetchEpisodesBySeasonID())
}

func MapAuthEpisodeRoutes(episodeGroup *gin.RouterGroup, h episode.Handler) {
	episodeGroup.PUT("/", h.PutEpisode())
	episodeGroup.PATCH("/", h.PatchEpisode())
	episodeGroup.DELETE("/:id", h.DeleteEpisode())
}

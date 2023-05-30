package episode

import "github.com/gin-gonic/gin"

type Handler interface {
	FetchEpisodesByID() gin.HandlerFunc
	FetchEpisodesBySeasonID() gin.HandlerFunc
	PutEpisode() gin.HandlerFunc
	PatchEpisode() gin.HandlerFunc
	DeleteEpisode() gin.HandlerFunc
}

package episode

import "github.com/gin-gonic/gin"

type Handler interface {
	FetchEpisodeByID() gin.HandlerFunc
	FetchEpisodesBySeasonID() gin.HandlerFunc
	PostEpisode() gin.HandlerFunc
	PutEpisode() gin.HandlerFunc
	DeleteEpisode() gin.HandlerFunc
}

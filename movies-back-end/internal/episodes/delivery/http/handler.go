package http

import (
	"github.com/gin-gonic/gin"
	"log"
	"movies-service/internal/dto"
	"movies-service/internal/episodes"
	"movies-service/pkg/utils"
	"net/http"
	"strconv"
)

type episodeHandler struct {
	episodeService episodes.Service
}

func NewEpisodeHandler(episodeService episodes.Service) episodes.Handler {
	return &episodeHandler{
		episodeService: episodeService,
	}
}

func (e episodeHandler) FetchEpisodesByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		episodeID, _ := strconv.Atoi(id)

		allEpisodes, err := e.episodeService.GetEpisodesByID(c, episodeID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			c.Abort()
			return
		}
		c.JSON(http.StatusOK, allEpisodes)
	}
}

func (e episodeHandler) FetchEpisodesBySeasonID() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Query("seasonID")
		seasonID, _ := strconv.Atoi(id)

		allEpisodes, err := e.episodeService.GetEpisodesBySeasonID(c, seasonID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			c.Abort()
			return
		}
		c.JSON(http.StatusOK, allEpisodes)
	}
}

func (e episodeHandler) PutEpisode() gin.HandlerFunc {
	return func(c *gin.Context) {
		episode := &dto.EpisodeDto{}
		if err := utils.ReadRequest(c, episode); err != nil {
			log.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "error",
			})
			c.Abort()
			return
		}

		err := e.episodeService.AddEpisode(c, episode)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			c.Abort()
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "ok"})
	}
}

func (e episodeHandler) PatchEpisode() gin.HandlerFunc {
	return func(c *gin.Context) {
		episode := &dto.EpisodeDto{}
		if err := utils.ReadRequest(c, episode); err != nil {
			log.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "error",
			})
			c.Abort()
			return
		}

		err := e.episodeService.UpdateEpisode(c, episode)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			c.Abort()
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "ok"})
	}
}

func (e episodeHandler) DeleteEpisode() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		seasonID, err := strconv.Atoi(id)
		if err != nil {
			log.Println("Error during conversion")
			return
		}

		err = e.episodeService.DeleteEpisodeById(c, seasonID)
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			c.Abort()
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "ok"})
	}
}

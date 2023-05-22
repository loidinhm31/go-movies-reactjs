package http

import (
	"github.com/gin-gonic/gin"
	"log"
	"movies-service/internal/dto"
	"movies-service/internal/seasons"
	"movies-service/pkg/utils"
	"net/http"
	"strconv"
)

type seasonHandler struct {
	seasonService seasons.Service
}

func NewSeasonHandler(seasonService seasons.Service) seasons.Handler {
	return &seasonHandler{
		seasonService: seasonService,
	}
}

func (s seasonHandler) FetchSeasonsByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		seasonID, _ := strconv.Atoi(id)

		allSeasons, err := s.seasonService.GetSeasonsByID(c, seasonID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			c.Abort()
			return
		}
		c.JSON(http.StatusOK, allSeasons)
	}
}

func (s seasonHandler) FetchSeasonsByMovieID() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Query("movieID")
		movieID, _ := strconv.Atoi(id)

		allSeasons, err := s.seasonService.GetSeasonsByMovieID(c, movieID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			c.Abort()
			return
		}
		c.JSON(http.StatusOK, allSeasons)
	}
}

func (s seasonHandler) PutSeason() gin.HandlerFunc {
	return func(c *gin.Context) {
		season := &dto.SeasonDto{}
		if err := utils.ReadRequest(c, season); err != nil {
			log.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "error",
			})
			c.Abort()
			return
		}

		err := s.seasonService.AddSeason(c, season)
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

func (s seasonHandler) PatchSeason() gin.HandlerFunc {
	return func(c *gin.Context) {
		season := &dto.SeasonDto{}
		if err := utils.ReadRequest(c, season); err != nil {
			log.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "error",
			})
			c.Abort()
			return
		}

		err := s.seasonService.UpdateSeason(c, season)
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

func (s seasonHandler) DeleteSeason() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		seasonID, err := strconv.Atoi(id)
		if err != nil {
			log.Println("Error during conversion")
			return
		}

		err = s.seasonService.DeleteSeasonById(c, seasonID)
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

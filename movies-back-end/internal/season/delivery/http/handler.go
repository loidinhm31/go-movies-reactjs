package http

import (
	"github.com/gin-gonic/gin"
	"log"
	"movies-service/internal/common/dto"
	"movies-service/internal/season"
	"movies-service/pkg/util"
	"net/http"
	"strconv"
)

type seasonHandler struct {
	seasonService season.Service
}

func NewSeasonHandler(seasonService season.Service) season.Handler {
	return &seasonHandler{
		seasonService: seasonService,
	}
}

func (s seasonHandler) FetchSeasonsByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		seasonID, _ := strconv.Atoi(id)

		allSeasons, err := s.seasonService.GetSeasonsByID(c, uint(seasonID))
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

		allSeasons, err := s.seasonService.GetSeasonsByMovieID(c, uint(movieID))
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

func (s seasonHandler) PostSeason() gin.HandlerFunc {
	return func(c *gin.Context) {
		season := &dto.SeasonDto{}
		if err := util.ReadRequest(c, season); err != nil {
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

func (s seasonHandler) PutSeason() gin.HandlerFunc {
	return func(c *gin.Context) {
		theSeason := &dto.SeasonDto{}
		if err := util.ReadRequest(c, theSeason); err != nil {
			log.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "error",
			})
			c.Abort()
			return
		}

		err := s.seasonService.UpdateSeason(c, theSeason)
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

		err = s.seasonService.RemoveSeasonByID(c, uint(seasonID))
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

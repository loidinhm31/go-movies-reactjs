package http

import (
	"github.com/gin-gonic/gin"
	"log"
	"movies-service/internal/analysis"
	"movies-service/internal/dto"
	"movies-service/pkg/util"
	"net/http"
)

type researchHandler struct {
	researchService analysis.Service
}

func NewAnalysisHandler(researchService analysis.Service) analysis.Handler {
	return &researchHandler{
		researchService: researchService,
	}
}

func (rh *researchHandler) FetchNumberOfMoviesByGenre() gin.HandlerFunc {
	return func(c *gin.Context) {
		movieType := c.Query("type")

		result, err := rh.researchService.GetNumberOfMoviesByGenre(c, movieType)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			c.Abort()
			return
		}
		c.JSON(http.StatusOK, result)
	}
}

func (rh *researchHandler) FetchNumberOfMoviesByReleaseDate() gin.HandlerFunc {
	return func(c *gin.Context) {
		input := &dto.AnalysisDto{}

		if err := util.ReadRequest(c, input); err != nil {
			log.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "error",
			})
			c.Abort()
			return
		}

		result, err := rh.researchService.GetNumberOfMoviesByReleaseDate(c, input.Year, input.Months)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			c.Abort()
			return
		}
		c.JSON(http.StatusOK, result)
	}
}

func (rh *researchHandler) FetchNumberOfMoviesByCreatedDate() gin.HandlerFunc {
	return func(c *gin.Context) {
		input := &dto.AnalysisDto{}

		if err := util.ReadRequest(c, input); err != nil {
			log.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "error",
			})
			c.Abort()
			return
		}

		result, err := rh.researchService.GetNumberOfMoviesByCreatedDate(c, input.Year, input.Months)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			c.Abort()
			return
		}
		c.JSON(http.StatusOK, result)
	}
}

func (rh *researchHandler) FetchViewsByGenreAndViewedDate() gin.HandlerFunc {
	return func(c *gin.Context) {
		var err error

		input := &dto.RequestData{}
		if err = util.ReadRequest(c, input); err != nil {
			log.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "error",
			})
			c.Abort()
			return
		}

		var result *dto.ResultDto
		if input.IsCumulative {
			result, err = rh.researchService.GetCumulativeViewsByGenreAndViewedDate(c, input)
		} else {
			result, err = rh.researchService.GetNumberOfViewsByGenreAndViewedDate(c, input)
		}
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			c.Abort()
			return
		}
		c.JSON(http.StatusOK, result)
	}
}

func (rh *researchHandler) FetchNumberOfViewsByViewedDate() gin.HandlerFunc {
	return func(c *gin.Context) {
		request := &dto.RequestData{}
		if err := util.ReadRequest(c, request); err != nil {
			log.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "error",
			})
			c.Abort()
			return
		}

		result, err := rh.researchService.GetNumberOfViewsByViewedDate(c, request)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			c.Abort()
			return
		}
		c.JSON(http.StatusOK, result)
	}
}

func (rh *researchHandler) FetchNumberOfMoviesByGenreAndReleaseDate() gin.HandlerFunc {
	return func(c *gin.Context) {
		input := &dto.RequestData{}

		if err := util.ReadRequest(c, input); err != nil {
			log.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "error",
			})
			c.Abort()
			return
		}

		result, err := rh.researchService.GetNumberOfMoviesByGenreAndReleasedDate(c, input)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			c.Abort()
			return
		}
		c.JSON(http.StatusOK, result)
	}
}

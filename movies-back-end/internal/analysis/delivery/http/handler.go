package http

import (
	"github.com/gin-gonic/gin"
	"log"
	"movies-service/internal/analysis"
	"movies-service/internal/dto"
	"movies-service/pkg/utils"
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
		result, err := rh.researchService.GetNumberOfMoviesByGenre(c.Request.Context())
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

		if err := utils.ReadRequest(c, input); err != nil {
			log.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "error",
			})
			c.Abort()
			return
		}

		result, err := rh.researchService.GetNumberOfMoviesByReleaseDate(c.Request.Context(), input.Year, input.Months)
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

		if err := utils.ReadRequest(c, input); err != nil {
			log.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "error",
			})
			c.Abort()
			return
		}

		result, err := rh.researchService.GetNumberOfMoviesByCreatedDate(c.Request.Context(), input.Year, input.Months)
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

func (rh *researchHandler) FetchNumberOfViewsByGenreAndViewedDate() gin.HandlerFunc {
	return func(c *gin.Context) {
		input := &dto.AnalysisDto{}

		if err := utils.ReadRequest(c, input); err != nil {
			log.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "error",
			})
			c.Abort()
			return
		}

		result, err := rh.researchService.GetNumberOfViewsByGenreAndViewedDate(c.Request.Context(), input.Genre,
			input.Year, input.Months)
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

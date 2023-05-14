package http

import (
	"github.com/gin-gonic/gin"
	"movies-service/internal/dto"
	"movies-service/internal/views"
	"movies-service/pkg/utils"
	"net/http"
	"strconv"
)

type viewHandler struct {
	viewService views.Service
}

func NewViewHandler(viewService views.Service) views.Handler {
	return &viewHandler{
		viewService: viewService,
	}
}

func (h *viewHandler) RecognizeViewForMovie() gin.HandlerFunc {
	return func(c *gin.Context) {
		viewer := &dto.Viewer{}

		if err := utils.ReadRequest(c, viewer); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "error",
			})
			c.Abort()
			return
		}

		err := h.viewService.RecognizeViewForMovie(c, viewer)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": err.Error(),
			})
			c.Abort()
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "OK"})
	}
}

func (h *viewHandler) FetchNumberOfViewsByMovieId() gin.HandlerFunc {
	return func(c *gin.Context) {
		movieKey := c.Param("movieId")
		movieId, _ := strconv.Atoi(movieKey)

		totalViews, err := h.viewService.GetNumberOfViewsByMovieId(c, movieId)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": err.Error(),
			})
			c.Abort()
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "OK",
			"views":   totalViews,
		})
	}
}

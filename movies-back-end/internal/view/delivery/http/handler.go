package http

import (
	"github.com/gin-gonic/gin"
	"log"
	"movies-service/internal/common/dto"
	"movies-service/internal/view"
	"movies-service/pkg/util"
	"net/http"
	"strconv"
)

type viewHandler struct {
	viewService view.Service
}

func NewViewHandler(viewService view.Service) view.Handler {
	return &viewHandler{
		viewService: viewService,
	}
}

func (h *viewHandler) RecognizeViewForMovie() gin.HandlerFunc {
	return func(c *gin.Context) {
		viewer := &dto.Viewer{}
		if err := util.ReadRequest(c, viewer); err != nil {
			log.Println(err)
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

		totalViews, err := h.viewService.GetNumberOfViewsByMovieId(c, uint(movieId))
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

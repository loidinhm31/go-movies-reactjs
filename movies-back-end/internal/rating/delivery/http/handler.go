package http

import (
	"github.com/gin-gonic/gin"
	"movies-service/internal/rating"
	"net/http"
)

type ratingHandler struct {
	ratingService rating.Service
}

func NewRatingHandler(ratingService rating.Service) rating.Handler {
	return &ratingHandler{
		ratingService: ratingService,
	}
}

func (rh *ratingHandler) FetchRatings() gin.HandlerFunc {
	return func(c *gin.Context) {
		allRatings, err := rh.ratingService.GetAllRatings(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			c.Abort()
			return
		}
		c.JSON(http.StatusOK, allRatings)
	}
}

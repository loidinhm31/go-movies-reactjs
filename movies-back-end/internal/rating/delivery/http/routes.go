package http

import (
	"github.com/gin-gonic/gin"
	"movies-service/internal/rating"
)

// MapRatingRoutes Map auth routes
func MapRatingRoutes(ratingGroup *gin.RouterGroup, h rating.Handler) {
	ratingGroup.GET("/", h.FetchRatings())
}

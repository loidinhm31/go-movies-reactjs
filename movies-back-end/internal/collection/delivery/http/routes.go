package http

import (
	"github.com/gin-gonic/gin"
	"movies-service/internal/collection"
)

// MapCollectionRoutes Map collection routes
func MapCollectionRoutes(collectionGroup *gin.RouterGroup, h collection.Handler) {
	collectionGroup.GET("/", h.FetchCollectionByUsernameAndRefID())
}

func MapAuthCollectionRoutes(collectionGroup *gin.RouterGroup, h collection.Handler) {
	collectionGroup.PUT("/", h.PutCollection())
	collectionGroup.GET("/", h.FetchCollectionsByUsername())
	collectionGroup.POST("/", h.FetchCollectionsByUsername())
}

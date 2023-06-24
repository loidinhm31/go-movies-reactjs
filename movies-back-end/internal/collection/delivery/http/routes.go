package http

import (
	"github.com/gin-gonic/gin"
	"movies-service/internal/collection"
)

// MapCollectionRoutes Map collection routes
func MapCollectionRoutes(collectionGroup *gin.RouterGroup, h collection.Handler) {
	collectionGroup.GET("/", h.FetchCollectionByUserAndRefID())
}

func MapAuthCollectionRoutes(collectionGroup *gin.RouterGroup, h collection.Handler) {
	collectionGroup.POST("/", h.PostCollection())
	collectionGroup.GET("/page", h.FetchCollectionsByUsername())
	collectionGroup.POST("/page", h.FetchCollectionsByUsername())
	collectionGroup.DELETE("/refs/:id", h.DeleteCollectionByTypeCodeAndRefID())
}

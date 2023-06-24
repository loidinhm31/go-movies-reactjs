package collection

import "github.com/gin-gonic/gin"

type Handler interface {
	PostCollection() gin.HandlerFunc
	FetchCollectionsByUsername() gin.HandlerFunc
	FetchCollectionByUserAndRefID() gin.HandlerFunc
	DeleteCollectionByTypeCodeAndRefID() gin.HandlerFunc
}

package collection

import "github.com/gin-gonic/gin"

type Handler interface {
	PutCollection() gin.HandlerFunc
	FetchCollectionsByUsername() gin.HandlerFunc
	FetchCollectionByUsernameAndRefID() gin.HandlerFunc
}

package http

import (
	"github.com/gin-gonic/gin"
	"log"
	"movies-service/internal/collection"
	"movies-service/internal/dto"
	"movies-service/pkg/pagination"
	"movies-service/pkg/util"
	"net/http"
	"strconv"
)

type collectionHandler struct {
	collectionService collection.Service
}

func NewCollectionHandler(collectionService collection.Service) collection.Handler {
	return &collectionHandler{
		collectionService: collectionService,
	}
}

func (r collectionHandler) PutCollection() gin.HandlerFunc {
	return func(c *gin.Context) {
		theCollection := &dto.CollectionDto{}

		if err := util.ReadRequest(c, theCollection); err != nil {
			log.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "error",
			})
			c.Abort()
			return
		}

		err := r.collectionService.AddCollection(c, theCollection)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			c.Abort()
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "ok",
		})
	}
}

func (r collectionHandler) FetchCollectionsByUsername() gin.HandlerFunc {
	return func(c *gin.Context) {
		keyword := c.Query("q")
		movieType := c.Query("type")

		pageable, _ := pagination.ReadPageRequest(c)

		if err := util.ReadRequest(c, pageable); err != nil {
			log.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "error",
			})
			c.Abort()
			return
		}
		results, err := r.collectionService.GetCollectionsByUsernameAndType(c, movieType, keyword, pageable)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			c.Abort()
			return
		}
		c.JSON(http.StatusOK, results)
	}
}

func (r collectionHandler) FetchCollectionByUsernameAndRefID() gin.HandlerFunc {
	return func(c *gin.Context) {
		username := c.Query("username")
		typeCode := c.Query("type")
		refIdStr := c.Query("refId")
		refID, _ := strconv.Atoi(refIdStr)

		results, err := r.collectionService.GetCollectionByUsernameAndRefID(c, username, typeCode, uint(refID))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			c.Abort()
			return
		}
		c.JSON(http.StatusOK, results)
	}
}

package http

import (
	"github.com/gin-gonic/gin"
	"log"
	"movies-service/internal/model"
	"movies-service/internal/search"
	"movies-service/pkg/util"
	"net/http"
)

type graphHandler struct {
	graphService search.Service
}

func NewGraphHandler(graphService search.Service) search.Handler {
	return &graphHandler{
		graphService: graphService,
	}
}

func (gh *graphHandler) Search() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the query from the request
		filter := &model.SearchParams{}

		if err := util.ReadRequest(c, filter); err != nil {
			log.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "error",
			})
			c.Abort()
			return
		}

		result, err := gh.graphService.SearchMovie(c, filter)
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

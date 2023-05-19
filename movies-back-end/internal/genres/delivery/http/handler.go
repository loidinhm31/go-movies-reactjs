package http

import (
	"github.com/gin-gonic/gin"
	"log"
	"movies-service/internal/genres"
	"movies-service/pkg/utils"
	"net/http"
)

type genreHandler struct {
	genreService genres.Service
}

func NewGenreHandler(genreService genres.Service) genres.GenreHandler {
	return &genreHandler{
		genreService: genreService,
	}
}

func (mh *genreHandler) FetchGenres() gin.HandlerFunc {
	return func(c *gin.Context) {
		allGenres, err := mh.genreService.GetAllGenres(c)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			c.Abort()
			return
		}
		c.JSON(http.StatusOK, allGenres)
	}
}

func (mh *genreHandler) PostGenres() gin.HandlerFunc {
	return func(c *gin.Context) {
		type Request struct {
			Genres []string `json:"genres"`
		}

		request := &Request{}
		if err := utils.ReadRequest(c, request); err != nil {
			log.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "error",
			})
			c.Abort()
			return
		}

		err := mh.genreService.AddGenres(c, request.Genres)
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err,
			})
			c.Abort()
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "ok",
		})
	}
}

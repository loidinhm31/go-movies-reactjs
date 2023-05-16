package http

import (
	"github.com/gin-gonic/gin"
	"log"
	"movies-service/internal/dto"
	"movies-service/internal/users"
	"movies-service/pkg/pagination"
	"movies-service/pkg/utils"
	"net/http"
)

type userHandler struct {
	userService users.Service
}

func NewUserHandler(userService users.Service) users.Handler {
	return &userHandler{
		userService: userService,
	}
}

func (u userHandler) FetchUsers() gin.HandlerFunc {
	return func(c *gin.Context) {
		searchKey := c.Query("q")
		isNew := c.Query("isNew")

		pageable, _ := pagination.ReadPageRequest(c)
		if err := utils.ReadRequest(c, pageable); err != nil {
			log.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "error",
			})
			c.Abort()
			return
		}

		allUsers, err := u.userService.GetUsers(c, pageable, searchKey, isNew)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			c.Abort()
			return
		}
		c.JSON(http.StatusOK, allUsers)
	}
}

func (u userHandler) PatchUserRole() gin.HandlerFunc {
	return func(c *gin.Context) {
		user := &dto.UserDto{}
		if err := utils.ReadRequest(c, user); err != nil {
			log.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "error",
			})
			c.Abort()
			return
		}

		err := u.userService.UpdateUserRole(c, user)
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

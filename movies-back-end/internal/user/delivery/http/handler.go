package http

import (
	"github.com/gin-gonic/gin"
	"log"
	"movies-service/internal/common/dto"
	"movies-service/internal/user"
	"movies-service/pkg/pagination"
	"movies-service/pkg/util"
	"net/http"
	"strconv"
)

type userHandler struct {
	userService user.Service
}

func NewUserHandler(userService user.Service) user.Handler {
	return &userHandler{
		userService: userService,
	}
}

func (u userHandler) FetchUsers() gin.HandlerFunc {
	return func(c *gin.Context) {
		searchKey := c.Query("q")
		isNew := c.Query("isNew")
		isNewBool, _ := strconv.ParseBool(isNew)

		pageable, _ := pagination.ReadPageRequest(c)
		if err := util.ReadRequest(c, pageable); err != nil {
			log.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "error",
			})
			c.Abort()
			return
		}

		allUsers, err := u.userService.GetUsers(c, pageable, searchKey, isNewBool)
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
		if err := util.ReadRequest(c, user); err != nil {
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

func (u userHandler) PutOidcUSer() gin.HandlerFunc {
	return func(c *gin.Context) {
		user := &dto.UserDto{}
		if err := util.ReadRequest(c, user); err != nil {
			log.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "error",
			})
			c.Abort()
			return
		}

		_, err := u.userService.AddOidcUser(c, user)
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

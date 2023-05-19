package http

import (
	"github.com/gin-gonic/gin"
	"movies-service/internal/roles"
	"net/http"
)

type roleHandler struct {
	roleService roles.Service
}

func NewRoleHandler(roleService roles.Service) roles.Handler {
	return &roleHandler{
		roleService: roleService,
	}
}

func (r roleHandler) FetchRoles() gin.HandlerFunc {
	return func(c *gin.Context) {
		allRoles, err := r.roleService.GetAllRoles(c)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			c.Abort()
			return
		}
		c.JSON(http.StatusOK, allRoles)
	}
}

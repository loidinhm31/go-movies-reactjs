package util

import (
	"github.com/gin-gonic/gin"
)

// GetConfigPath Get config path for config yml file
func GetConfigPath(profile string) string {
	if profile == "docker" {
		return "config-docker"
	}
	return "config-local"
}

// ReadRequest Read request body and validate
func ReadRequest(ctx *gin.Context, request interface{}) error {
	if err := ctx.Bind(request); err != nil {
		return err
	}
	return validate.StructCtx(ctx.Request.Context(), request)
}

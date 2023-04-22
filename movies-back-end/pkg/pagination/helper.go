package pagination

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

func ReadPageRequest(ctx *gin.Context) (*PageRequest, error) {
	pageNumber := 0
	pageNumStr := ctx.Query("page")
	if pageNumStr != "" {
		pageNumber, _ = strconv.Atoi(pageNumStr)
	}

	pageSize := 5
	pageSizeStr := ctx.Query("size")
	if pageSizeStr != "" {
		pageSize, _ = strconv.Atoi(pageSizeStr)
	}

	pageable := PageRequestOf().Set(pageNumber, pageSize)
	return pageable, nil
}

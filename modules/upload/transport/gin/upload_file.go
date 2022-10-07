package uploadgin

import (
	"fmt"
	"food-delivery/common"
	component "food-delivery/components"

	"github.com/gin-gonic/gin"
)

func UploadHandler(appCtx component.AppContext) func(*gin.Context) {
	return func(c *gin.Context) {
		fileHeader, err := c.FormFile("file")

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		if err := c.SaveUploadedFile(fileHeader, fmt.Sprintf("./static/%s", fileHeader.Filename)); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		c.JSON(200, common.SimpleSuccessResponse(true))
	}
}
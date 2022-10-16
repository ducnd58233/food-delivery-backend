package usergin

import (
	"food-delivery/common"
	component "food-delivery/components"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetProfileHandler(appCtx component.AppContext) func(*gin.Context) {
	return func(c *gin.Context) {
		data := c.MustGet(common.CurrentUser).(common.Requester)

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
	}
}

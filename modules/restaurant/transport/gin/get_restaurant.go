package restaurantgin

import (
	"food-delivery/common"
	"food-delivery/components"
	restaurantbiz "food-delivery/modules/restaurant/biz"
	restaurantstorage "food-delivery/modules/restaurant/storage"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetRestaurantHandler(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := restaurantstorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := restaurantbiz.NewGetRestaurantBiz(store)

		data, err := biz.GetRestaurant(c.Request.Context(), id)

		if err != nil {
			panic(err) // AppCtx error
		}
	
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
	}
}

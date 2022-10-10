package restaurantgin

import (
	"food-delivery/common"
	component "food-delivery/components"
	restaurantbiz "food-delivery/modules/restaurant/biz"
	restaurantstorage "food-delivery/modules/restaurant/storage"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetRestaurantHandler(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		// id, err := strconv.Atoi(c.Param("id"))
		uid, err := common.FromBase58(c.Param("id"))

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := restaurantstorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := restaurantbiz.NewGetRestaurantBiz(store)

		data, err := biz.GetRestaurant(c.Request.Context(), int(uid.GetLocalID()))

		if err != nil {
			panic(err) // AppCtx error
		}

		data.Mask(false)
		
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
	}
}

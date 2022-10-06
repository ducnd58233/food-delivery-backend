package restaurantgin

import (
	"food-delivery/common"
	"food-delivery/components"
	restaurantbiz "food-delivery/modules/restaurant/biz"
	restaurantmodel "food-delivery/modules/restaurant/model"
	restaurantstorage "food-delivery/modules/restaurant/storage"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func UpdateRestaurantHandler(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data restaurantmodel.RestaurantUpdate

		id, err := strconv.Atoi(c.Param("id"))
		// uid, err := common.FromBase58(c.Param("restaurant_id"))

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		if err := c.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		storage := restaurantstorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := restaurantbiz.NewUpdateRestaurantBiz(storage)

		if err := biz.UpdateRestaurantById(c.Request.Context(), id, &data); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}

package restaurantgin

import (
	"food-delivery/common"
	component "food-delivery/components"
	restaurantbiz "food-delivery/modules/restaurant/biz"
	restaurantmodel "food-delivery/modules/restaurant/model"
	restaurantstorage "food-delivery/modules/restaurant/storage"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UpdateRestaurantHandler(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data restaurantmodel.RestaurantUpdate

		// id, err := strconv.Atoi(c.Param("id"))
		uid, err := common.FromBase58(c.Param("id"))

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		if err := c.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		storage := restaurantstorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := restaurantbiz.NewUpdateRestaurantBiz(storage)

		if err := biz.UpdateRestaurantById(c.Request.Context(), int(uid.GetLocalID()), &data); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}

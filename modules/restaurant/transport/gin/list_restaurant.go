package restaurantgin

import (
	"food-delivery/common"
	"food-delivery/components"
	restaurantbiz "food-delivery/modules/restaurant/biz"
	restaurantmodel "food-delivery/modules/restaurant/model"
	restaurantstorage "food-delivery/modules/restaurant/storage"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetListRestaurantsHandler(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var filter restaurantmodel.Filter

		if err := c.ShouldBind(&filter); err != nil {
			c.JSON(401, gin.H{
				"error": err.Error(),
			})
	
			return
		}

		var paging common.Paging

		if err := c.ShouldBind(&paging); err != nil {
			c.JSON(401, gin.H{
				"error": err.Error(),
			})
	
			return
		}

		paging.Fulfill()
	
		store := restaurantstorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := restaurantbiz.NewListRestaurantBiz(store)
	
		result, err := biz.ListRestaurant(c.Request.Context(), &filter, &paging)

		if err != nil {
			c.JSON(401, gin.H{
				"error": err.Error(),
			})
	
			return
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(result, paging, filter))
	}
}

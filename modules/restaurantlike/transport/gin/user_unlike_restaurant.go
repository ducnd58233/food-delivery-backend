package restaurantlikegin

import (
	"food-delivery/common"
	component "food-delivery/components"
	restaurantstorage "food-delivery/modules/restaurant/storage"
	restaurantlikebiz "food-delivery/modules/restaurantlike/biz"
	restaurantlikestorage "food-delivery/modules/restaurantlike/storage"
	"net/http"

	"github.com/gin-gonic/gin"
)

// DELETE /v1/restaurants/:id/unlike

func UserUnlikeRestaurantHandler(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, err := common.FromBase58(c.Param("id"))

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		requester := c.MustGet(common.CurrentUser).(common.Requester)

		store := restaurantlikestorage.NewSQLStore(appCtx.GetMainDBConnection())
		descStore := restaurantstorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := restaurantlikebiz.NewUserUnlikeRestaurantBiz(store, descStore)

		if err := biz.UnlikeRestaurant(c.Request.Context(), requester.GetUserId(), int(uid.GetLocalID())); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}

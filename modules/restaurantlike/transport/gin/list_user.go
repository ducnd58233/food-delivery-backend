package restaurantlikegin

import (
	"food-delivery/common"
	component "food-delivery/components"
	restaurantlikemodel "food-delivery/modules/restaurantlike/model"
	restaurantlikestorage "food-delivery/modules/restaurantlike/storage"
	restaurantlikebiz "food-delivery/modules/restaurantlike/biz"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetListUsersLikeRestaurantHandler(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		// var filter restaurantlikemodel.Filter

		// if err := c.ShouldBind(&filter); err != nil {
		// 	panic(common.ErrInvalidRequest(err))
		// }
		uid, err := common.FromBase58(c.Param("id"))

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		filter := restaurantlikemodel.Filter{
			RestaurantId: int(uid.GetLocalID()),
		}

		var paging common.Paging

		if err := c.ShouldBind(&paging); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		paging.Fulfill()

		store := restaurantlikestorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := restaurantlikebiz.NewListUsersLikeRestaurantBiz(store)

		result, err := biz.ListUsers(c.Request.Context(), &filter, &paging)

		if err != nil {
			panic(err)
		}

		for i := range result {
			result[i].Mask(false)
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(result, paging, filter))
	}
}

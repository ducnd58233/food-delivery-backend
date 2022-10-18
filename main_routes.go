package main

import (
	"food-delivery/common"
	component "food-delivery/components"
	middleware "food-delivery/middlewares"
	restaurantgin "food-delivery/modules/restaurant/transport/gin"
	restaurantlikegin "food-delivery/modules/restaurantlike/transport/gin"
	uploadgin "food-delivery/modules/upload/transport/gin"
	usergin "food-delivery/modules/user/transport/gin"
	"net/http"

	"github.com/gin-gonic/gin"
)

func mainRoute(router *gin.Engine, appCtx component.AppContext) {
	v1 := router.Group("/v1")
	{
		v1.POST("/register", usergin.RegisterHandler(appCtx))
		v1.POST("/login", usergin.LoginHandler(appCtx))
		v1.POST("/profile", middleware.RequiredAuth(appCtx), usergin.GetProfileHandler(appCtx))
		
		v1.POST("/upload", uploadgin.UploadHandler(appCtx))

		restaurants := v1.Group("/restaurants", middleware.RequiredAuth(appCtx))
		{
			restaurants.POST("", restaurantgin.CreateRestaurantHandler(appCtx))
			restaurants.GET("/:id", restaurantgin.GetRestaurantHandler(appCtx))
			restaurants.GET("", restaurantgin.GetListRestaurantsHandler(appCtx))
			restaurants.PATCH("/:id", restaurantgin.UpdateRestaurantHandler(appCtx))
			restaurants.DELETE("/:id", restaurantgin.DeleteRestaurantHandler(appCtx))

			restaurants.GET("/:id/liked-users", restaurantlikegin.GetListUsersLikeRestaurantHandler(appCtx))
			restaurants.POST("/:id/like", restaurantlikegin.UserLikeRestaurantHandler(appCtx))
			restaurants.DELETE("/:id/unlike", restaurantlikegin.UserUnlikeRestaurantHandler(appCtx))
		
		}

		v1.GET("encode-uid", func(c *gin.Context) {
			type reqData struct {
				DbType int `form:"type"`
				RealId int `form:"id"`
			}

			var d reqData
			c.ShouldBind(&d)

			c.JSON(http.StatusOK, gin.H{
				"id": common.NewUID(uint32(d.RealId), d.DbType, 1),
			})
		})
	}
}

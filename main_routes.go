package main

import (
	restaurantgin "food-delivery/modules/restaurant/transport/gin"
	uploadgin "food-delivery/modules/upload/transport/gin"
	"food-delivery/components"

	"github.com/gin-gonic/gin"
)

func mainRoute(router *gin.Engine, appCtx component.AppContext) {
	v1 := router.Group("/v1")
	{
		restaurants := v1.Group("/restaurants")
		{
			restaurants.POST("", restaurantgin.CreateRestaurantHandler(appCtx))
			restaurants.GET("/:id", restaurantgin.GetRestaurantHandler(appCtx))
			restaurants.GET("", restaurantgin.GetListRestaurantsHandler(appCtx))
			restaurants.PATCH("/:id", restaurantgin.UpdateRestaurantHandler(appCtx))
			restaurants.DELETE("/:id", restaurantgin.DeleteRestaurantHandler(appCtx))
		}

		v1.POST("/upload", uploadgin.UploadHandler(appCtx))
	}
}

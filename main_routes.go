package main

import (
	restaurantgin "food-delivery/modules/restaurant/transport/gin"
	"food-delivery/components"

	"github.com/gin-gonic/gin"
)

func mainRoute(router *gin.Engine, appCtx component.AppContext) {
	v1 := router.Group("/v1")
	{
		restaurants := v1.Group("/restaurants")
		{
			restaurants.POST("", restaurantgin.CreateRestaurantHandler(appCtx))
			// restaurants.GET("/:restaurant_id", restaurantgin.GetRestaurantHandler(appCtx))
			restaurants.GET("", restaurantgin.GetListRestaurantsHandler(appCtx))
			// restaurants.PATCH("/:restaurant_id", restaurantgin.UpdateRestaurantHandler(appCtx))
			// restaurants.DELETE("/:restaurant_id", restaurantgin.DeleteRestaurantHandler(appCtx))
		}
	}
}

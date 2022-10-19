package subscriber

import (
	"context"
	"food-delivery/common"
	component "food-delivery/components"
	restaurantstorage "food-delivery/modules/restaurant/storage"
	restaurantlikemodel "food-delivery/modules/restaurantlike/model"
)

func IncreaseLikeCountAfterUserLikeRestaurant(appCtx component.AppContext, ctx context.Context) {
	c, _ := appCtx.GetPubsub().Subscribe(ctx, common.TopicUserLikeRestaurant)
	store := restaurantstorage.NewSQLStore(appCtx.GetMainDBConnection())

	go func() {
		defer common.AppRecover()
		for {
			msg := <-c
			likeData := msg.Data().(*restaurantlikemodel.Like)
			_ = store.IncreaseLikeCount(ctx, likeData.RestaurantId)
		}
	}()
}

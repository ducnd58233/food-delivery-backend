package subscriber

import (
	"context"
	component "food-delivery/components"
	restaurantstorage "food-delivery/modules/restaurant/storage"
	"food-delivery/pubsub"
)

func RunDecreaseLikeCountAfterUserUnlikeRestaurant(appCtx component.AppContext) consumerJob {
	return consumerJob{
		Title: "Decrease like count after user unlikes restaurant",
		Hld: func(ctx context.Context, message *pubsub.Message) error {
			store := restaurantstorage.NewSQLStore(appCtx.GetMainDBConnection())
			likeData := message.Data().(HasRestaurantId)
			return store.DecreaseLikeCount(ctx, likeData.GetRestaurantId())
		},
	}
}

package subscriber

import (
	"context"
	component "food-delivery/components"
	restaurantstorage "food-delivery/modules/restaurant/storage"
	"food-delivery/pubsub"
	"food-delivery/skio"
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

func EmitRealtimeAfterUserUnlikeRestaurant(appCtx component.AppContext, rtEngine skio.RealtimeEngine) consumerJob {
	return consumerJob {
		Title: "Emit realtime decrease like count after user likes restaurant",
		Hld: func(ctx context.Context, message *pubsub.Message) error {
			likeData := message.Data().(HasRestaurantId)

			return rtEngine.EmitToUser(likeData.GetRestaurantId(), string(message.Channel()), likeData)
		},
	}
}
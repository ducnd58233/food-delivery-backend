package subscriber

import (
	"context"
	component "food-delivery/components"
	restaurantstorage "food-delivery/modules/restaurant/storage"
	"food-delivery/pubsub"
	"food-delivery/skio"
)

type HasRestaurantId interface {
	GetRestaurantId() int
	GetUserLikeId() int
}

// func IncreaseLikeCountAfterUserLikeRestaurant(appCtx component.AppContext, ctx context.Context) {
// 	c, _ := appCtx.GetPubsub().Subscribe(ctx, common.TopicUserLikeRestaurant)
// 	store := restaurantstorage.NewSQLStore(appCtx.GetMainDBConnection())

// 	go func() {
// 		defer common.AppRecover()
// 		for {
// 			msg := <-c
// 			likeData := msg.Data().(HasRestaurantId) // Convert to HasRestaurantId type
// 			_ = store.IncreaseLikeCount(ctx, likeData.GetRestaurantId())
// 		}
// 	}()
// }

// Wish do something like this
// func RunIncreaseLikeCountAfterUserLikeRestaurant(appCtx component.AppContext) func(ctx context.Context, message *pubsub.Message) error {
// 	store := restaurantstorage.NewSQLStore(appCtx.GetMainDBConnection())

// 	return func(ctx context.Context, message *pubsub.Message) error {
// 		likeData := message.Data().(HasRestaurantId)
// 		return store.IncreaseLikeCount(ctx, likeData.GetRestaurantId())
// 	}
// }

// Convert from above implement to this implement - (SDK implement mindset)
func RunIncreaseLikeCountAfterUserLikeRestaurant(appCtx component.AppContext) consumerJob {
	return consumerJob{
		Title: "Increase like count after user likes restaurant",
		Hld: func(ctx context.Context, message *pubsub.Message) error {
			store := restaurantstorage.NewSQLStore(appCtx.GetMainDBConnection())
			likeData := message.Data().(HasRestaurantId)

			return store.IncreaseLikeCount(ctx, likeData.GetRestaurantId())
		},
	}
}

func EmitRealtimeAfterUserLikeRestaurant(appCtx component.AppContext, rtEngine skio.RealtimeEngine) consumerJob {
	return consumerJob {
		Title: "Emit realtime increase like count after user likes restaurant",
		Hld: func(ctx context.Context, message *pubsub.Message) error {
			likeData := message.Data().(HasRestaurantId)

			return rtEngine.EmitToUser(likeData.GetUserLikeId(), string(message.Channel()), likeData)
		},
	}
}
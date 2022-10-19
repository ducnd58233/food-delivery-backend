package restaurantlikebiz

import (
	"context"
	"food-delivery/common"
	restaurantlikemodel "food-delivery/modules/restaurantlike/model"
	"food-delivery/pubsub"
)

type UserLikeRestaurantStore interface {
	FindDataByCondition(
		ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string,
	) (*restaurantlikemodel.Like, error)
	Create(ctx context.Context, data *restaurantlikemodel.Like) error
}

// type IncreaseLikeCountStore interface {
// 	IncreaseLikeCount(ctx context.Context, id int) error
// }

type userLikeRestaurantBiz struct {
	store UserLikeRestaurantStore
	// incStore IncreaseLikeCountStore
	pubsub pubsub.Pubsub
}

func NewUserLikeRestaurantBiz(
	store UserLikeRestaurantStore,
	// incStore IncreaseLikeCountStore,
	pubsub pubsub.Pubsub,
) *userLikeRestaurantBiz {
	return &userLikeRestaurantBiz{
		store: store,
		// incStore: incStore,
		pubsub: pubsub,
	}
}

func (biz *userLikeRestaurantBiz) LikeRestaurant(
	ctx context.Context,
	data *restaurantlikemodel.Like,
) error {
	likeExist, _ := biz.store.FindDataByCondition(ctx, map[string]interface{}{"user_id": data.UserId, "restaurant_id": data.RestaurantId})

	if likeExist != nil {
		return common.NewCustomError(nil, "user already like restaurant", "AlreadyLikeRestaurant")
	}

	err := biz.store.Create(ctx, data)

	if err != nil {
		return restaurantlikemodel.ErrCannotLikeRestaurant(err)
	}

	// side effect
	// go func() {
	// 	defer common.AppRecover()

	// 	job := asyncjob.NewJob(func(ctx context.Context) error {
	// 		return biz.incStore.IncreaseLikeCount(ctx, data.RestaurantId)
	// 	})

	// 	_ = asyncjob.NewGroup(true, job).Run(ctx)
	// }()
	// New solution: use pubsub
	// Do not inject directly here, hard to unit test. Inject through struct instead
	biz.pubsub.Publish(ctx, common.TopicUserLikeRestaurant, pubsub.NewMessage(data))

	return nil
}

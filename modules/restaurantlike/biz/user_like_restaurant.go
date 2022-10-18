package restaurantlikebiz

import (
	"context"
	"food-delivery/common"
	"food-delivery/components/asyncjob"
	restaurantlikemodel "food-delivery/modules/restaurantlike/model"
)

type UserLikeRestaurantStore interface {
	FindDataByCondition(
		ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string,
	) (*restaurantlikemodel.Like, error)
	Create(ctx context.Context, data *restaurantlikemodel.Like) error
}

type IncreaseLikeCountStore interface {
	IncreaseLikeCount(ctx context.Context, id int) error
}

type userLikeRestaurantBiz struct {
	store    UserLikeRestaurantStore
	incStore IncreaseLikeCountStore
}

func NewUserLikeRestaurantBiz(store UserLikeRestaurantStore, incStore IncreaseLikeCountStore) *userLikeRestaurantBiz {
	return &userLikeRestaurantBiz{store: store, incStore: incStore}
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
	go func() {
		defer common.AppRecover()
		
		job := asyncjob.NewJob(func(ctx context.Context) error {
			return biz.incStore.IncreaseLikeCount(ctx, data.RestaurantId)
		})
	
		_ = asyncjob.NewGroup(true, job).Run(ctx)
	}()

	return nil
}

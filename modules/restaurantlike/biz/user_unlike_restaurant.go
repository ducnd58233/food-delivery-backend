package restaurantlikebiz

import (
	"context"
	"food-delivery/common"
	restaurantlikemodel "food-delivery/modules/restaurantlike/model"
)

type UserUnlikeRestaurantStore interface {
	FindDataByCondition(
		ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string,
	) (*restaurantlikemodel.Like, error)
	Delete(ctx context.Context, userId, restaurantId int) error
}

type DecreaseLikeCountStore interface {
	DecreaseLikeCount(ctx context.Context, id int) error
}

type userUnlikeRestaurantBiz struct {
	store     UserUnlikeRestaurantStore
	descStore DecreaseLikeCountStore
}

func NewUserUnlikeRestaurantBiz(store UserUnlikeRestaurantStore, descStore DecreaseLikeCountStore) *userUnlikeRestaurantBiz {
	return &userUnlikeRestaurantBiz{store: store, descStore: descStore}
}

func (biz *userUnlikeRestaurantBiz) UnlikeRestaurant(
	ctx context.Context,
	userId,
	restaurantId int,
) error {
	_, err := biz.store.FindDataByCondition(ctx, map[string]interface{}{"user_id": userId, "restaurant_id": restaurantId})

	if err != nil {
		return common.NewCustomError(nil, "user has not like restaurant", "NotLikeRestaurant")
	}

	if err := biz.store.Delete(ctx, userId, restaurantId); err != nil {
		return restaurantlikemodel.ErrCannotUnlikeRestaurant(err)
	}

	// side effect
	go func() {
		defer common.AppRecover()
		_ = biz.descStore.DecreaseLikeCount(ctx, restaurantId)
	}()

	return nil
}

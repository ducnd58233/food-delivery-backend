package restaurantbiz

import (
	"context"
	"food-delivery/common"
	restaurantmodel "food-delivery/modules/restaurant/model"
)

type ListRestaurantRepo interface {
	ListRestaurant(
		ctx context.Context,
		filter *restaurantmodel.Filter,
		paging *common.Paging,
	) ([]restaurantmodel.Restaurant, error)
}

type listRestaurantBiz struct {
	repo ListRestaurantRepo
}

func NewListRestaurantBiz(repo ListRestaurantRepo) *listRestaurantBiz {
	return &listRestaurantBiz{repo: repo}
}

func (biz *listRestaurantBiz) ListRestaurant(
	ctx context.Context,
	filter *restaurantmodel.Filter,
	paging *common.Paging,
) ([]restaurantmodel.Restaurant, error) {
	result, err := biz.repo.ListRestaurant(ctx, filter, paging)

	if err != nil {
		return nil, common.ErrCannotListEntity(restaurantmodel.EntityName, err)
	}

	return result, nil
}

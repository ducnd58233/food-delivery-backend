package restaurantlikestorage

import (
	"context"
	"food-delivery/common"
	restaurantlikemodel "food-delivery/modules/restaurantlike/model"
)

type sqlData struct {
	RestaurantId int `gorm:"column:restaurant_id;"`
	LikeCount    int `gorm:"column:count;"`
}

func (s *sqlStore) GetRestaurantLikes(
	ctx context.Context,
	ids []int,
) (map[int]int, error) {
	result := make(map[int]int)

	var listLike []sqlData

	if err := s.db.Table(restaurantlikemodel.Like{}.TableName()).Select("restaurant_id, count(restaurant_id) as count").
		Where("restaurant_id in (?)", ids).
		Group("restaurant_id").
		Find(&listLike).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	for _, item := range listLike {
		result[item.RestaurantId] = item.LikeCount
	}

	return result, nil
}

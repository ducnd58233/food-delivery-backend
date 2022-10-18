package restaurantstorage

import (
	"context"
	"food-delivery/common"
	restaurantmodel "food-delivery/modules/restaurant/model"

	"gorm.io/gorm"
)

func (s *sqlStore) Update(
	ctx context.Context,
	id int,
	data *restaurantmodel.RestaurantUpdate,
) error {
	db := s.db

	if err := db.Where("id = ?", id).Updates(data).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}

func (s *sqlStore) IncreaseLikeCount(
	ctx context.Context,
	id int,
) error {
	db := s.db.Table(restaurantmodel.Restaurant{}.TableName())

	if err := db.Where("id = ?", id).
		Update("liked_count", gorm.Expr("liked_count + ?", 1)).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}

func (s *sqlStore) DecreaseLikeCount(
	ctx context.Context,
	id int,
) error {
	db := s.db.Table(restaurantmodel.Restaurant{}.TableName())

	if err := db.Where("id = ?", id).
		Update("liked_count", gorm.Expr("liked_count - ?", 1)).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}

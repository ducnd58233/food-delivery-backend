package uploadstorage

import (
	"context"
	"food-delivery/common"
)

func (store *sqlStore) List(
	ctx context.Context,
	ids []int,
	moreKeys ...string,
) ([]common.Image, error) {
	db := store.db.Table(common.Image{}.TableName())
	var result []common.Image

	db = db.Table(common.Image{}.TableName())

	if err := db.Where("id in (?)", ids).Find(&result).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return result, nil
}
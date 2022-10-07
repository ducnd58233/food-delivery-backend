package uploadstorage

import (
	"context"
	"food-delivery/common"
)

func (store *sqlStore) Delete(ctx context.Context, ids []int) error {
	db := store.db.Table(common.Image{}.TableName())

	if err := db.
		Where("id in (?)", ids).
		Delete(nil).Error; err != nil {
		return err
	}

	return nil
}

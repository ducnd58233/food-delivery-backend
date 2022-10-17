package userbiz

import (
	"context"
	"food-delivery/common"
	usermodel "food-delivery/modules/user/model"
)

type RegisterStore interface {
	Find(ctx context.Context, conditions map[string]interface{}, moreInfo ...string) (*usermodel.User, error)
	Create(ctx context.Context, data *usermodel.UserCreate) error
}

// use interface to test
type Hasher interface {
	Hash(data string) string
}

type registerBiz struct {
	registerStore RegisterStore
	hasher        Hasher
}

func NewRegisterBiz(registerStore RegisterStore, hasher Hasher) *registerBiz {
	return &registerBiz{
		registerStore: registerStore,
		hasher:        hasher,
	}
}

func (biz *registerBiz) Register(ctx context.Context, data *usermodel.UserCreate) error {
	user, err := biz.registerStore.Find(ctx, map[string]interface{}{"email": data.Email})

	if user != nil {
		return usermodel.ErrEnailExisted
	}

	if err != nil && err == common.ErrRecordNotFound {
		salt := common.GenSalt(50)

		data.Password = biz.hasher.Hash(data.Password + salt)
		data.Salt = salt
		data.Role = "user" // hard code
		data.Status = 1

		if err := biz.registerStore.Create(ctx, data); err != nil {
			return common.ErrCannotCreateEntity(usermodel.EntityName, err)
		}

		return nil
	}

	return common.ErrDB(err)
}

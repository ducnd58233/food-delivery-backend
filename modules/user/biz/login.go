package userbiz

import (
	"context"
	"food-delivery/common"
	component "food-delivery/components"
	"food-delivery/components/tokenprovider"
	usermodel "food-delivery/modules/user/model"
)

type LoginStore interface {
	Find(ctx context.Context, conditions map[string]interface{}, moreInfo ...string) (*usermodel.User, error)
}

// type TokenConfig interface {
// 	GetAtExp() int
// 	GetRtExp() int
// }

type loginBiz struct {
	appCtx        component.AppContext
	userStore     LoginStore
	tokenProvider tokenprovider.Provider
	hasher        Hasher
	expiry        int
	// tkCfg         TokenConfig
}

func NewLoginBiz(userStore LoginStore, tokenProvider tokenprovider.Provider, hasher Hasher, expiry int) *loginBiz {
	return &loginBiz{
		userStore:     userStore,
		tokenProvider: tokenProvider,
		hasher:        hasher,
		expiry:        expiry,
		// tkCfg:         tkCfg,
	}
}

/*
1. Find user, email
2. Hash pass from input and compare with pass in db
3. Provider: issue JWT token for client
3.1. Access token and refresh token
4. Return token(s)
*/
func (biz *loginBiz) Login(ctx context.Context, data *usermodel.UserLogin) (*usermodel.Account, error) {
	user, err := biz.userStore.Find(ctx, map[string]interface{}{"email": data.Email})

	if err != nil {
		return nil, usermodel.ErrUsernameOrPasswordInvalid
	}

	hashedPassword := biz.hasher.Hash(data.Password + user.Salt)

	if user.Password != hashedPassword {
		return nil, usermodel.ErrUsernameOrPasswordInvalid
	}

	payload := tokenprovider.TokenPayload{
		UserId: user.Id,
		Role:   user.Role,
	}

	accessToken, err := biz.tokenProvider.Generate(payload, biz.expiry)
	if err != nil {
		return nil, common.ErrInternal(err)
	}

	refreshToken, err := biz.tokenProvider.Generate(payload, biz.expiry)
	if err != nil {
		return nil, common.ErrInternal(err)
	}

	account := usermodel.NewAccount(accessToken, refreshToken)

	return account, nil
}

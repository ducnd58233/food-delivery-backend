package usergin

import (
	"food-delivery/common"
	component "food-delivery/components"
	"food-delivery/components/hasher"
	"food-delivery/components/tokenprovider/jwt"
	userbiz "food-delivery/modules/user/biz"
	usermodel "food-delivery/modules/user/model"
	userstorage "food-delivery/modules/user/storage"
	"net/http"

	"github.com/gin-gonic/gin"
)

func LoginHandler(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var loginUserData usermodel.UserLogin

		if err := c.ShouldBind(&loginUserData); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		db := appCtx.GetMainDBConnection()
		tokenProvider := jwt.NewTokenJWTProvider(appCtx.SecretKey())

		store := userstorage.NewSQLStore(db)
		md5 := hasher.NewMd5Hash()

		biz := userbiz.NewLoginBiz(store, tokenProvider, md5, 60 * 60 * 24 * 7)
		account, err := biz.Login(c.Request.Context(), &loginUserData)

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(account))
	}
}

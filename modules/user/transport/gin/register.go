package usergin

import (
	"food-delivery/common"
	component "food-delivery/components"
	"food-delivery/components/hasher"
	userbiz "food-delivery/modules/user/biz"
	usermodel "food-delivery/modules/user/model"
	userstorage "food-delivery/modules/user/storage"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterHandler(appCtx component.AppContext) func(*gin.Context) {
	return func(c *gin.Context) {
		var data usermodel.UserCreate

		db := appCtx.GetMainDBConnection()

		if err := c.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := userstorage.NewSQLStore(db)
		md5 := hasher.NewMd5Hash()
		biz := userbiz.NewRegisterBiz(store, md5)

		if err := biz.Register(c.Request.Context(), &data); err != nil {
			panic(err)
		}

		data.Mask(false)

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data.FakeId.String()))
	}
}

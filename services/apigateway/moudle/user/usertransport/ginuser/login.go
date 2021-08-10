package ginuser

import (
	"github.com/chibao13/news_feed_practice/common"
	"github.com/chibao13/news_feed_practice/component/appctx"
	"github.com/chibao13/news_feed_practice/component/hasher"
	"github.com/chibao13/news_feed_practice/component/tokenprovider/jwt"
	"github.com/chibao13/news_feed_practice/services/apigateway/moudle/user/userbusiness"
	"github.com/chibao13/news_feed_practice/services/apigateway/moudle/user/usermodel"
	"github.com/chibao13/news_feed_practice/services/apigateway/moudle/user/userstorage"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var loginUserData usermodel.UserLogin

		if err := c.ShouldBind(&loginUserData); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		db := appCtx.GetMainDBConnection()
		tokenProvider := jwt.NewTokenJWTProvider(appCtx.SecretKey())

		store := userstorage.NewSQLStore(db)
		md5 := hasher.NewMd5Hash()

		business := userbusiness.NewLoginBusiness(store, tokenProvider, md5, appctx.NewTokenConfig())
		account, err := business.Login(c.Request.Context(), &loginUserData)

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(account))
	}
}

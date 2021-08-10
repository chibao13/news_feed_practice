package ginpost

import (
	api "github.com/chibao13/news_feed_practice/api/v1/post"
	"github.com/chibao13/news_feed_practice/common"
	"github.com/chibao13/news_feed_practice/component/appctx"
	"github.com/chibao13/news_feed_practice/services/apigateway/moudle/post/postbusiness"
	"github.com/chibao13/news_feed_practice/services/apigateway/moudle/post/postmodel"
	"github.com/chibao13/news_feed_practice/services/apigateway/moudle/post/poststorage"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"net/http"
)

func CreatePost(ctx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data postmodel.CreatePost
		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))
			return
		}
		requester := c.MustGet(common.CurrentUser).(common.Requester)

		data.UserId = requester.GetUserId()
		conn, err := grpc.Dial(common.HostgRPCPostService, grpc.WithInsecure(), grpc.WithBlock())
		defer conn.Close()
		if err != nil {
			panic(common.ErrInternal(err))
		}

		client := api.NewPostRPCClient(conn)
		storage := poststorage.NewGrpcStore(client)
		biz := postbusiness.NewCreatePostBusiness(storage, ctx.GetPubsub())

		if err := biz.CreatePost(c.Request.Context(), &data); err != nil {
			panic(err)
			return
		}
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data.Id))
	}
}

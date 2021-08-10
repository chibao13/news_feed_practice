package ginuserfriend

import (
	api "github.com/chibao13/news_feed_practice/api/v1/userfriend"
	"github.com/chibao13/news_feed_practice/common"
	"github.com/chibao13/news_feed_practice/component/appctx"
	"github.com/chibao13/news_feed_practice/services/apigateway/moudle/userfriend/userfriendbiz"
	"github.com/chibao13/news_feed_practice/services/apigateway/moudle/userfriend/userfriendstorage"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"net/http"
)

//func getUserfriendGRPCClient() api.UserFriendRPCClient {
//	conn, err := grpc.Dial(common.HostgRPCUserFriendService, grpc.WithInsecure(), grpc.WithBlock())
//	if err != nil {
//		log.Fatalf("did not connect: %v", err)
//	}
//
//	client := api.NewUserFriendRPCClient(conn)
//	return client
//}

func ListFriend(ctx appctx.AppContext) gin.HandlerFunc {
	//TODO reuse connection -- close connect
	//client := getUserfriendGRPCClient()

	return func(c *gin.Context) {
		userIdParam := c.Param("user-id")
		uid, err := common.FromBase58(userIdParam)
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		userId := uid.GetLocalID()
		conn, err := grpc.Dial(common.HostgRPCUserFriendService, grpc.WithInsecure(), grpc.WithBlock())
		defer conn.Close()
		if err != nil {
			panic(common.ErrInternal(err))
		}

		client := api.NewUserFriendRPCClient(conn)
		storage := userfriendstorage.NewGrpcStore(client)
		listFriendBiz := userfriendbiz.NewListFriendBusiness(storage)
		data, err := listFriendBiz.ListFriendIds(c.Request.Context(), map[string]interface{}{
			"user_id": userId,
		})
		if err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
	}
}

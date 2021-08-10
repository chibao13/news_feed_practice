package grpcpost

import (
	"context"
	api "github.com/chibao13/news_feed_practice/api/v1/post"
	"github.com/chibao13/news_feed_practice/component/appctx"
	"github.com/chibao13/news_feed_practice/services/postservice/moudle/post/postbusiness"
	"github.com/chibao13/news_feed_practice/services/postservice/moudle/post/postmodel"
	"github.com/chibao13/news_feed_practice/services/postservice/moudle/post/poststorage"
)

type grpcPostServer struct {
	api.UnimplementedPostRPCServer
	appCtx appctx.AppContext
}

func NewgRpcPostServer(appCtx appctx.AppContext) *grpcPostServer {
	return &grpcPostServer{
		appCtx: appCtx,
	}
}

func (g *grpcPostServer) CreatePost(ctx context.Context, request *api.CreatePostRequest) (*api.PostResponse, error) {

	dataCreate := postmodel.CreatePost{
		UserId:  int(request.UserId),
		Content: &request.Content,
	}

	db := g.appCtx.GetMainDBConnection()
	storage := poststorage.NewSQLStore(db)
	biz := postbusiness.NewCreatePostBusiness(storage, g.appCtx.GetPubsub())

	if err := biz.CreatePost(ctx, &dataCreate); err != nil {
		//panic(err)
		return nil, err
	}
	return &api.PostResponse{PostId: uint32(dataCreate.Id)}, nil
}

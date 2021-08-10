package grpcuserfriend

import (
	"context"
	api "github.com/chibao13/news_feed_practice/api/v1/userfriend"
	"github.com/chibao13/news_feed_practice/component/appctx"
	"github.com/chibao13/news_feed_practice/services/userfriendservice/moudle/userfriend/userfriendstorage"
)

type grpcUserServer struct {
	api.UnimplementedUserFriendRPCServer
	appCtx appctx.AppContext
}

func NewGrpcUser(appCtx appctx.AppContext) *grpcUserServer {
	return &grpcUserServer{appCtx: appCtx}
}

func (g *grpcUserServer) GetListFriends(ctx context.Context, request *api.ConditionRequest) (*api.ListFriendIdsResponse, error) {
	db := g.appCtx.GetMainDBConnection()
	storage := userfriendstorage.NewSqlStore(db)

	//TODO remove stupid conditions
	conditions := make(map[string]interface{}, len(request.Conditions))
	for key, value := range request.Conditions {
		conditions[key] = value
	}

	data, err := storage.ListFriendIdsWithCondition(ctx, conditions)
	if err != nil {
		return nil, err
	}
	return &api.ListFriendIdsResponse{Ids: data}, nil
}

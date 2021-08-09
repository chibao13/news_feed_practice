package grpcuser

import (
	"context"
	api "github.com/chibao13/news_feed_practice/api/v1"
	"github.com/chibao13/news_feed_practice/component/appctx"
	"github.com/chibao13/news_feed_practice/services/authenservice/moudle/user/userstorage"
)

type grpcUserServer struct {
	api.UnimplementedUserRPCServer
	appCtx appctx.AppContext
}

func NewGrpcUser(appCtx appctx.AppContext) *grpcUserServer {
	return &grpcUserServer{appCtx: appCtx}
}

func (g *grpcUserServer) FindUser(ctx context.Context, request *api.UserIdRequest) (*api.FindUserResponse, error) {

	db := g.appCtx.GetMainDBConnection()
	storage := userstorage.NewSQLStore(db)
	//TODO caching in this server
	conditions := make(map[string]interface{}, len(request.Conditions))

	for key, value := range request.Conditions {
		conditions[key] = value
	}
	data, err := storage.FindUser(ctx, conditions)
	if err != nil {
		return nil, err
	}

	return &api.FindUserResponse{
		Id:     uint32(data.Id),
		Status: int32(data.Status),
		Email:  data.Email,
	}, nil
}

func (g *grpcUserServer) GetListFriends(ctx context.Context, request *api.UserIdRequest) (*api.ListFriendIdsResponse, error) {
	//db := g.appCtx.GetMainDBConnection()
	//storage := userfriendstorage.NewSqlStore(db)
	//
	////TODO remove stupid conditions
	//conditions := make(map[string]interface{}, len(request.Conditions))
	//for key, value := range request.Conditions {
	//	conditions[key] = value
	//}
	//
	//data, err := storage.ListFriendIdsWithCondition(ctx, conditions)
	//if err != nil {
	//	return nil, err
	//}
	//return &api.ListFriendIdsResponse{Ids: data}, nil

	return nil, nil
}

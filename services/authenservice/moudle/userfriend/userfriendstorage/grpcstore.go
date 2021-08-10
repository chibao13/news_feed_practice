package userfriendstorage

import api "github.com/chibao13/news_feed_practice/api/v1/userfriend"

type grpcStore struct {
	client api.UserFriendRPCClient
}

func NewGrpcStore(client api.UserFriendRPCClient) *grpcStore {
	return &grpcStore{client: client}
}

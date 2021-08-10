package poststorage

import api "github.com/chibao13/news_feed_practice/api/v1/post"

type grpcStore struct {
	client api.PostRPCClient
}

func NewGrpcStore(client api.PostRPCClient) *grpcStore {
	return &grpcStore{client: client}
}

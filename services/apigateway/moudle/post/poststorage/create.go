package poststorage

import (
	"context"
	api "github.com/chibao13/news_feed_practice/api/v1/post"
	"github.com/chibao13/news_feed_practice/services/apigateway/moudle/post/postmodel"
	"log"
)

func (store *grpcStore) CreatePost(ctx context.Context, data *postmodel.CreatePost) error {

	response, err := store.client.CreatePost(ctx, &api.CreatePostRequest{
		UserId:  uint32(data.UserId),
		Content: *data.Content,
	})

	if err != nil {
		return err
	}

	log.Printf("%v ", response.PostId)
	return nil
}

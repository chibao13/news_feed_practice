package postbusiness

import (
	"context"
	"github.com/chibao13/news_feed_practice/common"
	"github.com/chibao13/news_feed_practice/pubsub"
	"github.com/chibao13/news_feed_practice/services/apigateway/moudle/post/postmodel"
)

type CreatePostStorage interface {
	CreatePost(ctx context.Context, data *postmodel.CreatePost) error
}

type createPostBusiness struct {
	postStorage CreatePostStorage
	pb          pubsub.Pubsub
}

func NewCreatePostBusiness(storage CreatePostStorage, pb pubsub.Pubsub) *createPostBusiness {
	return &createPostBusiness{postStorage: storage, pb: pb}
}

func (biz *createPostBusiness) CreatePost(ctx context.Context, data *postmodel.CreatePost) error {

	data.Status = 1
	if err := biz.postStorage.CreatePost(ctx, data); err != nil {
		return common.ErrCannotCreateEntity(postmodel.EntityName, err)
	}

	go func() {
		//Test error here
		defer common.AppRecover()
		biz.pb.Publish(ctx, common.TopicPostCreated, pubsub.NewMessage(data))
	}()
	return nil
}

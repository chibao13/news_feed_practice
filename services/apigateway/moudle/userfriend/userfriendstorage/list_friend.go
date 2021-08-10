package userfriendstorage

import (
	"context"
	api "github.com/chibao13/news_feed_practice/api/v1/userfriend"
	"github.com/chibao13/news_feed_practice/common"
	"log"
	"strconv"
)

func (store *grpcStore) ListFriendIdsWithCondition(ctx context.Context, conditions map[string]interface{}) ([]uint32, error) {
	listCondition := make(map[string]string, len(conditions))
	for key, value := range conditions {
		switch value.(type) {
		case string:
			listCondition[key] = value.(string)
		case uint32:
			listCondition[key] = strconv.Itoa(int(value.(uint32)))
		default:
			log.Printf("Not yet supported %t \n", value)
		}
	}
	response, err := store.client.GetListFriends(ctx, &api.ConditionRequest{
		Conditions: listCondition,
	})

	if err != nil {
		return nil, common.ErrInternal(err)
	}
	return response.Ids, nil
}

package userfriendstorage

import (
	"context"
	"github.com/chibao13/news_feed_practice/common"
	"github.com/chibao13/news_feed_practice/services/userfriendservice/moudle/userfriend/userfriendmodel"
)

func (store *sqlStore) ListFriendIdsWithCondition(ctx context.Context, cond map[string]interface{}) ([]uint32, error) {

	db := store.db.Table(userfriendmodel.UserFriend{}.TableName())
	db = db.Joins("inner join users on users.id = user_friends.friend_id and users.status <> 0")
	//db = db.Preload("User", "status <> 0")
	db = db.Where(cond)
	var ids []uint32
	if err := db.Select("friend_id").Find(&ids).Error; err != nil {
		return nil, common.ErrDB(err)
	}
	//data := make([]string, 0, len(ids))

	//for _, id := range ids {
	//	//data = append(data, common.NewUID(id.FriendId, common.DbTypeUser, 1).String())
	//	data = append(data, common.NewUID(id, common.DbTypeUser, 1).String())
	//}
	return ids, nil
}

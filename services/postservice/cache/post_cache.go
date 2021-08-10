package cache

import (
	"context"
	"fmt"
	"github.com/chibao13/news_feed_practice/common"
	"github.com/chibao13/news_feed_practice/memcache"
	"github.com/chibao13/news_feed_practice/services/postservice/moudle/post/postmodel"
	"log"
)

//var lock = sync.Mutex{}

//var singletonUserPostCaching *userPostCaching = nil

//Mono-singleton
//func GetPostCaching(db *gorm.DB) *userPostCaching {
//
//	if singletonUserPostCaching == nil {
//		lock.Lock()
//		defer lock.Unlock()
//		store := poststorage2.NewSQLStore(db)
//		singletonUserPostCaching = NewUserPostCaching(NewListCaching(30), store)
//	}
//	return singletonUserPostCaching
//}

func NewUserPostCaching(store memcache.ListCaching, postStore RealPostStore) *userPostCaching {
	return &userPostCaching{store: store, realStore: postStore}
}

type RealPostStore interface {
	ListPostWithCondition(ctx context.Context, condition map[string]interface{}, paging *common.Paging, moredata ...string) ([]postmodel.Post, error)
}

type userPostCaching struct {
	store     memcache.ListCaching
	realStore RealPostStore
}

func (c *userPostCaching) GetCachedStore() memcache.ListCaching {
	return c.store
}
func (u *userPostCaching) ListPostWithCondition(ctx context.Context,
	condition map[string]interface{}, paging *common.Paging, moredata ...string) ([]postmodel.Post, error) {
	var userId int = 0
	if value, ok := condition["user_id"]; ok {
		switch value.(type) {
		case int:
			userId = value.(int)
		}
	} else {
		//Test
		userId = 2
	}

	key := fmt.Sprintf("user:%d", userId)
	log.Printf("key %s", key)
	cachedList := u.store.RRange(key, 0, -1)
	if cachedList != nil {
		var listPost []postmodel.Post
		for i := len(cachedList) - 1; i >= 0; i-- {
			cachedPost := cachedList[i].(*postmodel.CreatePost)

			post := postmodel.Post{
				SQLModel: cachedPost.SQLModel,
				UserId:   userId,
				Content:  *cachedPost.Content,
			}
			post.Mask()
			listPost = append(listPost, post)
		}
		log.Println("Read from cached")
		return listPost, nil
	}
	log.Println("Read from store")
	return u.realStore.ListPostWithCondition(ctx, condition, paging, moredata...)
}

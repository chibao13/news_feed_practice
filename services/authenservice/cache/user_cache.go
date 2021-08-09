package cache

import (
	"context"
	"fmt"
	"github.com/chibao13/news_feed_practice/memcache"
	"github.com/chibao13/news_feed_practice/services/authenservice/moudle/user/usermodel"
	"time"
)

type RealStore interface {
	FindUser(ctx context.Context, conditions map[string]interface{}, moreInfo ...string) (*usermodel.User, error)
}
type userCaching struct {
	store     memcache.Caching
	realStore RealStore
}

func NewUserCaching(store memcache.Caching, realStore RealStore) *userCaching {
	return &userCaching{store: store, realStore: realStore}
}
func (u *userCaching) FindUser(ctx context.Context, conditions map[string]interface{}, moreInfo ...string) (*usermodel.User, error) {
	userId := conditions["id"].(int)
	key := fmt.Sprintf("users-%d", userId)
	userCached := u.store.Read(key)
	if userCached != nil {
		return userCached.(*usermodel.User), nil
	}
	user, err := u.realStore.FindUser(ctx, conditions, moreInfo...)
	if err != nil {
		return nil, err
	}
	u.store.WriteTTL(key, user, 5*time.Hour)
	return user, nil
}

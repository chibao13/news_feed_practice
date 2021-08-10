package poststorage

import (
	"context"
	"fmt"
	"github.com/btcsuite/btcutil/base58"
	"github.com/chibao13/news_feed_practice/common"
	"github.com/chibao13/news_feed_practice/services/postservice/moudle/post/postmodel"
	"time"
)

const timeLayout = "2006-01-02 15:04:05"

func (s *sqlStore) ListPostWithCondition(ctx context.Context, cond map[string]interface{}, paging *common.Paging, moreDatas ...string) ([]postmodel.Post, error) {
	db := s.db.Table(postmodel.Post{}.TableName())

	db = db.Where("status <> 0")
	db = db.Where(cond)
	db = db.Order("created_at desc")

	if err := db.Count(&paging.Total).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	var data []postmodel.Post
	db = db.Table(postmodel.Post{}.TableName()).Limit(paging.Limit)

	if v := paging.FakeCursor; v != "" {
		timeCreated, err := time.Parse(timeLayout, string(base58.Decode(v)))
		if err != nil {
			return nil, err
		}
		db = db.Where("created_at < ?", timeCreated.Format(timeLayout))
	} else {
		db = db.Offset((paging.Page - 1) * paging.Limit)
	}

	for i := range moreDatas {
		db = db.Preload(moreDatas[i])
	}

	if err := db.Find(&data).Error; err != nil {
		return nil, err
	}
	for i := range data {
		data[i].Mask()
		if i == len(data)-1 {
			paging.NextCursor = base58.Encode([]byte(fmt.Sprintf("%v", data[i].CreatedAt.Format(timeLayout))))
		}
	}
	return data, nil
}

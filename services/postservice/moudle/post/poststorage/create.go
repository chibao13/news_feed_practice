package poststorage

import (
	"context"
	"github.com/chibao13/news_feed_practice/common"
	"github.com/chibao13/news_feed_practice/services/postservice/moudle/post/postmodel"
)

func (s *sqlStore) CreatePost(ctx context.Context, data *postmodel.CreatePost) error {
	db := s.db.Begin()

	if err := db.Table(data.TableName()).Create(&data).Error; err != nil {
		db.Rollback()
		return common.ErrDB(err)
	}
	if err := db.Commit().Error; err != nil {
		db.Rollback()
		return common.ErrDB(err)
	}

	return nil
}

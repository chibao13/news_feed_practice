package userstorage

import (
	"context"
	"github.com/chibao13/news_feed_practice/common"
	"github.com/chibao13/news_feed_practice/services/apigateway/moudle/user/usermodel"
)

func (s *sqlStore) CreateUser(ctx context.Context, data *usermodel.UserCreate) error {
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

package postmodel

import "github.com/chibao13/news_feed_practice/common"

type CreatePost struct {
	common.SQLModel `json:",inline"`
	UserId          int     `json:"-" gorm:"column:user_id;"`
	Content         *string `json:"content" form:"content" gorm:"column:content;" form:"content"`
}

func (CreatePost) TableName() string {
	return "posts"
}

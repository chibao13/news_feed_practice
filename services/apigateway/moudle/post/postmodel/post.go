package postmodel

import "github.com/chibao13/news_feed_practice/common"

const EntityName = "Post"

type Post struct {
	common.SQLModel `json:",inline"`
	UserId          int         `json:"-" gorm:"column:user_id;"`
	FakeUserId      *common.UID `json:"user_id" gorm:"-"`
	Content         string      `json:"content" gorm:"column:content"`

	User *common.SimpleUser `json:"user" gorm:"foreignKey:UserId;"`
}

func (Post) TableName() string {
	return "posts"
}

func (p *Post) Mask() {
	p.GenUID(common.DbTypeNote)

	uid := common.NewUID(uint32(p.UserId), common.DbTypeUser, 1)
	p.FakeUserId = &uid

	if p.User != nil {
		p.User.GenUID(common.DbTypeUser)
	}
}

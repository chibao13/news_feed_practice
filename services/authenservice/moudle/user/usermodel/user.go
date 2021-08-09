package usermodel

import "github.com/chibao13/news_feed_practice/common"

const EntityName = "User"

type User struct {
	common.SQLModel `json:",inline"`
	Email           string `json:"email" gorm:"column:email;"`
	Password        string `json:"-" gorm:"column:password;"`
	Salt            string `json:"-" gorm:"column:salt;"`
}

func (u *User) GetUserId() int {
	return u.Id
}
func (u *User) GetEmail() string {
	return u.Email
}

func (User) TableName() string {
	return "users"
}

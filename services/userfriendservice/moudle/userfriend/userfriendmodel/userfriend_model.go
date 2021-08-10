package userfriendmodel

import (
	"time"
)

const EntityName = "UserFriend"

type UserFriend struct {
	UserId    uint32     `json:"-" gorm:"column:user_id;"`
	FriendId  uint32     `json:"-" gorm:"column:friend_id;"`
	CreatedAt *time.Time `json:"created_at" gorm:"created_at"`
	//User *common.SimpleUser `json:"friend" gorm:"foreignKey:FriendId;"`
}

func (UserFriend) TableName() string {
	return "user_friends"
}

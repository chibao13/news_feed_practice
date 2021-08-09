package appctx

import (
	"github.com/chibao13/news_feed_practice/pubsub"
	"gorm.io/gorm"
)

type AppContext interface {
	GetMainDBConnection() *gorm.DB
	SecretKey() string
	GetPubsub() pubsub.Pubsub
}

type ctx struct {
	mainDB *gorm.DB
	secret string
	ps     pubsub.Pubsub
}

func New(mainDB *gorm.DB, secret string, pubsub pubsub.Pubsub) *ctx {
	return &ctx{mainDB: mainDB, secret: secret, ps: pubsub}
}

func (c ctx) GetMainDBConnection() *gorm.DB {
	return c.mainDB
}

func (c ctx) SecretKey() string {
	return c.secret
}
func (c ctx) GetPubsub() pubsub.Pubsub {
	return c.ps
}

type tokenExpiry struct {
	atExp int
	rtExp int
}

func NewTokenConfig() tokenExpiry {
	return tokenExpiry{
		atExp: 60 * 60 * 24 * 7,
		rtExp: 60 * 60 * 24 * 7 * 2,
	}
}

func (tk tokenExpiry) GetAtExp() int {
	return tk.atExp
}

func (tk tokenExpiry) GetRtExp() int {
	return tk.rtExp
}

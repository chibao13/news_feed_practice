package memcache

import (
	"github.com/chibao13/news_feed_practice/common"
	"sync"
	"time"
)

type caching struct {
	store  map[string]interface{}
	locker *sync.RWMutex
	timer  *time.Timer
}
type Caching interface {
	Write(key string, value interface{})
	Read(key string) interface{}
	WriteTTL(key string, value interface{}, exp time.Duration)
}

func NewCaching() *caching {
	return &caching{
		store:  make(map[string]interface{}),
		locker: new(sync.RWMutex),
		//quit:   make(chan bool),
	}
}

func (c *caching) Read(key string) interface{} {
	c.locker.RLock()
	defer c.locker.RUnlock()
	if value, ok := c.store[key]; ok {
		return value
	}

	return nil
}
func (c *caching) Write(key string, value interface{}) {
	c.locker.Lock()
	defer c.locker.Unlock()
	c.store[key] = value
}

func (c *caching) WriteTTL(key string, value interface{}, exp time.Duration) {
	c.locker.Lock()
	defer c.locker.Unlock()
	c.store[key] = value

	if c.timer != nil {
		c.timer.Stop()
	}
	go func() {
		defer common.AppRecover()
		c.timer = time.NewTimer(exp)
		select {
		case <-c.timer.C:
			c.Write(key, nil)
			c.timer = nil
		}
	}()
}

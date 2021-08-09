package memcache

import (
	"github.com/chibao13/news_feed_practice/common"
	"sync"
	"time"
)

type ListCaching interface {
	RPush(key string, value interface{})
	RRange(key string, start, end int) []interface{}
	RPop(key string) interface{}
	RPushTTL(key string, value interface{}, exp time.Duration)
}

type listCaching struct {
	store  map[string][]interface{}
	locker *sync.RWMutex
	maxLen int
	//timer  *time.Timer
}

func NewListCaching(maxLen int) *listCaching {
	return &listCaching{
		store:  make(map[string][]interface{}),
		locker: new(sync.RWMutex),
		maxLen: maxLen,
	}
}
func (c *listCaching) RPush(key string, value interface{}) {
	c.locker.Lock()
	defer c.locker.Unlock()

	newSlice := append(c.store[key], value)
	if len(newSlice) >= c.maxLen {
		_, newSlice = newSlice[0], newSlice[1:]
	}
	c.store[key] = newSlice
}

func (c *listCaching) RRange(key string, start, end int) []interface{} {
	c.locker.RLock()
	defer c.locker.RUnlock()

	value, ok := c.store[key]
	if !ok {
		return nil
	}
	if end < 0 || end >= len(value) {
		return value
	}
	return value[start:end]
}

func (c *listCaching) RPushTTL(key string, value interface{}, exp time.Duration) {
	c.RPush(key, value)
	go func() {
		common.AppRecover()
		timer := time.NewTimer(exp)
		<-timer.C
		c.locker.Lock()
		defer c.locker.Unlock()
		if slice, ok := c.store[key]; ok {
			c.store[key] = slice[1:]
		}
	}()
}

func (c *listCaching) RPop(key string) interface{} {
	c.locker.Lock()
	defer c.locker.Unlock()

	if slice, ok := c.store[key]; ok {
		if len(slice) >= 1 {
			rValue := slice[len(slice)-1]
			c.store[key] = slice[:len(slice)-1]
			return rValue
		}
	}
	return nil
}

package lru

import (
	"github.com/ccb1900/tinyweb/cache/contract"
	"github.com/ccb1900/tinyweb/log"
	lru "github.com/hashicorp/golang-lru"
	"time"
)

type LruCache struct {
	Cache *lru.ARCCache
}

func (l *LruCache) Get(key string) interface{} {
	if v,ok :=l.Cache.Get(key);ok {
		return v
	} else {
		return nil
	}
}

func (l *LruCache) Set(key string, value interface{}) {
	l.Cache.Add(key,value)
}

func (l *LruCache) SetWithExpire(key string, value interface{}, expire time.Time) {
	l.Cache.Add(key,value)
}

func (l *LruCache) Delete(key string) error {
	l.Cache.Remove(key)
	return nil
}

func New() contract.ICache  {
	c,err := lru.NewARC(1024)
	if err != nil {
		log.Fatal(err)
	}
	return &LruCache{
		Cache: c,
	}
}

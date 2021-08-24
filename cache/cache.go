package cache

import (
	"github.com/ccb1900/tinyweb/cache/contract"
	"github.com/ccb1900/tinyweb/cache/lru"
	"github.com/ccb1900/tinyweb/cache/redis"
	"github.com/ccb1900/tinyweb/config"
)
var c contract.ICache

func Init()  {
	switch config.Get("cache.driver") {
	case "redis":
		c = redis.New()
		break
	case "lru":
		c = lru.New()
		break
	}
}

func Get(key string) interface{}  {
	return c.Get(key)
}

func Set(key string,value interface{})  {
	c.Set(key,value)
}

func Delete(key string) error  {
	return c.Delete(key)
}





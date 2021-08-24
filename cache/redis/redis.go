package redis

import (
	"github.com/ccb1900/tinyweb/cache/contract"
	"github.com/ccb1900/tinyweb/config"
	_ "github.com/go-redis/cache/v8"
	rv8 "github.com/go-redis/redis/v8"
	"time"
)

func client() *rv8.Client   {
	rdb := rv8.NewClient(&rv8.Options{
		Addr:     config.Get("cache.redis.host")+":"+config.Get("cache.redis.port"),
		Password: config.Get("cache.redis.password"), // no password set
		DB:       config.GetInt("cache.redis.db"),  // use default DB
	})

	return rdb
}

type RedisCache struct {
	Cache *rv8.Client
}
func (r *RedisCache) Get(key string) interface{} {
	panic("implement me")
}

func (r *RedisCache) Set(key string, value interface{}) {
	panic("implement me")
}

func (l *RedisCache) SetWithExpire(key string, value interface{}, expire time.Time) {
	panic("implement me")
}

func (l *RedisCache) Delete(key string) error {
	panic("implement me")
}

func New() contract.ICache  {
	return &RedisCache{Cache: client()}
}


package redis

import (
	"fmt"
	"github.com/ccb1900/tinyweb/config"
	"github.com/ccb1900/tinyweb/log"
	"github.com/go-redis/redis/v8"
)
var rdb *redis.Client
func client()   {
	rdb = redis.NewClient(&redis.Options{
		Addr:     config.Get("redis.host")+":"+config.Get("redis.port"),
		Password: config.Get("redis.password"), // no password set
		DB:       config.GetInt("redis.db"),  // use default DB
	})
}
func client2()  {
	opt, err := redis.ParseURL(fmt.Sprintf("redis://%s:%s@%s:%s/%s",
		config.GetInt("redis.host"),
		config.GetInt("redis.port"),
		config.GetInt("redis.password"),
		config.GetInt("redis.db"),
		config.GetInt("redis.db"),
	))
	if err != nil {
		log.Debug(err)
		panic(err)
	}

	rdb = redis.NewClient(opt)
}
func Init()  {
	client()
}

func GetRDB() *redis.Client  {
	return rdb
}

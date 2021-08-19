package redis

import (
	"context"
	"fmt"
	"github.com/ccb1900/tinyweb/config"
	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()
func client()  {
	rdb := redis.NewClient(&redis.Options{
		Addr:     config.Get("redis.host")+":"+config.Get("redis.port"),
		Password: config.Get("redis.password"), // no password set
		DB:       config.GetInt("redis.db"),  // use default DB
	})

	err := rdb.Set(ctx, "key", "value", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := rdb.Get(ctx, "key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)

	val2, err := rdb.Get(ctx, "key2").Result()
	if err == redis.Nil {
		fmt.Println("key2 does not exist")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("key2", val2)
	}
	// Output: key value
	// key2 does not exist
}

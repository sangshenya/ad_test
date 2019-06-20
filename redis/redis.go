package redis

import (
	"github.com/go-redis/redis"
	"ad_test/config"
)

var(
	cluster *redis.Client
)

func init()  {
	// 初始化
	c := redis.NewClient(&redis.Options{
		Addr:config.REDISURL,
		Password:config.REDISPASSWORD,
		DB:0,
	})

	cluster = c
}
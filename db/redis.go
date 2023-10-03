package db

import (
	"context"
	"github.com/redis/go-redis/v9"
	"log"
)

/*
redis服务器
*/

var Rdb *redis.Client

func init() {
	client := redis.NewClient(&redis.Options{
		Addr:     "192.168.0.103:6379",
		Password: "admin123", // no password set
		DB:       0,          // uses default DB
		PoolSize: 1000,
	})
	ctx := context.Background()
	ping := client.Ping(ctx)
	if ping.String() == "ping: PONG" {
		log.Println("连接redis 成功!")
	}
	Rdb = client
}

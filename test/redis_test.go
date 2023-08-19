package test

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"log"
	"testing"
)

func TestRedis(t *testing.T) {
	client := redis.NewClient(&redis.Options{
		Addr:     "192.168.0.107:6379",
		Password: "admin123", // no password set
		DB:       0,          // uses default DB
		PoolSize: 1000,
	})
	ctx := context.Background()
	ping := client.Ping(ctx)
	fmt.Println(ping.String())
	if ping.String() == "ping: PONG" {
		log.Println("连接redis 成功!")
	}
}

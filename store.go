package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"time"

	redis "github.com/redis/go-redis/v9"
)

func redis_con() *redis.Client {
	address := fmt.Sprint(os.Getenv("REDIS_HOST"))
	client := redis.NewClient(&redis.Options{
		Addr:     net.JoinHostPort(address, os.Getenv("REDIS_PORT")),
		Password: os.Getenv("REDIS_PASSWORD"),
	})

	// context and ping to verify redis connection
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err := client.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}

	return client
}

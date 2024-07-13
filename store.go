package main

import (
	"fmt"
	"net"
	"os"

	redis "github.com/redis/go-redis/v9"
)

func redis_con() *redis.Client {
	address := fmt.Sprint(os.Getenv("REDIS_HOST"))
	client := redis.NewClient(&redis.Options{
		Addr:     net.JoinHostPort(address, os.Getenv("REDIS_PORT")),
		Password: os.Getenv("REDIS_PASSWORD"),
	})

	return client
}
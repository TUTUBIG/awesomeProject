package main

import (
	"context"
	"github.com/go-redis/redis/v8"
	"log"
	"time"
)

func main() {
	c := redis.NewClient(&redis.Options{
		Addr: "192.168.1.91:6379",
	})

	if e := c.Set(context.Background(), "test", "alvin", time.Minute).Err(); e != nil {
		log.Fatal(e)
	}
	r, e := c.Get(context.Background(), "test").Result()
	if e != nil {
		log.Fatal(e)
	}
	log.Println(r)
}

package main

import (
	"golang.org/x/time/rate"
	"time"
)

func main() {
	limit := rate.Every(time.Second)
	limiter := rate.NewLimiter(limit, 10)
	limiter.Burst()
}

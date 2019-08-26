package redis

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)
import "git.ezbuy.me/ezbuy/base/storage/redis"

var cfg *redis.Config

type GenerateReturnCode struct {
	totalCount int
	storeNumbers []int
	distributeLock *DistributeLock
	locker *sync.Mutex
}

type DistributeLock struct {
	redisClient *redis.Client
	expiration time.Duration
}

func (g *GenerateReturnCode) Init(totalCount int) {
	g.storeNumbers = make([]int,totalCount)
	for i := range g.storeNumbers {
		g.storeNumbers[i] = i
	}
	g.locker = new(sync.Mutex)
	g.distributeLock = &DistributeLock{
		redisClient:redis.New(cfg),
		expiration:(14 * 24 + 1) * time.Hour,
	}
}

func (g *GenerateReturnCode) getRandom() int {
	if len(g.storeNumbers) <= 0 {
		return -1
	}

	randomIndex := rand.Intn(len(g.storeNumbers))
	randomNumber := g.storeNumbers[randomIndex]

	g.locker.Lock()
	defer g.locker.Unlock()

	g.storeNumbers = append(g.storeNumbers[:randomIndex],g.storeNumbers[randomIndex+1:]...)

	if err := g.distributeLock.lock(fmt.Sprint(randomNumber));err != nil {
		return -1
	}
	return randomNumber
}

func (g *GenerateReturnCode) put(number int)  {

	g.locker.Lock()
	defer g.locker.Unlock()

	if err := g.distributeLock.unlock(fmt.Sprint(number)); err != nil {
		g.storeNumbers = append(g.storeNumbers,number)
	}

	return
}

func (l *DistributeLock) lock(key string) error {
	return l.redisClient.SetNX(key,true,l.expiration).Err()
}

func (l *DistributeLock) unlock(key string) error {
	return l.redisClient.Del(key).Err()
}
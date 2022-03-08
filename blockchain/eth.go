package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/prometheus/common/log"
	"math/big"
	"sync"
	"time"
)

const urlWs = "wss://apis-sj.ankr.com/wss/9561f7e4c4a74e3eb6d2ee8bee84764e/7af9c7a27032e18bc95fda33ab009406/binance/full/main"

func main() {
	websocket()
}

func websocket() {
	client, err := ethclient.Dial(urlWs)
	if err != nil {
		panic(err)
	}

	client.BalanceAt()

	headerChannel := make(chan *types.Header, 1024)
	subNewHeader, err := client.SubscribeNewHead(context.Background(), headerChannel)
	if err != nil {
		panic(err)
	}

	logChannel := make(chan types.Log, 1024)
	lf := ethereum.FilterQuery{
		Addresses: []common.Address{common.HexToAddress("0x000000000000000000000000bb4cdb9cbd36b01bd1cbaebf2de08d9173bc095c")},
		Topics:    [][]common.Hash{{common.HexToHash("0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c")}},
	}
	subLogs, err := client.SubscribeFilterLogs(context.Background(), lf, logChannel)
	if err != nil {
		panic(err)
	}

	fmt.Println("---subscribe-----")

	go func() {
		time.Sleep(5 * time.Second)
		fmt.Println("---unsubscribe-----")
		subNewHeader.Unsubscribe()
		subLogs.Unsubscribe()
	}()

	go func() {
		for c := range headerChannel {
			fmt.Println("new header: ", c.Number)
		}
	}()

	go func() {
		for c := range logChannel {
			fmt.Println("new log: ", c.BlockNumber, c.Address.Hash())
		}
	}()

	wg := new(sync.WaitGroup)

	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("subLogs end", <-subLogs.Err())
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("subNewHeader end", <-subNewHeader.Err())
	}()

	wg.Wait()
}

func rpc() {
	client, err := ethclient.Dial("http://218.30.91.134:21145")
	if err != nil {
		panic(err)
	}

	number := big.NewInt(10865975)
	block, err := client.BlockByNumber(context.Background(), number)
	if err != nil {
		panic(err)
	}

	log.Info(block.Number().Uint64())
	log.Info(block.Hash().String())
}

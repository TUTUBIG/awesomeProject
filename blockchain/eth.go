package main

import (
	"context"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/prometheus/common/log"
	"math/big"
)

func main() {
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

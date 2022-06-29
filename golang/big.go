package main

import (
	"fmt"
	"math"
	"math/big"
)

var eth = big.NewFloat(math.Pow10(18))

func main() {
	hex := "0x51b9b6215e5e1ec72a0000"
	f, _ := new(big.Float).SetString(hex)
	fmt.Println(f.String())
	f = new(big.Float).Quo(f, eth)
	f = f.Mul(f, big.NewFloat(1000))
	fmt.Println(f.Text('f', 8), f.String())
	fmt.Println(f.Float64())

	fmt.Printf("%.2f", 0.9/3)
}

package main

import (
	"fmt"
	"math/rand"
	"time"
)

var r = rand.New(rand.NewSource(time.Now().Unix()))

func main() {

	fmt.Printf("0x%0x", 11)

	a := []string{"1", "w", "4"}
	b := a[1:]
	fmt.Println()

	for {
		fmt.Println(r.Intn(10))
		time.Sleep(time.Second)
	}
}

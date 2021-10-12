package main

import (
	"fmt"
	"runtime"
	"time"
)

type Test struct {
	A int
}

func main() {
	tester := Test{A: 10}

	for i := 0; i < 10; i++ {
		go func(j int) {
			time.Sleep(time.Second)
			fmt.Println(j)
		}(i)
	}

	runtime.SetFinalizer(&tester, func() {
		fmt.Println("get back tester")
	})

	runtime.KeepAlive(&tester)

	time.Sleep(time.Minute)
}

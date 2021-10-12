package main

import (
	"runtime"
	"time"
)

func TT() {
	i := 1
	i = i + 1
}

func main() {
	runtime.GOMAXPROCS(1)
	go func() {
		for {
			TT()
		}
	}()

	time.Sleep(time.Second)

}

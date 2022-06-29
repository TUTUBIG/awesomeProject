package main

import (
	"fmt"
	"time"
)

type A struct {
	A1 string
	B1 int
}

func T(m map[string]string) {
	fmt.Println(m["a"])
}

func main() {
	m := make(map[A]int)

	a1 := A{
		A1: "hallo",
		B1: 1,
	}

	a2 := A{
		A1: "hh",
		B1: 2,
	}

	m[a1] = 1
	m[a2] = 2

	a3 := A{
		A1: "hallo",
		B1: 2,
	}

	T(nil)
	m1 := make(map[string]string, 0)
	m1["a"] = "a"
	m1["b"] = "b"
	T(m1)

	fmt.Println(m[a1], m[a3])

	/*m2 := make(map[string]map[string]string, 0)
	m2["hello"]["haha"] = "1"

	fmt.Println(m2["hello"]["haha"])*/

	sm := make(map[string]string, 1)
	sm["1"] = "1"

	for i := 0; i < 50; i++ {
		go func() {
			for {
				fmt.Println(sm["1"])
			}
		}()
	}

	time.Sleep(time.Minute)
}

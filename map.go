package main

import "fmt"

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
		A1:"hallo",
		B1:1,
	}

	a2 := A{
		A1:"hh",
		B1:2,
	}

	m[a1] = 1
	m[a2] = 2

	a3 := A{
		A1:"hallo",
		B1:2,
	}

	T(nil)
	m1 := make(map[string]string,0)
	m1["a"]="a"
	m1["b"]="b"
	T(m1)

	fmt.Println(m[a1],m[a3])
}

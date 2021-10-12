package main

import "fmt"

type A1 struct {
	A string
	B string
}

type A struct {
	a *A1
	b string
}

func main() {
	a := &A{
		a: &A1{
			A: "A",
			B: "B",
		},
		b: "b",
	}

	fmt.Printf("%v", a.a)
}

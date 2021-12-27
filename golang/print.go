package main

import "fmt"

type A1 struct {
	A string
	B string
}

type AA struct {
	a *A1
	b string
}

func main() {
	a := &AA{
		a: &A1{
			A: "A",
			B: "B",
		},
		b: "b",
	}

	fmt.Printf("%v", a.a)
}

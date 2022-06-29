package main

import (
	"fmt"
	"reflect"
)

type P struct {
	A string
}

func main() {
	p := new(P)
	fmt.Println(p.A)

	p1 := new(P)
	p1.A = "hallo"
	p2 := new(P)
	p2 = p1
	p2.A = "haha"
	fmt.Println(p1.A, p2.A)
	p3 := new(P)
	*p3 = *p2
	p3.A = "ahaha"
	fmt.Println(p1, p2, p3)

	a := []int{1, 2, 3, 4, 5, 6, 7}
	b := make([]int, 4)

	fmt.Println(copy(b, a[:3]))

	fmt.Println(b, len(b), cap(b))

	var p10 *P
	test(p10)
	p11 := new(P)
	test(p11)
	test(nil)
}

func test(i interface{}) {
	if i == nil {
		fmt.Println("nil nil")
		return
	}
	if reflect.ValueOf(i).Kind() == reflect.Ptr {
		fmt.Println("ptr")
	}

	if reflect.ValueOf(i).IsNil() {
		fmt.Println("nil")
	}
}

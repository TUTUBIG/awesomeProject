package main

import (
	"fmt"
	"reflect"
)

type TI struct {
	A string
}

func t(i interface{}) {
	fmt.Println(reflect.ValueOf(i).Kind(), reflect.ValueOf(i).IsNil())
	ti := &TI{A: "a"}
	fmt.Println(reflect.TypeOf(i), reflect.ValueOf(i).Elem().CanSet())
	reflect.ValueOf(i).Elem().Set(reflect.ValueOf(ti))
}
func t1(i interface{}) {
	fmt.Println(reflect.TypeOf(i))
	reflect.ValueOf(i).Elem().Set(reflect.Zero(reflect.TypeOf(i).Elem()))
}

func main() {
	ti := new(TI)
	t(&ti)
	fmt.Println(ti)
	t1(&ti)
	fmt.Println(ti)

	tt()
}

func tt() {
	k := TI{}
	var i interface{}
	i = k
	ii, ok := i.(TI)
	if !ok {
		fmt.Println("oooo")
		return
	}
	ii.A = "bbbbb"
	fmt.Println("------", k)

}

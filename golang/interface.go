package main

//https://m.jb51.cc/go/795663.html

import (
	"fmt"
	"reflect"
)

type IT struct {
	A string
}

func TF(i interface{}) {
	fmt.Println("can set: ", reflect.ValueOf(i).Elem().CanSet())
	reflect.ValueOf(i).Elem().Set(reflect.Zero(reflect.ValueOf(i).Elem().Type()))
}

func TF1(i interface{}) {
	fmt.Println("TF1", i, reflect.ValueOf(i).IsNil())
	if reflect.ValueOf(i).IsNil() {
		tn := reflect.New(reflect.TypeOf(i).Elem())
	}
}

func main() {
	it := new(IT)
	fmt.Println("before", reflect.TypeOf(it), it)
	TF(&it)
	fmt.Println("after", reflect.TypeOf(it), it)

	var it1 *IT
	fmt.Println("it1 before", reflect.TypeOf(it1), it1, &it1)

	TF1(it1)
	fmt.Println("it1", reflect.TypeOf(it1), it1)

}

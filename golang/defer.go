package main

import "fmt"

func test1() (x int) {
	defer fmt.Printf("in defer: x = %d\n", x)

	x = 7
	return 9
}

func test2() (x int) {
	defer func() {
		fmt.Printf("in defer: x = %d\n", x)
	}()

	x = 7
	return 9
}

func test3() (x int) {
	x = 7
	defer fmt.Printf("in defer: x = %d\n", x)
	return 9
}

func test4() (x int) {
	defer func(n int) {
		fmt.Printf("in defer x as parameter: x = %d\n", n)
		fmt.Printf("in defer x after return: x = %d\n", x)
	}(x)

	x = 7
	return 9
}

func (a A) Func() {
	fmt.Println("a")
}

type B struct {
	A
}

func (b B) Func() {
	fmt.Println("b")
}

func main() {
	new(A).Func()
	new(B).Func()

	fmt.Println("test1")
	fmt.Printf("in main: x = %d\n", test1())
	fmt.Println("test2")
	fmt.Printf("in main: x = %d\n", test2())
	fmt.Println("test3")
	fmt.Printf("in main: x = %d\n", test3())
	fmt.Println("test4")
	fmt.Printf("in main: x = %d\n", test4())

	Aa()
}

func Aa() {
	var i *int
	return
	defer fmt.Println(i)
}

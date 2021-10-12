package main

import (
	"fmt"
	"strings"
)

func FS(s []int) {
	s = append(s, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20)
	fmt.Println("前 ", len(s), cap(s), s)
}

func FA(a [10]int) {
	a[1] = 10000
	fmt.Println("前 ", len(a), cap(a), a)
}

func FC(c []string) {
	fmt.Println("fc ", c, len(c), strings.Join(c, ","))
}

func main() {
	s := make([]int, 5, 10)
	fmt.Println(len(s), cap(s), s)
	s = append(s, 1)
	fmt.Println(len(s), cap(s), s)
	s = append(s, 2)
	fmt.Println(len(s), cap(s), s)
	s = append(s, 3, 4, 5, 6, 7, 8)
	fmt.Println(len(s), cap(s), s)
	FS(s)
	fmt.Println("后", len(s), cap(s), s)
	a := [10]int{1, 2, 3, 4, 5}
	FA(a)
	fmt.Println("后", len(a), cap(a), a)

	fmt.Println("------------")

	b := make([]int, 2)
	fmt.Println("b before ", b, len(b))
	b = append(b, 3)
	fmt.Println("b after ", b, len(b))

	b1 := make([]int, 0, 2)
	fmt.Println("b1 before ", b1, len(b1))
	b1 = append(b1, 3, 4, 5, 6)
	fmt.Println("b1 after", b1, len(b1))

	fmt.Println("------------")

	FC(nil)

	c := []int{1, 2, 3, 4, 5}
	fmt.Println("----c----", c)
	c = append(c[:2], c[3:]...)
	fmt.Println("-----c---", c)
	c = append(c[:2], c[3:]...)
	fmt.Println("-----c---", c)
	c = append(c, 6)
	fmt.Println("-----c---", c)
	c = append(c[:2], c[3:]...)
	fmt.Println("-----c---", c)

	var d []int

	d = append(d, 1)
	fmt.Println(d)

}

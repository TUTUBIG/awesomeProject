package main

import (
	"fmt"
	"unsafe"
)

const d  = "hello,alvin,nih,hu,ka"


type AS struct {
	A string
	B int
	c int
}

func main() {
	var a  = "hello,alvin,nih,hu,ka,fsdsfafas"
	var b int16 = 12
	var c  = 1
	var e  = false
	var f AS

	fs := make([]AS,3,3)
	fs3 := [3]AS{}

	fmt.Println(unsafe.Offsetof(f.B),unsafe.Offsetof(f.A),unsafe.Alignof(f.B))

	fmt.Println(unsafe.Sizeof(d))

	fmt.Println(unsafe.Sizeof(a),unsafe.Sizeof(b),unsafe.Sizeof(c),unsafe.Sizeof(e),unsafe.Sizeof(f),unsafe.Sizeof(fs),unsafe.Sizeof(fs3))


	fs = append(fs,fs3[:]...)

	var fs4 [4]AS


	fmt.Println(unsafe.Sizeof(fs),unsafe.Sizeof(fs4))
}

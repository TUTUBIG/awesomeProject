package main

import (
	"fmt"
	"strings"
)

func main() {
	a := strings.Split("",",")
	fmt.Println(len(a),a[0])

	fmt.Println(strings.Join([]string{},";"))
}

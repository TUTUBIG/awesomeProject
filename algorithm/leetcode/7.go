package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(reverse(1563847412))
}

func reverse(x int32) int32 {
	var y int32
	for x != 0 {
		p := x % 10
		x = x / 10
		if y > math.MaxInt32/10 || (y == math.MaxInt32/10 && p > 7) {
			return 0
		}
		if y < math.MinInt32/10 || (y == math.MinInt32/10 && p < -8) {
			return 0
		}
		y = 10*y + p
	}
	return y
}

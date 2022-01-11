package main

import "fmt"

func main() {
	t := candy([]int{1, 3, 2, 5, 2})
	fmt.Println(t)
}

func candy(ratings []int) int {
	candies := make([]int, len(ratings))
	candies[0] = 1
	for i := 1; i < len(ratings); i++ {
		if ratings[i] > ratings[i-1] {
			candies[i] = candies[i-1] + 1
		} else {
			candies[i] = 1
		}
	}

	fmt.Println(candies)

	for i := len(candies) - 2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] {
			if candies[i+1]+1 > candies[i] {
				candies[i] = candies[i+1] + 1
			}
		}
	}

	fmt.Println(candies)

	var total int
	for i := range candies {
		total += candies[i]
	}
	return total
}

package algorithm

import "fmt"

func rotate(nums []int, k int) {
	if k == 0 {
		return
	}
	if k > len(nums) {
		k = k % len(nums)
		if k == 0 {
			return
		}
	}
	fmt.Println(k)
	temp := make([]int, k)
	var j = 0
	for i := len(nums) - k; i < len(nums); i++ {
		temp[j] = nums[i]
		j++
	}

	for i := len(nums) - k - 1; i >= 0; i-- {
		nums[i+k] = nums[i]
	}

	for i := range temp {
		nums[i] = temp[i]
	}
	return
}

func main() {
	rotate([]int{1, 2}, 3)
}

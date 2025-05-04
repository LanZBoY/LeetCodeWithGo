package main

import (
	"fmt"
	"math"
)

func maxSlidingWindow(nums []int, k int) []int {
	N := len(nums)
	res := make([]int, N-k+1)

	for i := range len(res) {
		slic_win := nums[i : i+k]
		maxNum := math.MinInt32
		for _, num := range slic_win {
			maxNum = max(num, maxNum)
		}
		res[i] = maxNum
	}

	return res
}

func main() {
	fmt.Printf("%v\n", maxSlidingWindow([]int{1, -1}, 1))
	fmt.Printf("%v\n", maxSlidingWindow([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3))
	fmt.Printf("%v\n", maxSlidingWindow([]int{1}, 1))
}

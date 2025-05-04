package main

import (
	"fmt"
)

func maxSlidingWindow(nums []int, k int) []int {
	deQueue := make([]int, 0)
	for i := range k {
		for len(deQueue) > 0 && deQueue[len(deQueue)-1] < nums[i] {
			deQueue = deQueue[:len(deQueue)-1]
		}
		deQueue = append(deQueue, nums[i])
	}

	res := make([]int, len(nums)-k+1)

	res[0] = deQueue[0]

	for i := k; i < len(nums); i++ {
		removeNum := nums[i-k]
		if removeNum == deQueue[0] {
			deQueue = deQueue[1:]
		}
		for len(deQueue) > 0 && deQueue[len(deQueue)-1] < nums[i] {
			deQueue = deQueue[:len(deQueue)-1]
		}
		deQueue = append(deQueue, nums[i])
		res[i-k+1] = deQueue[0]
	}
	return res
}

func main() {
	fmt.Printf("%v\n", maxSlidingWindow([]int{1, 3, 1, 2, 0, 5}, 3))
	fmt.Printf("%v\n", maxSlidingWindow([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3))
	fmt.Printf("%v\n", maxSlidingWindow([]int{1, -1}, 1))
	fmt.Printf("%v\n", maxSlidingWindow([]int{1}, 1))
}

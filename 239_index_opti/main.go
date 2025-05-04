package main

import (
	"fmt"
)

/*
	會多出index opti版本是因為我覺得index queue更能精確判斷sliding windows滑動的狀況
*/

func maxSlidingWindow(nums []int, k int) []int {
	index_dequeue := make([]int, 0)
	for i := range k {
		for len(index_dequeue) > 0 && nums[index_dequeue[len(index_dequeue)-1]] < nums[i] {
			index_dequeue = index_dequeue[:len(index_dequeue)-1]
		}
		index_dequeue = append(index_dequeue, i)
	}

	res := make([]int, len(nums)-k+1)

	res[0] = nums[index_dequeue[0]]

	for i := k; i < len(nums); i++ {
		removeIndex := i - k
		if removeIndex == index_dequeue[0] {
			index_dequeue = index_dequeue[1:]
		}
		for len(index_dequeue) > 0 && nums[index_dequeue[len(index_dequeue)-1]] < nums[i] {
			index_dequeue = index_dequeue[:len(index_dequeue)-1]
		}
		index_dequeue = append(index_dequeue, i)
		res[i-k+1] = nums[index_dequeue[0]]
	}
	return res
}

func main() {
	fmt.Printf("%v\n", maxSlidingWindow([]int{1, 3, 1, 2, 0, 5}, 3))
	fmt.Printf("%v\n", maxSlidingWindow([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3))
	fmt.Printf("%v\n", maxSlidingWindow([]int{1, -1}, 1))
	fmt.Printf("%v\n", maxSlidingWindow([]int{1}, 1))
}

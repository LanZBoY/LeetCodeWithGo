package main

import "fmt"

func pivotIndex(nums []int) int {
	prefixSum := make([]int, len(nums))
	prefixSum[0] = nums[0]
	suffixSum := make([]int, len(nums))
	suffixSum[len(nums)-1] = nums[len(nums)-1]

	for i := 1; i < len(nums); i++ {
		back_i := len(nums) - i - 1

		prefixSum[i] = prefixSum[i-1] + nums[i]
		suffixSum[back_i] = suffixSum[back_i+1] + nums[back_i]
	}

	for i := range nums {
		leftSum := 0
		rightSum := 0
		if i != 0 {
			leftSum = prefixSum[i-1]
		}
		if i != len(nums)-1 {
			rightSum = suffixSum[i+1]
		}

		if leftSum == rightSum {
			return i
		}
	}

	return -1
}

func main() {
	fmt.Printf("%v\n", pivotIndex([]int{1, 7, 3, 6, 5, 6}))
	fmt.Printf("%v\n", pivotIndex([]int{1, 2, 3}))
	fmt.Printf("%v\n", pivotIndex([]int{2, 1, -1}))
}

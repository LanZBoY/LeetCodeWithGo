package main

import "fmt"

func isZeroArray(nums []int, queries [][]int) bool {

	for _, query := range queries {

		for i := query[0]; i <= query[1]; i++ {
			nums[i] = max(0, nums[i]-1)
		}
	}

	for _, num := range nums {
		if num > 0 {
			return false
		}
	}

	return true
}

func main() {

	fmt.Printf("%v\n", isZeroArray([]int{1, 0, 1}, [][]int{{0, 2}}))
	fmt.Printf("%v\n", isZeroArray([]int{4, 3, 2, 1}, [][]int{{1, 3}, {0, 2}}))
}

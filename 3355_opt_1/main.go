package main

import "fmt"

func isZeroArray(nums []int, queries [][]int) bool {

	for i := range nums {
		rangeCount := 0
		for j := range queries {
			if queries[j][0] <= i && queries[j][1] >= i {
				rangeCount++
			}
		}
		if nums[i] > rangeCount {
			return false
		}
	}

	return true
}

func main() {

	fmt.Printf("%v\n", isZeroArray([]int{1, 0, 1}, [][]int{{0, 2}}))
	fmt.Printf("%v\n", isZeroArray([]int{4, 3, 2, 1}, [][]int{{1, 3}, {0, 2}}))
}

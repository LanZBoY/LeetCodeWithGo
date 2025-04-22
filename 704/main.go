package main

import "fmt"

func search(nums []int, target int) int {
	i := 0
	j := len(nums) - 1

	for i <= j {
		mid := (i + j) / 2

		if nums[mid] > target {
			j = mid - 1
		} else if nums[mid] < target {
			i = mid + 1
		} else {
			return mid
		}
	}

	return -1
}
func main() {
	fmt.Println(search([]int{-1, 0, 3, 5, 9, 12}, 9))
	fmt.Println(search([]int{-1, 0, 3, 5, 9, 12}, 13))
	fmt.Println(search([]int{5}, 5))
}

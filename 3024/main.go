package main

import "fmt"

func triangleType(nums []int) string {

	if nums[0]+nums[1] <= nums[2] || nums[1]+nums[2] <= nums[0] || nums[0]+nums[2] <= nums[1] {
		return "none"
	}

	if nums[0] == nums[1] && nums[1] == nums[2] {
		return "equilateral"
	}

	if nums[0] == nums[1] || nums[1] == nums[2] || nums[2] == nums[0] {
		return "isosceles"
	}

	return "scalene"
}

func main() {
	fmt.Printf("%v\n", triangleType([]int{3, 3, 3}))
	fmt.Printf("%v\n", triangleType([]int{3, 4, 5}))
}

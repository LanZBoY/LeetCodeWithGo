package main

import "fmt"

func findDuplicate(nums []int) int {

	slow := nums[0]
	fast := nums[0]

	for {
		slow = nums[slow]
		fast = nums[nums[fast]]
		if slow == fast {
			break
		}
	}

	slow = nums[0]

	for slow != fast {
		slow = nums[slow]
		fast = nums[fast]
	}

	return slow
}

func main() {

	fmt.Printf("%v\n", findDuplicate([]int{1, 3, 4, 2, 2}))
	fmt.Printf("%v\n", findDuplicate([]int{3, 1, 3, 4, 2}))
	fmt.Printf("%v\n", findDuplicate([]int{3, 3, 3, 3, 3}))
}

package main

import "fmt"

func findDuplicate(nums []int) int {
	dupSet := make([]bool, len(nums))

	for _, num := range nums {
		if exist := dupSet[num]; exist {
			return num
		}
		dupSet[num] = true
	}

	return -1
}

func main() {

	fmt.Printf("%v\n", findDuplicate([]int{1, 3, 4, 2, 2}))
	fmt.Printf("%v\n", findDuplicate([]int{3, 1, 3, 4, 2}))
	fmt.Printf("%v\n", findDuplicate([]int{3, 3, 3, 3, 3}))
}

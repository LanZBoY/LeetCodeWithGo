package main

import "fmt"

func isZeroArray(nums []int, queries [][]int) bool {
	deltaArray := make([]int, len(nums)+1)

	for _, query := range queries {
		deltaArray[query[0]]++
		deltaArray[query[1]+1]--
	}
	operationCounts := make([]int, len(deltaArray))

	currentOperation := 0

	for i, delta := range deltaArray {
		currentOperation += delta
		operationCounts[i] = currentOperation
	}
	for i := range nums {
		if operationCounts[i] < nums[i] {
			return false
		}
	}
	return true
}

func main() {

	fmt.Printf("%v\n", isZeroArray([]int{1, 0, 1}, [][]int{{0, 2}}))
	fmt.Printf("%v\n", isZeroArray([]int{4, 3, 2, 1}, [][]int{{1, 3}, {0, 2}}))
}

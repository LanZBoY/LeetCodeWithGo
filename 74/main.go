package main

import "fmt"

func getElement(matrix [][]int, ele_i int) int {
	n := len(matrix[0])
	return matrix[ele_i/n][ele_i%n]
}

func searchMatrix(matrix [][]int, target int) bool {
	i := 0
	j := (len(matrix) * len(matrix[0])) - 1
	for i <= j {
		mid := (i + j) / 2
		mid_ele := getElement(matrix, mid)

		if mid_ele > target {
			j = mid - 1
		} else if mid_ele < target {
			i = mid + 1
		} else {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println(searchMatrix([][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}, 3))
	fmt.Println(searchMatrix([][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}, 13))
}

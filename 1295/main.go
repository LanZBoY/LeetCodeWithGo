package main

import "fmt"

func getDigits(num int) int {
	digitNum := 0

	for num > 0 {
		digitNum++
		num = num / 10
	}
	return digitNum
}

func findNumbers(nums []int) int {
	ret := 0
	for _, num := range nums {
		if digitNum := getDigits(num); digitNum%2 == 0 {
			ret++
		}

	}
	return ret
}

func main() {
	fmt.Printf("%v\n", findNumbers([]int{12, 345, 2, 6, 7896}))
	fmt.Printf("%v\n", findNumbers([]int{555, 901, 482, 1771}))
}

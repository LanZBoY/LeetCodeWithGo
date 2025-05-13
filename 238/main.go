package main

import "fmt"

func productExceptSelf(nums []int) []int {
	N := len(nums)
	prefixProduct := make([]int, N)
	suffixProduct := make([]int, N)

	prefixProduct[0] = nums[0]
	suffixProduct[N-1] = nums[N-1]
	for i := 1; i < N; i++ {
		revert_i := N - i - 1
		prefixProduct[i] = prefixProduct[i-1] * nums[i]
		suffixProduct[revert_i] = suffixProduct[revert_i+1] * nums[revert_i]
	}

	retProduct := make([]int, N)
	retProduct[0] = suffixProduct[1]
	retProduct[N-1] = prefixProduct[N-2]
	for i := 1; i < N-1; i++ {
		retProduct[i] = prefixProduct[i-1] * suffixProduct[i+1]
	}

	return retProduct
}

func main() {
	fmt.Printf("%v\n", productExceptSelf([]int{1, 2, 3, 4}))
	fmt.Printf("%v\n", productExceptSelf([]int{-1, 1, 0, -3, 3}))
}

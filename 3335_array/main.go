package main

import "fmt"

func lengthAfterTransformations(s string, t int) int {
	const MOD = 1000000007
	arr := [26]int{0}

	for _, r := range s {
		arr[r%'a']++
	}

	for range t {
		temp_arr := [26]int{0}
		for i := range 26 {
			if arr[i] == 0 {
				continue
			}

			v := arr[i]

			if i < 25 {
				temp_arr[i+1] = (temp_arr[i+1] + v) % MOD
			} else {
				temp_arr[0] = (temp_arr[0] + v) % MOD
				temp_arr[1] = (temp_arr[1] + v) % MOD
			}
		}
		arr = temp_arr
	}

	sum := 0
	for i := range 26 {

		sum = (sum + arr[i]) % MOD
	}
	return sum
}

func main() {
	fmt.Printf("%v\n", lengthAfterTransformations("abcyy", 2))
	fmt.Printf("%v\n", lengthAfterTransformations("azbk", 1))
	fmt.Printf("%v\n", lengthAfterTransformations("jqktcurgdvlibczdsvnsg", 7517))
}

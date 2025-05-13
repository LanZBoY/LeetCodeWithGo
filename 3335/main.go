package main

import "fmt"

func lengthAfterTransformations(s string, t int) int {
	counterMap := map[rune]int{}

	for _, r := range s {
		counterMap[r]++
	}

	for range t {
		tempMap := map[rune]int{}

		for k, v := range counterMap {
			if k < 'z' {
				tempMap[k+1] = (tempMap[k+1] + v) % 1000000007
			} else {
				tempMap['a'] = (tempMap['a'] + v) % 1000000007
				tempMap['b'] = (tempMap['b'] + v) % 1000000007
			}
		}
		counterMap = tempMap
	}

	sum := 0

	for _, v := range counterMap {
		sum = (sum + v) % 1000000007
	}

	return sum
}

func main() {
	fmt.Printf("%v\n", lengthAfterTransformations("jqktcurgdvlibczdsvnsg", 7517))
	fmt.Printf("%v\n", lengthAfterTransformations("abcyy", 2))
	fmt.Printf("%v\n", lengthAfterTransformations("azbk", 1))
}

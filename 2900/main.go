package main

import "fmt"

func getLongestSubsequence(words []string, groups []int) []string {
	N := len(words)

	ret := make([]string, 0)
	ret = append(ret, words[0])
	for i := 1; i < N; i++ {

		if groups[i] != groups[i-1] {
			ret = append(ret, words[i])
		}
	}
	return ret
}

func main() {

	fmt.Printf("%v\n", getLongestSubsequence([]string{"e", "a", "b"}, []int{0, 0, 1}))
	fmt.Printf("%v\n", getLongestSubsequence([]string{"a", "b", "c", "d"}, []int{1, 0, 1, 1}))
}

package main

import "fmt"

func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}
	countMap := map[byte]int{}

	for _, c := range s1 {
		countMap[byte(c)]++
	}
	i, j := 0, 0
	for j < len(s2) {
		r := s2[j]
		if _, ok := countMap[r]; ok && countMap[r] > 0 {
			countMap[r]--
			j++
			if j-i == len(s1) {
				return true
			}
		} else {
			temp_i := i
			for temp_i < j {
				if _, ok := countMap[s2[i]]; ok {
					countMap[s2[temp_i]]++
				}
				temp_i++
			}
			i++
			j = i
		}
	}
	return false
}

func main() {
	fmt.Printf("%v\n", checkInclusion("ab", "eidbaooo"))
	fmt.Printf("%v\n", checkInclusion("ab", "eidboaoo"))
	fmt.Printf("%v\n", checkInclusion("bc", "ad"))
	fmt.Printf("%v\n", checkInclusion("adc", "dcda"))
	fmt.Printf("%v\n", checkInclusion("hello", "ooolleoooleh"))
}

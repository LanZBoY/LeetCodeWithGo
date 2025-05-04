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
		c := s2[j]
		_, exist := countMap[c]

		if exist {
			countMap[c]--
			j++
			for i < j && countMap[c] < 0 {
				countMap[s2[i]]++
				i++
			}

			if j-i == len(s1) {
				return true
			}
		} else {
			for i < j {
				if _, ok := countMap[s2[i]]; ok {
					countMap[s2[i]]++
				}
				i++
			}
			j++
			i = j
		}

	}
	return false
}

func main() {
	fmt.Printf("%v\n", checkInclusion("hello", "ooolleoooleh"))
	fmt.Printf("%v\n", checkInclusion("ab", "eidbaooo"))
	fmt.Printf("%v\n", checkInclusion("ab", "eidboaoo"))
	fmt.Printf("%v\n", checkInclusion("adc", "dcda"))
	fmt.Printf("%v\n", checkInclusion("bc", "ad"))
}

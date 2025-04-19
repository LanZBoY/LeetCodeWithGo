package main

import "fmt"

func isValid(s string) bool {
	if len(s) == 0 {
		return false
	}
	c_stack := make([]rune, 0)
	closeMap := map[rune]rune{
		'(': ')',
		'[': ']',
		'{': '}',
	}
	for _, c := range s {
		if _, ok := closeMap[c]; ok {
			c_stack = append(c_stack, c)
		} else if len(c_stack) == 0 || c != closeMap[c_stack[len(c_stack)-1]] {
			return false
		} else {
			// pop
			c_stack = c_stack[:len(c_stack)-1]
		}
	}
	return len(c_stack) == 0
}

func main() {
	fmt.Printf("%v\n", isValid("()"))
	fmt.Printf("%v\n", isValid("()[]{}"))
	fmt.Printf("%v\n", isValid("(]"))
	fmt.Printf("%v\n", isValid("([])"))
}

package main

import (
	"fmt"
	"strings"
)

func main() {

	tests := []string{
		"A man, a plan, a canal: Panama",
		"race a car",
		" ",
	}

	for i, test := range tests {
		fmt.Printf("test %d: %v - %v\n", i+1, test, isPalindrome(test))
	}

}

func isPalindrome(s string) bool {

	lowerS := strings.ToLower(s)
	i, j := 0, len(s)-1

	for i < j {
		if !isValid(lowerS[i]) {
			i++
			continue
		}
		if !isValid(lowerS[j]) {
			j--
			continue
		}

		if lowerS[i] != lowerS[j] {
			return false
		}

		i++
		j--
	}
	return true
}

func isValid(a byte) bool {
	if (a >= 'a' && a <= 'z') || (a >= 'A' && a <= 'Z') || (a >= '0' && a <= '9') {
		return true
	}
	return false
}

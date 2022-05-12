package main

import (
	"fmt"
	"strings"
)

func isPalindrome(s string) bool {
	i, j := 0, len(s)-1
	for i < j {
		if !isValid(s[i]) {
			i++
			continue
		}
		if !isValid(s[j]) {
			j--
			continue
		}
		if !strings.EqualFold(string(s[i]), string(s[j])) {
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

func main() {
	test1 := "A man, a plan, a canal: Panama"
	test2 := "race a car"
	test3 := " "

	fmt.Println(isPalindrome(test1))
	fmt.Println(isPalindrome(test2))
	fmt.Println(isPalindrome(test3))
}

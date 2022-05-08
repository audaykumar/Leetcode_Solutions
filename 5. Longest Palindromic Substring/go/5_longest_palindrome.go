package main

import "fmt"

func main() {
	test1 := "babad"
	// test2 := "cbbd"

	result := longestPalindrome(test1)
	fmt.Printf("Result: %s", result)
}

func longestPalindrome(s string) string {
	temp_len := 0
	res := ""
	for i, _ := range s {
		odd := check(s, i, i)
		even := check(s, i, i+1)
	}
	return ""
}

func check(s string, lower int, upper int) string {
	for lower >= 0 && upper < len(s) && s[lower] == s[upper] {
		lower = lower - 1
		upper = upper + 1
	}

	res := s[lower+1 : upper]
	return res
}

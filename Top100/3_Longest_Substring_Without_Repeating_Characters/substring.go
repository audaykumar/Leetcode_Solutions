package main

import (
	"fmt"
	"math"
)

func main() {
	s1 := "abcafabcbb"
	// s2 := "bbbbb"
	// s3 := "pwwkew"
	// s4 := ""
	result := lengthOfLongestSubstring(s1)

	fmt.Printf("Length of longest substring: %d\n", result)
}

func lengthOfLongestSubstring(s string) int {
	start, maxLength := 0, 0
	charMap := make(map[rune]int)
	for i, c := range s {
		fmt.Printf("char: %c, start: %d, len:%d\n", c, start, maxLength)
		_, ok := charMap[c]
		// If the character is present in the map
		// and the start value is lesser than the position of the character move start to index + 1
		if ok && start <= charMap[c] {
			start = charMap[c] + 1
		} else {
			maxLength = int(math.Max(float64(maxLength), float64(i-start+1)))
		}
		charMap[c] = i
	}
	return maxLength
}

package main

import "fmt"

func main() {
	s := "aaabbbccddd"
	result := minDeletions(s)
	fmt.Printf("Result: %d\n", result)
}

func minDeletions(s string) int {
	var result int
	charMap := make(map[rune]int)
	used := make(map[int32]bool)
	for _, char := range s {
		charMap[char] += 1
	}

	for _, freq := range charMap {
		for freq > 0 && used[int32(freq)] {
			freq--
			result++
		}

		used[int32(freq)] = true
	}
	return result
}

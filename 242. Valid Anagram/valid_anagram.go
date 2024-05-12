package main

import (
	"fmt"
	"sort"
)

func main() {

	tests := [][]string{
		{"anagram", "nagaram"},
		{"rat", "card"},
	}

	for i, test := range tests {
		fmt.Printf("test %d: %v - %v\n", i+1, test, isAnagram(test[0], test[1]))
	}

}

func isAnagram(s string, t string) bool {

	if len(s) != len(t) {
		return false
	}

	// convert to byte array
	byteS := []byte(s)
	byteT := []byte(t)

	sort.Slice(byteS, func(i, j int) bool {
		return byteS[i] < byteS[j]
	})

	sort.Slice(byteT, func(i, j int) bool {
		return byteT[i] < byteT[j]
	})

	// fmt.Printf("%s\n", byteS)
	// fmt.Printf("%s\n", byteT)
	for i := 0; i < len(s); i++ {
		if byteS[i] != byteT[i] {
			return false
		}
	}

	return true

}

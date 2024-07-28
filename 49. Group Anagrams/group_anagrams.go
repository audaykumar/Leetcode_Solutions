package main

import (
	"fmt"
	"sort"
)

func main() {
	tests := [][]string{
		{"eat", "tea", "tan", "ate", "nat", "bat"},
		{""},
		{"a"},
	}

	for i, test := range tests {
		fmt.Printf("test %d: %v - %v\n", i+1, test, groupAnagrams(test))
	}

}

func groupAnagrams(strs []string) [][]string {

	result := [][]string{}

	anagramMap := make(map[string][]string)
	for _, s := range strs {
		byteS := []byte(s)
		sort.Slice(byteS, func(i, j int) bool {
			return byteS[i] < byteS[j]
		})

		if _, ok := anagramMap[string(byteS)]; !ok {
			anagramMap[string(byteS)] = []string{s}
		} else {
			list := anagramMap[string(byteS)]
			list = append(list, s)
			anagramMap[string(byteS)] = list
		}
	}

	for _, value := range anagramMap {
		result = append(result, value)
	}
	return result
}

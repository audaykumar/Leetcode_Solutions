package main

import "fmt"

func main() {

	tests := []string{
		"23",
		"",
	}

	for i, test := range tests {
		fmt.Printf("test %d: %v - %v\n", i+1, test, letterCombinations(test))
	}
}

var charMap = map[string][]string{
	"2": {"a", "b", "c"},
	"3": {"d", "e", "f"},
	"4": {"g", "h", "i"},
	"5": {"j", "k", "l"},
	"6": {"m", "n", "o"},
	"7": {"p", "q", "r", "s"},
	"8": {"t", "u", "v"},
	"9": {"w", "x", "y", "z"},
}

func letterCombinations(digits string) []string {
	result := []string{}

	if digits == "" {
		return result
	}

	result = solution(digits, "", result)

	return result
}

func solution(digits string, str string, result []string) []string {
	if digits == "" {
		result = append(result, str)
		return result
	}
	char := digits[:1]

	for i := 0; i < len(charMap[char]); i++ {
		newstr := str
		newstr += charMap[char][i]

		result = solution(digits[1:], newstr, result)

	}

	return result
}

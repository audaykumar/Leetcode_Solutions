package main

import (
	"fmt"
	"strconv"
)

func main() {
	tests := []string{
		"444947137",
		"00009",
		"39878",
		"0000",
		"54321",
	}

	for i, test := range tests {
		fmt.Printf("test %d: %v - %v\n", i+1, test, largestPalindromic(test))
	}

}

func largestPalindromic(num string) string {

	numMap := make(map[int]int, 0)

	for _, n := range num {
		n_int, _ := strconv.Atoi(string(n))
		numMap[n_int] += 1
	}

	leftStr := ""
	rightStr := ""
	middle := ""
	for i := 9; i >= 0; i-- {
		number := strconv.Itoa(i)
		for {
			if numMap[i] < 2 {
				break
			}

			if leftStr == "" && i == 0 {
				break
			}
			leftStr = leftStr + number
			rightStr = number + rightStr
			numMap[i] = numMap[i] - 2
		}
		if middle == "" && numMap[i] != 0 {
			middle = number
		}

	}
	result := leftStr + middle + rightStr

	return result

}

package main

import (
	"fmt"
	"math"
)

func main() {
	tests := [][]int{
		{7, 1, 5, 3, 6, 4},
		{7, 6, 4, 3, 1},
	}

	for i, test := range tests {
		fmt.Printf("test %d: %v - %v\n", i+1, test, maxProfit(test))
	}

}

func maxProfit(prices []int) int {
	maxProfit := 0
	start := prices[0]
	for i := 1; i < len(prices); i++ {
		if prices[i] < start {
			start = prices[i]
		} else {
			maxProfit = int(math.Max(float64(maxProfit), float64(prices[i]-start)))
		}

	}

	return maxProfit
}

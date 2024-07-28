package main

import (
	"fmt"
	"sort"
)

type testcase struct {
	nums []int
	k    int
}

func main() {
	tests := []testcase{
		{
			nums: []int{1, 1, 1, 1, 2, 2, 3},
			k:    2,
		},
		{
			nums: []int{1},
			k:    1,
		},
	}

	for i, test := range tests {
		fmt.Printf("test %d: num: %v, k: %v} - %v\n", i+1, test.nums, test.k, topKFrequent(test.nums, test.k))
	}
}

func topKFrequent(nums []int, k int) []int {

	numMap := make(map[int]int)

	for _, num := range nums {
		numMap[num] += 1
	}

	numSlice := make([]int, 0)
	for k := range numMap {
		numSlice = append(numSlice, k)
	}

	sort.Slice(numSlice, func(i, j int) bool {
		return numMap[numSlice[i]] >= numMap[numSlice[j]]
	})

	result := []int{}
	for i := 0; i < k; i++ {
		result = append(result, numSlice[i])
	}
	return result
}

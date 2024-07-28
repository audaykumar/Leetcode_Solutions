package main

import "fmt"

type testcase struct {
	nums   []int
	target int
}

func main() {

	tests := []testcase{
		{
			nums:   []int{-1, 0, 3, 5, 9, 12},
			target: 9,
		},
		{
			nums:   []int{-1, 0, 3, 5, 9, 12},
			target: 2,
		},
	}

	for i, test := range tests {
		fmt.Printf("test %d: {%v, %v} - %v\n", i+1, test.nums, test.target, search(test.nums, test.target))
	}

}

func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	left := 0
	right := len(nums) - 1

	for left <= right {

		middle := (left + right) / 2

		if nums[middle] == target {
			return middle
		} else if target > nums[middle] {
			left = middle + 1
		} else {
			right = middle - 1
		}
	}
	return -1
}

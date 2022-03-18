package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{34, 55, 79, 28, 46, 33, 2, 48, 31, -3, 84, 71, 52, -3, 93, 15, 21, -43, 57, -6, 86, 56, 94, 74, 83, -14, 28, -66, 46, -49, 62, -11, 43, 65, 77, 12, 47, 61, 26, 1, 13, 29, 55, -82, 76, 26, 15, -29, 36, -29, 10, -70, 69, 17, 49}
	result := threeSum(nums)
	fmt.Printf("Result: %v\n", result)
}

func threeSum(nums []int) [][]int {
	var result [][]int
	sort.Ints(nums)
	var target int
	for i := 0; i < len(nums); i++ {
		// The target sum we need from the other 2 numbers
		target = -1 * nums[i]
		// set front to the next number
		front := i + 1
		// set back to the last number
		back := len(nums) - 1
		// If we reach a positive number its target will be negative
		// which we have already covered so break out from the loop
		if target < 0 {
			break
		}

		// loop while front is lesser than back
		for front < back {
			sum := nums[front] + nums[back]

			// if the sum is less than target we need a smaller negative number
			// if sum is bigger than target we need a smaller positive number
			// Else we have target
			if sum < target {
				front++
			} else if sum > target {
				back--
			} else {
				triplet := []int{nums[i], nums[front], nums[back]}
				result = append(result, triplet)
				// Move front to the next unique number
				for (front < back) && (nums[front] == triplet[1]) {
					front++
				}
				// move back to the next unique number
				for (front < back) && (nums[back] == triplet[2]) {
					back--
				}
			}
		}

		for i+1 < len(nums) && nums[i+1] == nums[i] {
			i++
		}
	}
	return result
}

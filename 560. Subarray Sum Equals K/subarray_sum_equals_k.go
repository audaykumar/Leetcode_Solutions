package main

import "fmt"

type testcase struct {
	nums []int
	k    int
}

func main() {
	tests := []testcase{
		{
			nums: []int{1, 1, 1},
			k:    3,
		},
		{
			nums: []int{1, 2, -1, 3, -1, 2},
			k:    3,
		},
	}

	for i, test := range tests {
		fmt.Printf("test %d: num: %v, k: %v} - %v | %v\n", i+1, test.nums, test.k, subarraySum(test.nums, test.k), longestSubarraySum(test.nums, test.k))
	}
}

// Solution

// Calculate the sum for each index starting from 1
// Create a map with the sum as the key and the count of how many times we can reach that sum as value
// Basically when we calulcate the sum at an index it is the sum of all the nums before it
// if we subtract k from it that is the value from where we have to start to reach the current sum
// So if we find in the map the sum as the key then the value in the map at that point is the number of times
// we can reach that starting value

func subarraySum(nums []int, k int) int {

	count := 0
	sum := 0
	customMap := make(map[int]int)
	customMap[0] = 1

	for i := 0; i < len(nums); i++ {
		sum += nums[i]

		diff := sum - k

		if val, ok := customMap[diff]; ok {
			count += val
		}
		customMap[sum]++
	}

	return count
}

func longestSubarraySum(nums []int, k int) int {

	maxLen := 0
	sum := 0
	customMap := make(map[int]int)
	customMap[0] = 0

	for i := 0; i < len(nums); i++ {
		sum += nums[i]

		diff := sum - k

		if nums[i] == k {
			if maxLen < 1 {
				maxLen = 1
				continue
			}
		}
		if val, ok := customMap[diff]; ok {
			if i-val > maxLen {
				maxLen = i - val + 1
				fmt.Println(maxLen, i, val)
			}
		} else {
			customMap[sum] = i + 1
		}

	}

	fmt.Println(customMap)
	return maxLen
}

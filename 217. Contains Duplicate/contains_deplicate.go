package main

import "fmt"

func containsDuplicate(nums []int) bool {
	if len(nums) <= 1 {
		return false
	}
	numMap := make(map[int]struct{}, 0)

	for _, num := range nums {
		_, ok := numMap[num]
		if ok {
			return true
		} else {
			numMap[num] = struct{}{}
		}
	}
	return false
}

func main() {
	list := []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}

	result := containsDuplicate(list)
	fmt.Printf("Result: %v\n", result)
}

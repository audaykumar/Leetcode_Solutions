package main

import "fmt"

func main() {
	test1 := []int{1, 2, 3, 1}
	test2 := []int{1, 2, 3, 4}
	test3 := []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}

	fmt.Println(containsDuplicate(test1))
	fmt.Println(containsDuplicate(test2))
	fmt.Println(containsDuplicate(test3))
}

func containsDuplicate(nums []int) bool {
	if len(nums) <= 1 {
		return false
	}

	numMap := make(map[int]struct{})

	for _, num := range nums {
		if _, ok := numMap[num]; ok {
			return true
		}
		numMap[num] = struct{}{}
	}
	return false
}

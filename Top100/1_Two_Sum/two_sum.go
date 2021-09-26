package main

import "fmt"

func main() {
	list := []int{2, 11, 15, 13, 27, 7, 41}

	result := twoSum(list, 9)
	fmt.Printf("Result: %v\n", result)
}

func twoSum(nums []int, target int) []int {
	numMap := make(map[int]int)
	var result []int
	for index, value := range nums {
		compliment := target - value
		if val, ok := numMap[compliment]; ok {
			result = []int{val, index}
			return result
		}
		numMap[value] = index
	}
	return result
}

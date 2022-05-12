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
	mymap := make(map[int]int)

	for _, n := range nums {
		mymap[n]++
		if mymap[n] > 1 {
			return true
		}
	}

	return false
}

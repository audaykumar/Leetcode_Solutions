package main

import "fmt"

func longestConsecutive(nums []int) int {

	max := 0

	set := make(map[int]struct{})

	for _, num := range nums {
		if _, ok := set[num]; !ok {
			set[num] = struct{}{}
		}
	}

	count := 0
	for _, num := range nums {
		count = 1
		delete(set, num)
		// if _, ok := set[num-1]; !ok {
		// 	j := num
		// 	flag := true
		// 	for flag {
		// 		if _, ok := set[j]; ok {
		// 			count++
		// 			j++
		// 		} else {
		// 			flag = false
		// 		}
		// 	}
		// 	if count > max {
		// 		max = count
		// 	}
		// }

		right := num + 1
		left := num - 1
		for {
			if _, ok := set[left]; !ok {
				break
			}
			delete(set, left)
			count++
			left--
		}

		for {
			if _, ok := set[right]; !ok {
				break
			}
			delete(set, right)
			count++
			right++
		}

		if count > max {
			max = count
		}

	}
	return max
}

func main() {

	test1 := []int{100, 4, 200, 1, 3, 2, 100}

	test2 := []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}

	fmt.Println(longestConsecutive(test1))

	fmt.Println(longestConsecutive(test2))
}

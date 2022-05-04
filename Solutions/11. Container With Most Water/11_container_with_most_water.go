package main

import (
	"fmt"
)

func main() {

	// test1 := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}

	test1 := []int{2, 3, 4, 5, 18, 17, 6}
	fmt.Printf("Max Area: %v\n", maxArea(test1))
}

func maxArea(height []int) int {
	i := 0
	j := len(height) - 1
	temp_area := 0
	max_area := 0
	for {
		if i == j {
			break
		}
		// temp_area = (j - i) * int(math.Min(float64(height[i]), float64(height[j])))
		if height[i] < height[j] {
			temp_area = (j - i) * height[i]
			i++
		} else {
			temp_area = (j - i) * height[j]
			j--
		}
		if temp_area > max_area {
			max_area = temp_area
		}
	}
	return max_area
}

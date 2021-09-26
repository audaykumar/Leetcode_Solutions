package main

import (
	"fmt"
	"math"
)

func main() {

	nums := []int{2, 2, -3, 4, 5, -6, 7, 8, -9, 10}

	result := maxProduct(nums)

	fmt.Printf("Result: %d", result)
}

func maxProduct(nums []int) int {
	length := len(nums) - 1
	product1 := nums[0]
	product2 := nums[length]
	// Check for contiguous prouct from fornt and back
	for i, tempProduct, tempProduct2 := 0, 1, 1; i <= length; i++ {
		// Checking for product from front of array
		// If condition to check if the element is 0
		if nums[i] == 0 {
			tempProduct = 0
			product1 = int(math.Max(float64(tempProduct), float64(product1)))
			tempProduct = 1
		} else {
			tempProduct = tempProduct * nums[i]
			product1 = int(math.Max(float64(tempProduct), float64(product1)))
			product1 = int(math.Max(float64(nums[i]), float64(product1)))
		}

		// Checking for product from back of array
		if nums[length-i] == 0 {
			tempProduct2 = 0
			product2 = int(math.Max(float64(tempProduct2), float64(product2)))
			tempProduct2 = 1
		} else {
			tempProduct2 = tempProduct2 * nums[length-i]
			product2 = int(math.Max(float64(tempProduct2), float64(product2)))
			product2 = int(math.Max(float64(nums[length-i]), float64(product2)))
		}
	}

	product := int(math.Max(float64(product1), float64(product2)))
	return product
}

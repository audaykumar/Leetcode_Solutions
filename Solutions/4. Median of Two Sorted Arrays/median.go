package main

import (
	"fmt"
	"sort"
)

func main() {
	nums1 := []int{1, 3}
	nums2 := []int{2, 4}

	result := findMedianSortedArrays(nums1, nums2)
	fmt.Printf("Result: %g", result)
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var median float64
	merged := append(nums1, nums2...)
	sort.Ints(merged)
	if len(merged)%2 == 0 {
		median = float64(merged[len(merged)/2]+merged[(len(merged)/2)-1]) / 2
	} else {
		median = float64(merged[len(merged)/2])
	}

	return median
}

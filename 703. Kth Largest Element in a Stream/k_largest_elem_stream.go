package main

import (
	"fmt"
	"sort"
)

func main() {

	kLarge := Constructor(3, []int{4, 5, 8, 2})
	fmt.Println(kLarge)
	fmt.Println(kLarge.Add(3), kLarge)
}

type KthLargest struct {
	k    int
	nums []int
}

func Constructor(k int, nums []int) KthLargest {

	sort.Slice(nums, func(i, j int) bool {
		return nums[i] > nums[j]
	})

	kLrg := KthLargest{
		nums: nums,
		k:    k,
	}
	return kLrg
}

func (this *KthLargest) Add(val int) int {

	i := 0

	for ; i < len(this.nums); i++ {
		if val >= this.nums[i] {
			break
		}
	}

	this.nums = append(this.nums, 0)
	copy(this.nums[i+1:], this.nums[i:])
	this.nums[i] = val
	return this.nums[this.k-1]
}

/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */

//  2, 3, 4, 4, 5, 5, 8, 9, 10

// 8, 5 , 4 ,2

package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	// test1 := &ListNode{
	// 	Val: 1,
	// 	Next: &ListNode{
	// 		Val: 2,
	// 		Next: &ListNode{
	// 			Val: 3,
	// 			Next: &ListNode{
	// 				Val:  4,
	// 				Next: nil,
	// 			},
	// 		},
	// 	},
	// }

	// test2 := &ListNode{
	// 	Val:  1,
	// 	Next: nil,
	// }

	var test3 *ListNode
	test3 = nil

	result := swapPairs(test3)

	for {
		fmt.Printf("%d", result.Val)
		if result.Next == nil {
			break
		}
		result = result.Next
	}
}

/**
 * Definition for singly-linked list.
 */

func swapPairs(head *ListNode) *ListNode {
	var temp1 *ListNode
	var temp2 *ListNode
	var temp3 *ListNode
	count := 0
	root := head
	if head == nil {
		return nil
	} else if head.Next == nil {
		return head
	} else {
		root = head.Next
	}

	for {
		if head.Next == nil {
			break
		}
		if count%2 == 0 {
			// fmt.Printf("Count: %v, Value of head in mod 2: %v\n", count, head.Val)

			temp1 = head

			temp2 = head.Next.Next

			// fmt.Printf("value in head.next: %v\n", head.Next.Val)
			head = head.Next
			if temp3 != nil {
				temp3.Next = head
			}

			head.Next = temp1

			head.Next.Next = temp2

			head = head.Next
			count = count + 1
		} else {
			// fmt.Printf("Count: %v, Value in else: %v\n", count, head.Val)
			temp3 = head
			head = head.Next
			count = count + 1
		}
	}

	return root
}

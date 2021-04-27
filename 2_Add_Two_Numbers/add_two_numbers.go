package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	l1 := &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val:  3,
				Next: nil,
			},
		},
	}

	l2 := &ListNode{
		Val: 5,
		Next: &ListNode{
			Val: 6,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}
	ans := addTwoNumbers(l1, l2)
	for {
		fmt.Printf("%d", ans.Val)
		if ans.Next == nil {
			break
		}
		ans = ans.Next
	}

}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	ans := &ListNode{
		Val:  0,
		Next: nil,
	}

	carry := 0
	temp := ans
	num1 := l1.Val
	num2 := l2.Val
	// Run an infinite loop to go over all the nodes
	for {
		// Update the value of the node
		ans.Val = (num1 + num2 + carry) % 10
		// calculate the carry for the next two numbers
		carry = (num1 + num2 + carry) / 10

		// Check here if there are no more nodes left in either of the 2 linked lists.
		if l1.Next == nil && l2.Next == nil {
			if carry != 0 {
				ans.Next = &ListNode{
					Val:  carry,
					Next: nil,
				}
			}
			break
		}
		//  if nodes are present still check in which list it is present
		if l1.Next == nil {
			num1 = 0
		} else {
			l1 = l1.Next
			num1 = l1.Val
		}
		if l2.Next == nil {
			num2 = 0
		} else {
			l2 = l2.Next
			num2 = l2.Val
		}

		// Set the next node of the list
		ans.Next = &ListNode{
			Val:  0,
			Next: nil,
		}
		ans = ans.Next
	}

	return temp
}

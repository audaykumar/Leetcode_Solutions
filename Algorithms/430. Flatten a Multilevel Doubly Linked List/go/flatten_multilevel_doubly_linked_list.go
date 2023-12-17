package main

import "fmt"

// Definition for a Node.
type Node struct {
	Val   int
	Prev  *Node
	Next  *Node
	Child *Node
}

func flatten(root *Node) *Node {

	if root == nil {
		return root
	}
	DFS(root)
	return root

	// Alternate Solution
	// node := root
	// for node != nil {

	// 	child := node.Child
	// 	if child != nil {
	// 		next := node.Next

	// 		node.Next = child
	// 		child.Prev = node

	// 		for child.Next != nil {
	// 			child = child.Next
	// 		}

	// 		child.Next = next
	// 		if next != nil {
	// 			next.Prev = child
	// 		}
	// 		node.Child = nil
	// 	}

	// 	node = node.Next
	// }

	// return root
}

func DFS(node *Node) *Node {

	if node.Child != nil {
		next := node.Next

		// New next node
		node.Next = DFS(node.Child)

		// Set the previous for the next node to current node
		if node.Next != nil {
			node.Next.Prev = node
		}

		// Set the child of the current node to nil
		node.Child = nil

		// Get to the end of the sub Linked-List
		temp := node
		for {
			if temp.Next == nil {
				break
			}
			temp = temp.Next
		}
		// Set the next node as the original next node
		temp.Next = next

		// Set the previous of the original next node to the new previous
		if next != nil {
			next.Prev = temp
		}
		// node = node.Next
	}

	if node.Next == nil {
		return node
	} else {
		DFS(node.Next)
		return node
	}
}

func cutomPrint(head *Node) {
	for {
		if head.Next == nil {
			fmt.Printf("Val: %d - ", head.Val)
			if head.Prev != nil {
				fmt.Printf("Prev Val: %d - ", head.Prev.Val)
			}
			break
		}
		fmt.Printf("Val: %d - ", head.Val)
		if head.Prev != nil {
			fmt.Printf("Prev Val: %d - ", head.Prev.Val)
		}
		fmt.Printf("Next Val: %d\n", head.Next.Val)
		head = head.Next
	}
}

func main() {
	// test1 := &Node{
	// 	Val:   1,
	// 	Child: nil,
	// 	Next:  nil,
	// 	Prev:  nil,
	// }

	test2 := &Node{
		Val: 1,
		Child: &Node{Val: 2,
			Child: &Node{Val: 3,
				Child: nil,
				Prev:  nil,
				Next:  nil},
			Next: nil,
			Prev: nil},
		Next: nil,
		Prev: nil,
	}

	// cutomPrint(flatten(test1))

	cutomPrint(flatten(test2))

}

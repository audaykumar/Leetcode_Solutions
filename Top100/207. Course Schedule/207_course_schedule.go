package main

import (
	"fmt"
)

type input struct {
	numCourses    int
	prerequisites [][]int
}

func main() {
	// test1 := input{numCourses: 2, prerequisites: [][]int{{1, 0}}}
	test2 := input{numCourses: 3, prerequisites: [][]int{{1, 0}, {0, 2}, {2, 1}, {2, 0}}}

	// 20 [[0,10],[3,18],[5,5],[6,11],[11,14],[13,1],[15,1],[17,4]]
	// test3 := input{numCourses: 20, prerequisites: [][]int{{0, 10}, {3, 18}, {5, 5}, {6, 11}, {11, 14}, {13, 1}, {15, 1}, {17, 4}}}

	result := canFinish(test2.numCourses, test2.prerequisites)
	fmt.Printf("Result: %t\n", result)
}

func canFinish(numCourses int, prerequisites [][]int) bool {
	newmap := make(map[int][]int, numCourses)

	for _, v := range prerequisites {
		newmap[v[0]] = append(newmap[v[0]], v[1])
	}

	nodeVisited := make([]int, numCourses)
	for i := 0; i < numCourses; i++ {
		check := checkNodes(newmap, nodeVisited, i)

		if !check {
			return false
		}
	}

	return true
}

func checkNodes(nodesGraph map[int][]int, nodeVisited []int, node int) bool {

	// Check if already tested this node
	if nodeVisited[node] == 2 {
		return true
	}

	// Check if we find a loop while checking the nodes
	if nodeVisited[node] == 1 {
		return false
	}

	nodeVisited[node] = 1
	for _, x := range nodesGraph[node] {
		if !checkNodes(nodesGraph, nodeVisited, x) {
			return false
		}
	}
	nodeVisited[node] = 2

	return true
}

// func contains(arr []int, val int) bool {
// 	fmt.Printf("Array: %v\n", arr)
// 	fmt.Printf("Value to check: %v\n", val)
// 	for _, elem := range arr {
// 		if elem == val {
// 			return true
// 		}
// 	}
// 	return false
// }

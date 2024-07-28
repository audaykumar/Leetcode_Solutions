package main

import "fmt"

func main() {

	tests := [][]int{
		{2, 5, 1},
		{3, 2, 2},
	}

	for i, test := range tests {
		fmt.Printf("test %d: %v - %v\n", i+1, test, longestString(test[0], test[1], test[2]))
	}
}

func longestString(x int, y int, z int) int {
	ab := ""
	parent := ""
	for z > 0 {
		ab += "AB"
		z--
	}

	for x > 0 && y > 0 {
		parent += "AABB"
		x--
		y--
	}
	if x > 0 {
		parent += "AA"
	} else if y > 0 {
		parent = "BB" + parent
	}

	parent += ab
	return len(parent)
}

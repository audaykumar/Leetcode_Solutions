package main

import (
	"fmt"
	"math"
)

func kthGrammar(n int, k int) int {
	if n == 1 {
		return 0
	}

	size := int(math.Pow(2, float64(n-1)))

	if k <= size/2 {
		return kthGrammar(n-1, k)
	} else {
		return 1 - kthGrammar(n-1, k-(size/2))
	}

}

func main() {

	fmt.Printf("N: %v K: %v Res: %v\n", 1, 1, kthGrammar(1, 1))

	fmt.Printf("N: %v K: %v Res: %v\n", 2, 1, kthGrammar(2, 1))

	fmt.Printf("N: %v K: %v Res: %v\n", 30, 45, kthGrammar(30, 43))
}

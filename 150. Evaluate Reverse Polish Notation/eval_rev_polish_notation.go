package main

import (
	"fmt"
	"strconv"
)

func main() {

	tests := [][]string{
		{"2", "1", "+", "3", "*"},
		{"4", "13", "5", "/", "+"},
		{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"},
	}

	for i, test := range tests {
		fmt.Printf("test %d: %v - %v\n", i+1, test, evalRPN(test))
	}
}

type stack []int

func (s *stack) IsEmpty() bool {
	return len(*s) == 0
}

func (s *stack) Push(elem int) {
	*s = append(*s, elem)
}

func (s *stack) Pop() (int, bool) {
	if s.IsEmpty() {
		return 0, false
	}
	index := len(*s) - 1
	elem := (*s)[index]
	*s = (*s)[:index]

	return elem, true
}

func evalRPN(tokens []string) int {

	s := make(stack, 0)

	for _, token := range tokens {

		switch token {
		case "+":
			num1, _ := s.Pop()
			num2, _ := s.Pop()
			s.Push(num1 + num2)
		case "-":
			num1, _ := s.Pop()
			num2, _ := s.Pop()
			s.Push(num2 - num1)
		case "*":
			num1, _ := s.Pop()
			num2, _ := s.Pop()
			s.Push(num2 * num1)
		case "/":
			num1, _ := s.Pop()
			num2, _ := s.Pop()
			s.Push(num2 / num1)
		default:
			num, _ := strconv.ParseInt(token, 10, 64)
			s.Push(int(num))
		}

	}

	val, _ := s.Pop()
	return val
}

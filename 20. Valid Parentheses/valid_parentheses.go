package main

import "fmt"

func main() {

	tests := []string{
		"()",
		"()[]{}",
		"(]",
	}

	for i, test := range tests {
		fmt.Printf("test %d: %v - %v\n", i+1, test, isValid(test))
	}

}

type stack []rune

func (s stack) Push(c rune) stack {
	return append(s, c)
}

func (s stack) Pop() (rune, stack) {
	slen := len(s)

	if slen == 0 {
		return '#', s
	}

	toRet := s[slen-1]
	s = s[:slen-1]
	return toRet, s
}

func (s stack) Top() rune {
	slen := len(s)

	if slen == 0 {
		return '#'
	}
	toRet := s[slen-1]
	return toRet
}

func isValid(s string) bool {

	st := make(stack, 0)
	var popC rune
	for _, c := range s {
		if c == '(' || c == '{' || c == '[' {
			st = st.Push(c)
			continue
		}
		if c == ')' {
			popC, st = st.Pop()
			if popC != '(' || popC == '#' {
				return false
			}
			continue
		}
		if c == '}' {
			popC, st = st.Pop()
			if popC != '{' || popC == '#' {
				return false
			}
			continue
		}
		if c == ']' {
			popC, st = st.Pop()
			if popC != '[' || popC == '#' {
				return false
			}
			continue
		}

	}

	return len(st) == 0
}

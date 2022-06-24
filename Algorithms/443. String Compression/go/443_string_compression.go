package main

import "fmt"

func compress(chars []byte) int {
	i := 0
	ansindex := 0

	l := len(chars)

	for i < l {
		j := i
		for j < l && chars[j] == chars[i] {
			j++
		}
		freq := j - i

		chars[ansindex] = chars[i]
		ansindex++
		if freq > 1 {
			temp := fmt.Sprintf("%d", freq)

			for _, c := range []byte(temp) {
				chars[ansindex] = c
				ansindex++
			}
		}
		i = j
	}
	return ansindex
}

func main() {
	test1 := []byte{'a', 'a', 'b', 'b', 'c', 'c', 'c'}
	test2 := []byte{'a'}
	test3 := []byte{'a', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b'}

	fmt.Printf("%c\n %v\n", test1, compress(test1))
	fmt.Printf("%c\n %v\n", test2, compress(test2))
	fmt.Printf("%c\n %v\n", test3, compress(test3))
}

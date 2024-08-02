package main

import (
	"fmt"
	"strings"
)

type testcase struct {
	s        string
	wordDict []string
}

func main() {
	tests := []testcase{
		// {
		// 	s:        "leetcode",
		// 	wordDict: []string{"leet", "code"},
		// },
		{
			s:        "applepenapple",
			wordDict: []string{"apple", "pen"},
		},
		{
			s:        "catsandog",
			wordDict: []string{"cats", "dog", "sand", "and", "cat"},
		},
		// {
		// 	s:        "bb",
		// 	wordDict: []string{"a", "b", "bbb", "bbbb"},
		// },
		{
			s:        "aaaaaaa",
			wordDict: []string{"aaaa", "aaa"},
		},
	}

	for i, test := range tests {
		fmt.Printf("test %d: %s |  %v - %v \n", i+1, test.s, test.wordDict, wordBreak(test.s, test.wordDict))
	}
}

func wordBreak(s string, wordDict []string) bool {

	q := []string{s}
	memo := make(map[string]bool)
	for len(q) != 0 {
		remaining := q[0]
		q = q[1:] // Dequeue
		if _, ok := memo[remaining]; ok {
			continue
		}
		for _, word := range wordDict {
			fmt.Println(remaining, word)
			if word == remaining {
				return true
			}
			if strings.HasPrefix(remaining, word) {
				q = append(q, remaining[len(word):]) // Enqueue
				memo[remaining] = true
			}
			fmt.Println(q)
		}
		fmt.Println(memo)
	}
	return false
}

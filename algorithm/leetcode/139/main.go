package main

import "fmt"

func main() {
	s := "leetcode"
	wordDict := []string{"leet", "code"}
	fmt.Println("word break:", wordBreak(s, wordDict))
}

func wordBreak(s string, wordDict []string) bool {
	var maxLen int
	words := make(map[string]bool, len(wordDict))
	for _, word := range wordDict {
		words[word] = true
		maxLen = max(maxLen, len(word))
	}

	n := len(s)
	f := make([]bool, n+1)
	f[0] = true
	for i := 1; i <= n; i++ {
		for j := i - 1; j >= max(i-maxLen, 0); j-- {
			if f[j] && words[s[j:i]] {
				f[i] = true
			}
		}
	}
	return f[n]
}

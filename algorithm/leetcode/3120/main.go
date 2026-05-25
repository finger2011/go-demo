package main

import "fmt"

func main() {
	word := "aaAbcBC"
	fmt.Println("num:", numberOfSpecialChars(word))
}

func numberOfSpecialChars(word string) int {
	s := make(map[rune]bool, 52)
	for _, ch := range word {
		s[ch] = true
	}
	var ans int
	for ch := 'a'; ch <= 'z'; ch++ {
		if s[ch] && s[ch-'a'+'A'] {
			ans++
		}
	}
	return ans
}

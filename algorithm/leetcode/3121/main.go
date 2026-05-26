package main

import "fmt"

func main() {
	word := "AbcbDBdD"
	fmt.Println("num:", numberOfSpecialChars(word))
}

func numberOfSpecialChars(word string) int {
	s := [26]int{}
	var ans int
	for _, ch := range word {
		if ch >= 'a' && ch <= 'z' {
			switch s[ch-'a'] {
			case 0:
				s[ch-'a'] = 1
			case 2:
				s[ch-'a'] = 3
				ans--
			default:
			}
		} else {
			switch s[ch-'A'] {
			case 0:
				s[ch-'A'] = 3
			case 1:
				s[ch-'A'] = 2
				ans++
			default:
			}
		}
	}
	return ans
}

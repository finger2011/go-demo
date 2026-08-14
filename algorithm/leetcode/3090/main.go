package main

import "fmt"

func main() {
	s := "bcbbbcba"
	fmt.Println("max:", maximumLengthSubstring(s))
}

func maximumLengthSubstring(s string) int {
	chars := make([]int, 26)
	var ans, start int
	for i, ch := range s {
		if chars[ch-'a'] < 2 {
			chars[ch-'a']++
			ans = max(ans, i-start+1)
		} else {
			j := start
			for ; s[j] != byte(ch); j++ {
				chars[s[j]-'a']--
			}
			start = j + 1
		}
	}
	return ans
}

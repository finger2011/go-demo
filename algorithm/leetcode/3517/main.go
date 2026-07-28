package main

import "fmt"

func main() {
	s := "babab"
	fmt.Println("smallestPalindrome:", smallestPalindrome(s))
}

func smallestPalindrome(s string) string {
	n := len(s)
	has := (n % 2) != 0
	chars := make([]int, 26)
	ans := make([]byte, n)
	for _, ch := range s {
		chars[ch-'a']++
	}
	var pos int
	var mid, ch byte
	for i, num := range chars {
		if num == 0 {
			continue
		}
		ch = byte('a' + i)
		if has && (num%2) != 0 {
			num--
			mid = ch
		}
		for j := num / 2; j > 0; j-- {
			ans[pos] = ch
			ans[n-1-pos] = ch
			pos++
		}
	}
	if has {
		ans[n/2] = mid
	}
	return string(ans)
}

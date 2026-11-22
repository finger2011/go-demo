package main

import "fmt"

func main() {
	s := "babad"
	fmt.Println("res:", longestPalindrome(s))
}

func longestPalindrome(s string) string {
	n := len(s)
	ansLeft, ansRight := 0, 0

	// 奇回文串
	for i := range n {
		l, r := i, i
		for l >= 0 && r < n && s[l] == s[r] {
			l--
			r++
		}
		if r-l-1 > ansRight-ansLeft {
			ansLeft = l + 1
			ansRight = r // 左闭右开区间
		}
	}

	// 偶回文串
	for i := range n - 1 {
		l, r := i, i+1
		for l >= 0 && r < n && s[l] == s[r] {
			l--
			r++
		}
		if r-l-1 > ansRight-ansLeft {
			ansLeft = l + 1
			ansRight = r // 左闭右开区间
		}
	}

	return s[ansLeft:ansRight]
}

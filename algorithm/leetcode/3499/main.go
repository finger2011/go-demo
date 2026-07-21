package main

import (
	"fmt"
	"math"
)

func main() {
	s := "01"
	fmt.Println("max:", maxActiveSectionsAfterTrade(s))
}

func maxActiveSectionsAfterTrade(s string) int {
	var ans, mx int
	pre0 := math.MinInt
	cnt := 0
	for i := range len(s) {
		cnt++
		if i == len(s)-1 || s[i] != s[i+1] { // i 是这一段的末尾
			if s[i] == '1' {
				ans += cnt
			} else {
				mx = max(mx, pre0+cnt)
				pre0 = cnt
			}
			cnt = 0
		}
	}
	return ans + mx
}

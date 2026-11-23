package main

import (
	"fmt"
	"slices"
)

func main() {
	nums := []int{9, 4, 7, 2, 10}
	fmt.Println("res:", longestArithSeqLength(nums))
}

func longestArithSeqLength(nums []int) int {
	ans := 2
	m := slices.Max(nums)
	f := make([]int, m+1)
	for d := -m; d <= m; d++ {
		clear(f)
		for _, x := range nums {
			if 0 <= x-d && x-d <= m {
				f[x] = f[x-d] + 1
				ans = max(ans, f[x])
			} else {
				f[x] = 1
			}
		}
	}
	return ans
}

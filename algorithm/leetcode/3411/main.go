package main

import (
	"fmt"
	"slices"
)

func main() {
	nums := []int{2, 3, 4, 5, 6}
	fmt.Println("ans:", maxLength(nums))
}

func maxLength(nums []int) int {
	mx := slices.Max(nums)
	allLcm := 1
	for _, x := range nums {
		allLcm = lcm(allLcm, x)
	}
	ans := 0

	for i := range nums {
		m, l, g := 1, 1, 0
		for j := i; j < len(nums) && m <= allLcm*mx; j++ {
			x := nums[j]
			m *= x
			l = lcm(l, x)
			g = gcd(g, x)
			if m == l*g {
				ans = max(ans, j-i+1)
			}
		}
	}
	return ans
}

func gcd(a, b int) int {
	for a != 0 {
		a, b = b%a, a
	}
	return b
}
func lcm(a, b int) int {
	return a / gcd(a, b) * b
}

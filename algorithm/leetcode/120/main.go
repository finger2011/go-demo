package main

import (
	"fmt"
	"math"
)

func main() {
	triangle := [][]int{{2}, {3, 4}, {6, 5, 7}, {4, 1, 8, 3}}
	fmt.Println("min:", minimumTotal2(triangle))
}

func minimumTotal(triangle [][]int) int {
	m := len(triangle)
	if m == 1 {
		return triangle[0][0]
	}
	n := len(triangle[m-1])
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	dp[0][0] = triangle[0][0]
	ans := math.MaxInt
	for i := 1; i < m; i++ {
		k := len(triangle[i])
		for j := range k {
			if j == 0 {
				dp[i][j] = dp[i-1][j] + triangle[i][j]
			} else if j == k-1 {
				dp[i][j] = dp[i-1][j-1] + triangle[i][j]
			} else {
				dp[i][j] = min(dp[i-1][j-1], dp[i-1][j]) + triangle[i][j]
			}
			if i == m-1 {
				ans = min(ans, dp[i][j])
			}
		}

	}

	return ans
}

// O(n) 额外空间
func minimumTotal2(triangle [][]int) int {
	m := len(triangle)
	if m == 1 {
		return triangle[0][0]
	}
	n := len(triangle[m-1])
	dp := make([]int, n)
	dp[0] = triangle[0][0]
	ans := math.MaxInt
	for i := 1; i < m; i++ {
		k := len(triangle[i])
		for j := k - 1; j >= 0; j-- {
			if j == 0 {
				dp[j] = dp[j] + triangle[i][j]
			} else if j == k-1 {
				dp[j] = dp[j-1] + triangle[i][j]
			} else {
				dp[j] = min(dp[j-1], dp[j]) + triangle[i][j]
			}
			if i == m-1 {
				ans = min(ans, dp[j])
			}
		}
	}
	return ans
}

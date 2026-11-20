package main

import "fmt"

func main() {
	m, n := 3, 7
	fmt.Println("path:", uniquePaths2(m, n))
}

// dp[m][n] = dp[m-1][n] + dp[m][n-1]
func uniquePaths(m int, n int) int {
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	for i := range m {
		for j := range n {
			if i == 0 && j == 0 {
				dp[i][j] = 1
			} else if i == 0 {
				dp[i][j] = dp[i][j-1]
			} else if j == 0 {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = dp[i][j-1] + dp[i-1][j]
			}
		}
	}
	return dp[m-1][n-1]
}

// 因为之前状态无用，所以直接使用一维数组存即可
func uniquePaths2(m int, n int) int {
	dp := make([]int, n)
	for i := range m {
		for j := range n {
			if i == 0 {
				dp[j] = 1
			} else if j > 0 {
				dp[j] += dp[j-1]
			}
		}
	}
	return dp[n-1]
}

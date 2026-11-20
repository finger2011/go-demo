package main

import "fmt"

func main() {
	matrix := [][]int{{2, 1, 3}, {6, 5, 4}, {7, 8, 9}} // 13
	fmt.Println("min:", minFallingPathSum(matrix))
}

func minFallingPathSum(matrix [][]int) int {
	m, n := len(matrix), len(matrix[0])
	dp := make([][]int, m)
	dp[0] = matrix[0]

	for i := 1; i < m; i++ {
		dp[i] = make([]int, n)
		for j := range n {
			if j == 0 {
				dp[i][j] = min(dp[i-1][j], dp[i-1][j+1]) + matrix[i][j]
			} else if j == n-1 {
				dp[i][j] = min(dp[i-1][j], dp[i-1][j-1]) + matrix[i][j]
			} else {
				dp[i][j] = min(dp[i-1][j], dp[i-1][j+1], dp[i-1][j-1]) + matrix[i][j]
			}
		}
	}
	ans := dp[m-1][0]
	for i := 1; i < n; i++ {
		ans = min(ans, dp[m-1][i])
	}
	return ans
}

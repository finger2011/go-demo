package main

import "fmt"

func main() {
	// grid := [][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}} // 7
	grid := [][]int{{1, 2, 3}, {4, 5, 6}} // 12
	fmt.Println("min:", minPathSum(grid))
}

func minPathSum(grid [][]int) int {
	n := len(grid)
	m := len(grid[0])
	dp := make([]int, m)
	for i := range n {
		for j := range m {
			if i == 0 && j == 0 {
				dp[j] = grid[i][j]
			} else if i == 0 {
				dp[j] = dp[j-1] + grid[i][j]
			} else if j == 0 {
				dp[j] += grid[i][j]
			} else {
				dp[j] = min(dp[j], dp[j-1]) + grid[i][j]
			}
		}
	}
	return dp[m-1]
}

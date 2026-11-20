package main

import "fmt"

func main() {
	// obstacleGrid := [][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}} // 2
	obstacleGrid := [][]int{{1, 0}} // 0
	fmt.Println("path:", uniquePathsWithObstacles(obstacleGrid))
}

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m, n := len(obstacleGrid), len(obstacleGrid[0])
	dp := make([]int, n)
	dp[0] = 1
	for i := range m {
		for j := range n {
			if obstacleGrid[i][j] == 1 {
				dp[j] = 0
			} else if j > 0 {
				dp[j] += dp[j-1]
			}
		}
	}
	return dp[n-1]
}

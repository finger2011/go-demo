package main

import "fmt"

func main() {
	// 4
	// matrix := [][]byte{{'1', '0', '1', '0', '0'}, {'1', '0', '1', '1', '1'}, {'1', '1', '1', '1', '1'}, {'1', '0', '0', '1', '0'}}
	// 0
	// matrix := [][]byte{{'0'}}
	// 9
	matrix := [][]byte{
		{'0', '0', '0', '1'},
		{'1', '1', '0', '1'},
		{'1', '1', '1', '1'},
		{'0', '1', '1', '1'},
		{'0', '1', '1', '1'},
	}

	fmt.Println("max:", maximalSquare(matrix))
}

func maximalSquare(matrix [][]byte) int {
	m, n := len(matrix), len(matrix[0])
	dp := make([][]int, m)
	var ans int
	for i := range m {
		dp[i] = make([]int, n)
		for j := range n {
			if matrix[i][j] == '0' {
				dp[i][j] = 0
			} else {
				dp[i][j] = 1
			}
			if i > 0 && j > 0 && dp[i][j] == 1 && dp[i-1][j-1] > 0 {

				maxEdge := dp[i-1][j-1] + 1
				edge := dp[i][j]
				for k := 1; k < maxEdge; k++ {
					if dp[i-k][j] == 0 || dp[i][j-k] == 0 {
						break
					}
					edge++
				}
				dp[i][j] = edge
			}
			ans = max(ans, dp[i][j]*dp[i][j])
		}
	}
	return ans
}

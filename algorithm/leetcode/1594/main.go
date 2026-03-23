package main

import (
	"fmt"
	"math"
)

func main() {
	grid := [][]int{{1, -2, 1}, {1, -2, 1}, {3, -4, 1}}
	fmt.Println("maxProductPath(", grid, ") ==> ", maxProductPath(grid))
}

// dfs[i][j] = max(dfs[i - 1][j] * grid[i][j], dfs[i][j - 1] * grid[i][j]) 注意负数的情况
func maxProductPath(grid [][]int) int {
	const mod = 1000000007
	m, n := len(grid), len(grid[0])
	mem := make([][2]int, n)
	for i := range m {
		for j := range n {
			x := grid[i][j]
			if i == 0 && j == 0 {
				mem[j] = [2]int{x, x}
				continue
			}
			ansMin, ansMax := math.MaxInt, math.MinInt
			if i > 0 {
				mMin, mMax := mem[j][0], mem[j][1]
				ansMin = min(x*mMin, x*mMax)
				ansMax = max(x*mMin, x*mMax)
			}
			if j > 0 {
				mMin, mMax := mem[j-1][0], mem[j-1][1]
				ansMin = min(ansMin, x*mMin, x*mMax)
				ansMax = max(ansMax, x*mMin, x*mMax)
			}
			mem[j] = [2]int{ansMin, ansMax}
		}
	}
	ans := mem[n-1][1]
	if ans < 0 {
		return -1
	}
	return ans % mod
}

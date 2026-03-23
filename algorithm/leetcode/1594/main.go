package main

import (
	"fmt"
	"math"
)

func main() {
	grid := [][]int{{1, -2, 1}, {1, -2, 1}, {3, -4, 1}}
	fmt.Println("maxProductPath(", grid, ") ==> ", maxProductPath(grid))
}

func maxProductPath(grid [][]int) int {
	const mod = 1000000007
	m, n := len(grid), len(grid[0])
	mem := make([][][2]int, m)
	for i := range mem {
		mem[i] = make([][2]int, n)
		for j := range mem[i] {
			mem[i][j] = [2]int{math.MinInt, math.MinInt}
		}
	}
	var dfs func(i, j int) (int, int)
	dfs = func(i, j int) (int, int) {
		x := grid[i][j]
		if i == 0 && j == 0 {
			return x, x
		}
		if mem[i][j][0] != math.MinInt {
			return mem[i][j][0], mem[i][j][1]
		}
		ansMin, ansMax := math.MaxInt, math.MinInt
		if i > 0 {
			mMin, mMax := dfs(i-1, j)
			ansMin = min(x*mMin, x*mMax)
			ansMax = max(x*mMin, x*mMax)
		}
		if j > 0 {
			mMin, mMax := dfs(i, j-1)
			ansMin = min(ansMin, x*mMin, x*mMax)
			ansMax = max(ansMax, x*mMin, x*mMax)
		}
		mem[i][j] = [2]int{ansMin, ansMax}
		return ansMin, ansMax
	}
	_, ans := dfs(m-1, n-1)
	if ans >= 0 {
		return ans % mod
	}
	return -1
}

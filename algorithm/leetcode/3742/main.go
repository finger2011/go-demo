package main

import (
	"fmt"
	"math"
)

func main() {
	grid := [][]int{{0, 1}, {2, 0}}
	k := 1
	fmt.Println("maxPathScore:", maxPathScore(grid, k))
}

func maxPathScore(grid [][]int, k int) int {
	m, n := len(grid), len(grid[0])
	memo := make([][][]int, m)
	for i := range memo {
		memo[i] = make([][]int, n)
		for j := range memo[i] {
			memo[i][j] = make([]int, k+1)
			for p := range memo[i][j] {
				memo[i][j][p] = -1
			}
		}
	}

	var dfs func(int, int, int) int
	dfs = func(i, j, k int) int {
		if i < 0 || j < 0 || k < 0 {
			return math.MinInt
		}
		if i == 0 && j == 0 {
			return 0
		}
		p := &memo[i][j][k]
		if *p != -1 {
			return *p
		}
		x := grid[i][j]
		if x > 0 {
			k--
		}
		res := max(dfs(i-1, j, k), dfs(i, j-1, k)) + x
		*p = res
		return res
	}

	ans := dfs(m-1, n-1, k)
	if ans < 0 {
		return -1
	}
	return ans
}

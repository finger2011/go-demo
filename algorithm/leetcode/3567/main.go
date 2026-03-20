package main

import (
	"fmt"
	"math"
	"slices"
)

func main() {
	grid := [][]int{{1, -2, 3}, {2, 3, 5}}
	k := 2
	fmt.Println("minAbsDiff(", grid, ", ", k, ") ===> ", minAbsDiff(grid, k))
}

func minAbsDiff(grid [][]int, k int) [][]int {
	m, n := len(grid), len(grid[0])
	ans := make([][]int, m-k+1)
	sortArr := make([]int, k*k)
	for i := range ans {
		ans[i] = make([]int, n-k+1)
		for j := range ans[i] {
			a := sortArr[:0]
			for _, row := range grid[i : i+k] {
				a = append(a, row[j:j+k]...)
			}
			slices.Sort(a)
			res := math.MaxInt
			for l := 1; l < len(a); l++ {
				if a[l] > a[l-1] {
					res = min(res, a[l]-a[l-1])
				}
			}
			if res < math.MaxInt {
				ans[i][j] = res
			}
		}

	}
	return ans
}

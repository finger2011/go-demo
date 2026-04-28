package main

import (
	"fmt"
	"slices"
)

func main() {
	grid := [][]int{{931, 128}, {639, 712}}
	x := 73
	fmt.Println("minOperations:", minOperations(grid, x))
}

// 中位数定理：https://zhuanlan.zhihu.com/p/1922938031687595039
func minOperations(grid [][]int, x int) int {
	var ans int
	m, n := len(grid), len(grid[0])
	total := m * n
	pos := make([]int, 0, total)
	for i := range grid {
		for j := range grid[i] {
			if (grid[i][j]-grid[0][0])%x != 0 {
				return -1
			}
		}
		pos = append(pos, grid[i]...)
	}
	slices.Sort(pos)
	media := pos[total/2]
	for _, num := range pos {
		if num >= media {
			ans += num - media
		} else {
			ans += media - num
		}
	}
	return ans / x
}

func minOperations2(grid [][]int, x int) int {
	var ans int
	m, n := len(grid), len(grid[0])
	total := m * n
	pos := make([]int, total)
	for i := range grid {
		for j := range grid[i] {
			if (grid[i][j]-grid[0][0])%x != 0 {
				return -1
			}
			pos[i*n+j] = (grid[i][j] - grid[0][0]) / x
		}
	}
	slices.Sort(pos)
	var sum, p int
	start := pos[0]
	for i := range pos {
		pos[i] -= start
		sum += pos[i]
	}
	ans = sum
	for p < total {
		n := p
		for ; p < total && pos[p] <= pos[n]; p++ {
		}
		var delta int
		if p < total {
			delta = pos[p] - pos[n]
		} else {
			delta = pos[p-1] - pos[n]
		}
		sum += (p*2 - total) * delta
		ans = min(ans, sum)
	}
	return ans
}

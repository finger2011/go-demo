package main

import "fmt"

func main() {
	grid := [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}, {13, 14, 15, 16}}
	k := 2
	fmt.Println("rotate:", rotateGrid(grid, k))
}

func rotateGrid(grid [][]int, k int) [][]int {
	m, n := len(grid), len(grid[0])
	for i := 0; i < m/2 && i < n/2; i++ {
		sum := 2 * (m + n - 2 - 4*i)
		ops := k % sum
		if ops == 0 {
			continue
		}
		var op int
		nums := make([]int, sum)
		for j := i; j < n-i; j++ {
			nums[op] = grid[m-1-i][j]
			op++
		}
		for j := m - 2 - i; j >= i; j-- {
			nums[op] = grid[j][n-1-i]
			op++
		}
		for j := n - 2 - i; j >= i; j-- {
			nums[op] = grid[i][j]
			op++
		}
		for j := i + 1; j < m-1-i; j++ {
			nums[op] = grid[j][i]
			op++
		}
		nums = append(nums[sum-ops:], nums[:sum-ops]...)

		op = 0
		for j := i; j < n-i; j++ {
			grid[m-1-i][j] = nums[op]
			op++
		}
		for j := m - 2 - i; j >= i; j-- {
			grid[j][n-1-i] = nums[op]
			op++
		}
		for j := n - 2 - i; j >= i; j-- {
			grid[i][j] = nums[op]
			op++
		}
		for j := i + 1; j < m-1-i; j++ {
			grid[j][i] = nums[op]
			op++
		}
	}
	return grid
}

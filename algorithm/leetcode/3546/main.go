package main

import "fmt"

func main() {
	grid := [][]int{{1, 4}, {2, 3}}
	fmt.Println("canPartitionGrid(", grid, ") ===> ", canPartitionGrid(grid))
}

func canPartitionGrid(grid [][]int) bool {
	m, n := len(grid), len(grid[0])
	pre := make([][]int, m)
	for i := range m {
		pre[i] = make([]int, n)
		sum := 0
		for j := range n {
			sum += grid[i][j]
			if i == 0 {
				pre[i][j] = sum
			} else {
				pre[i][j] = sum + pre[i-1][j]
			}
		}
	}
	for i := range m {
		if pre[i][n-1]*2 == pre[m-1][n-1] {
			return true
		} else if pre[i][n-1]*2 > pre[m-1][n-1] {
			break
		}
	}
	for j := range n {
		if pre[m-1][j]*2 == pre[m-1][n-1] {
			return true
		} else if pre[m-1][j]*2 > pre[m-1][n-1] {
			break
		}
	}
	return false
}

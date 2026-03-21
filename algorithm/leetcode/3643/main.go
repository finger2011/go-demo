package main

import (
	"fmt"
)

func main() {
	// grid := [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}, {13, 14, 15, 16}}
	// x, y, k := 1, 0, 3
	// grid := [][]int{{3, 4, 2, 3}, {2, 3, 4, 2}}
	// x, y, k := 0, 2, 2
	grid := [][]int{{14, 3, 18, 16}, {2, 14, 11, 20}, {19, 19, 4, 15}, {11, 15, 18, 6}}
	x, y, k := 0, 0, 4
	fmt.Println("reverseSubmatrix(", grid, ", ", x, ", ", y, ", ", k, "):")
	fmt.Println(reverseSubmatrix(grid, x, y, k))
	fmt.Println("reverseSubmatrix2(", grid, ", ", x, ", ", y, ", ", k, "):")
	fmt.Println(reverseSubmatrix2(grid, x, y, k))
}

func reverseSubmatrix(grid [][]int, x int, y int, k int) [][]int {
	n := len(grid[0])
	row1, row2 := make([]int, n), make([]int, n)
	for i := range k / 2 {
		r1, r2 := row1[:0], row2[:0]
		if y > 0 {
			r1 = append(r1, grid[x+i][:y]...)
			r2 = append(r2, grid[x+k-1-i][:y]...)
		}
		r1 = append(r1, grid[x+k-1-i][y:y+k]...)
		r2 = append(r2, grid[x+i][y:y+k]...)
		if y+k < n {
			r1 = append(r1, grid[x+i][y+k:]...)
			r2 = append(r2, grid[x+k-1-i][y+k:]...)
		}
		copy(grid[x+i], r1)
		copy(grid[x+k-1-i], r2)
	}
	return grid
}

func reverseSubmatrix2(grid [][]int, x, y, k int) [][]int {
	l, r := x, x+k-1
	for l < r {
		for j := y; j < y+k; j++ {
			grid[l][j], grid[r][j] = grid[r][j], grid[l][j]
		}
		l++
		r--
	}
	return grid
}

package main

import "fmt"

func main() {
	grid := [][]byte{{'c', 'c', 'c', 'a'}, {'c', 'd', 'c', 'c'}, {'c', 'c', 'e', 'c'}, {'f', 'c', 'c', 'c'}}
	fmt.Println("containsCycle:", grid)
	fmt.Println("result:", containsCycle(grid))
}

var dir = [4][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

func containsCycle(grid [][]byte) bool {
	m, n := len(grid), len(grid[0])
	pos := make([][]bool, m)
	for i := range pos {
		pos[i] = make([]bool, n)
	}
	type point struct {
		x, y int
	}
	var dfs func(point, point) bool
	dfs = func(p1, p2 point) bool {
		pos[p1.x][p1.y] = true
		for _, d := range dir {
			x, y := p1.x+d[0], p1.y+d[1]
			if x < 0 || x >= m || y < 0 || y >= n || grid[p1.x][p1.y] != grid[x][y] || (x == p2.x && y == p2.y) {
				continue
			}
			if pos[x][y] || dfs(point{x, y}, p1) {
				return true
			}
		}
		return false
	}
	for i := range pos {
		for j, b := range pos[i] {
			if (!b) && dfs(point{i, j}, point{-1, 1}) {
				return true
			}
		}
	}

	return false
}

// func modifySlice(slice [][]int) {
// 	slice[0][0] = 999

// 	slice[1] = []int{8, 8, 8}
// 	slice = append(slice, []int{7, 7, 7})
// }

// func main() {
// 	original := [][]int{
// 		{1, 2, 3},
// 		{4, 5, 6},
// 		{7, 8, 9},
// 	}

// 	modifySlice(original)

// 	fmt.Println(original[0])
// 	fmt.Println(original[1])
// 	fmt.Println(len(original))
// }

package main

import "fmt"

type Input struct {
	grid   [][]int
	except bool
}

func main() {
	tests := []*Input{
		{
			grid:   [][]int{{4, 1, 3}, {6, 1, 2}},
			except: true,
		},
		{
			grid:   [][]int{{2}, {2}, {2}, {2}, {2}, {6}},
			except: true,
		},
		{
			grid:   [][]int{{1, 1, 2}},
			except: false,
		},
		{
			grid:   [][]int{{1, 1, 1, 1, 1, 1, 3}},
			except: true,
		},

		{
			grid:   [][]int{{1, 2, 1}, {1, 2, 1}},
			except: false,
		},
		{
			grid:   [][]int{{2, 4, 3}, {6, 5, 2}},
			except: true,
		},
		{
			grid:   [][]int{{1}},
			except: true,
		},
	}
	for _, test := range tests {
		fmt.Println("=========start")
		fmt.Println("grid:", test.grid)
		ans := hasValidPath(test.grid)
		fmt.Println("success:", ans == test.except)
		fmt.Println("=========end")
	}
}

var dirs = map[int][2][2]int{
	1: {{0, 1}, {0, -1}},
	2: {{-1, 0}, {1, 0}},
	3: {{1, 0}, {0, -1}},
	4: {{1, 0}, {0, 1}},
	5: {{-1, 0}, {0, -1}},
	6: {{-1, 0}, {0, 1}},
}

func hasValidPath(grid [][]int) bool {
	m, n := len(grid), len(grid[0])
	if m <= 1 && n <= 1 {
		return true
	}
	var pos [][]bool
	for _, dir := range dirs[grid[0][0]] {
		var x, y, preX, preY, nX, nY int
		x, y = dir[0], dir[1]
		pos = make([][]bool, m)
		for i := range pos {
			pos[i] = make([]bool, n)
		}
		pos[preX][preY] = false
		for x >= 0 && x < m && y >= 0 && y < n {
			if pos[x][y] {
				break
			}
			nx1, ny1 := dirs[grid[x][y]][0][0]+x, dirs[grid[x][y]][0][1]+y
			nx2, ny2 := dirs[grid[x][y]][1][0]+x, dirs[grid[x][y]][1][1]+y
			if nx1 == preX && ny1 == preY {
				nX, nY = nx2, ny2
			} else if nx2 == preX && ny2 == preY {
				nX, nY = nx1, ny1
			} else {
				break
			}
			if x == m-1 && y == n-1 {
				return true
			}
			pos[x][y] = true
			x, y, preX, preY = nX, nY, x, y
		}
	}
	return false
}

package main

import "bytes"

func main() {

}

func rotateTheBox(boxGrid [][]byte) [][]byte {
	m, n := len(boxGrid), len(boxGrid[0])
	grid := make([][]byte, n)
	for i := range grid {
		grid[i] = bytes.Repeat([]byte{'.'}, m)
	}
	for i, row := range boxGrid {
		k := n - 1
		for j := n - 1; j >= 0; j-- {
			if row[j] == '*' { // 障碍物
				grid[j][m-1-i] = '*'
				k = j - 1 // 障碍物左边最近的石头，在旋转后掉落到 j-1
			} else if row[j] == '#' { // 石头
				grid[k][m-1-i] = '#' // 旋转后，石头掉落到 k
				k--
			}
		}
	}
	return grid
}

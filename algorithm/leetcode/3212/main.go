package main

import "fmt"

func main() {
	grid := [][]byte{{'X', 'Y', '.'}, {'Y', '.', '.'}}
	fmt.Println("numberOfSubmatrices(", grid, ") ===> ", numberOfSubmatrices(grid))
}

func numberOfSubmatrices(grid [][]byte) int {
	var ans int
	cols := len(grid[0])
	xCount, yCount := make([]int, cols), make([]int, cols)
	for _, row := range grid {
		var sumX, sumY int
		for j, v := range row {
			switch v {
			case 'X':
				sumX++
			case 'Y':
				sumY++
			default:
			}
			xCount[j] += sumX
			yCount[j] += sumY

			if xCount[j] > 0 && xCount[j] == yCount[j] {
				ans++
			}
		}
	}

	return ans
}

package main

import "fmt"

func main() {
	grid := [][]int{{7, 2, 9}, {1, 5, 0}, {2, 6, 6}}
	k := 20
	fmt.Println("countSubmatrices(", grid, ", ", k, ") ===> ", countSubmatrices(grid, k))
}

func countSubmatrices(grid [][]int, k int) int {
	var ans int
	n := len(grid[0])
	pres := make([]int, n)
	for _, row := range grid {
		var rowSum int
		for i, v := range row {
			if pres[i] > k {
				break
			}
			rowSum += v
			if rowSum+pres[i] <= k {
				ans++
				pres[i] += rowSum
			} else {
				pres[i] += rowSum
				break
			}
		}
	}
	return ans
}

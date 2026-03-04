package main

import "fmt"

func main() {
	mat := [][]int{{1, 0, 0}, {0, 1, 0}, {0, 0, 1}}
	fmt.Println("numSpecial(", mat, ") ====> ", numSpecial(mat))
}

func numSpecial(mat [][]int) int {
	var ans int
	rows, cols := len(mat), len(mat[0])
	rNums, cNums := make([]int, rows), make([]int, cols)
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			rNums[i] += mat[i][j]
			cNums[j] += mat[i][j]
		}
	}
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if mat[i][j] == 1 && rNums[i] == 1 && cNums[j] == 1 {
				ans++
			}
		}
	}
	return ans
}

package main

import (
	"fmt"
	"slices"
)

func main() {
	matrix := [][]int{{0, 0, 1}, {1, 1, 1}, {1, 0, 1}}
	fmt.Println("largestSubmatrix(", matrix, ") ===> ", largestSubmatrix(matrix))
}

// 0 0 1		 1 0  0
// 1 1 1   =>	[1 1] 1
// 1 0 1		[1 1] 0
func largestSubmatrix(matrix [][]int) int {
	var ans int
	n := len(matrix[0])
	height := make([]int, n)
	for _, row := range matrix {
		for j, v := range row {
			if v == 0 {
				height[j] = 0
			} else {
				height[j]++
			}
		}
		hs := slices.Clone(height)
		slices.Sort(hs)
		for i, h := range hs {
			ans = max(ans, (n-i)*h)
		}
	}
	return ans
}

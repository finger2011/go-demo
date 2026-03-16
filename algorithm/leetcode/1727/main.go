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
// 1.枚举矩形的底，即从i,j往上全1是的个数，如果grid[i][j] = 0,那么无法作为矩形的底
// 2. 贪心的把全是1的底放在一起，那么最大矩形就是底长度 * 最小高度
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
		// 底的长度不同，高度也不同
		for i, h := range hs {
			if h == 0 {
				continue
			}
			ans = max(ans, (n-i)*h)
		}
	}
	return ans
}

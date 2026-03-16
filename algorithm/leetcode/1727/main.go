package main

import (
	"fmt"
	"slices"
)

func main() {
	matrix := [][]int{{0, 0, 1}, {1, 1, 1}, {1, 0, 1}}
	fmt.Println("largestSubmatrix (", matrix, ") ===> ", largestSubmatrix(matrix))
	fmt.Println("largestSubmatrix2(", matrix, ") ===> ", largestSubmatrix(matrix))
}

// 0 0 1		 1 0  0
// 1 1 1   =>	[1 1] 1
// 1 0 1		[1 1] 0
// 1.枚举矩形的底，即从i,j往上全1是的个数，如果grid[i][j] = 0,那么无法作为矩形的底
// 2. 贪心的把全是1的底放在一起，那么最大矩形就是底长度 * 最小高度
// 时间复杂度O(mnlogn):m行，对每一行都需要排序nlogn,瓶颈在排序上
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

// 优化排序
// 针对第i行，如果第i-1行已经有序，那么到第i行，只需要把i-1行中为0的排在前面，剩余的内容天然有序
// 时间复杂度为O(mn)
func largestSubmatrix2(matrix [][]int) int {
	var ans int
	n := len(matrix[0])
	height := make([]int, n)
	idx := make([]int, n)
	for i := range idx {
		idx[i] = i
	}
	_nonZero := make([]int, n)
	for _, row := range matrix {
		zeros := idx[:0]
		nonZero := _nonZero[:0]
		// 这里主要需要对idx进行遍历，而不是row
		for _, j := range idx {
			if row[j] == 0 {
				height[j] = 0
				zeros = append(zeros, j)
			} else {
				height[j]++
				nonZero = append(nonZero, j)
			}
		}
		idx = append(zeros, nonZero...)
		for i := len(zeros); i < n; i++ {
			ans = max(ans, (n-i)*height[idx[i]])
		}
	}
	return ans
}

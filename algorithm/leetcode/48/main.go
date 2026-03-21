package main

import (
	"fmt"
	"slices"
)

func main() {
	matrix := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	fmt.Println("before:", matrix)
	rotate(matrix)
	fmt.Println("after:", matrix)
}

// 顺时针旋转90° = 对角线转置 + 行翻转
// 第一行 -> 最后一列 第二行 -> 倒数第二列 ...
// 第一列 -> 第一行 第二列 -> 第二行
// (i, j)  -> (j, n - 1 - i)  = (i, j) -> (j, i) -> (j, n - 1 - i)
// (i, j) -> (j, i) 对应按对角线转置
// (j, i) -> (j, n - 1 - i) 行翻转
// 相同的，如果是逆时针旋转 = 对角线转置 + 列翻转
// (i, j) -> (n - 1 - j, i) = (i, j) -> (j, i) -> (n - 1 - j,i)
func rotate(matrix [][]int) {
	n := len(matrix)
	// 转置
	for i := range n {
		for j := range i {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
	// 行翻转
	for _, row := range matrix {
		slices.Reverse(row)
	}
}

package main

import (
	"fmt"
	"slices"
)

func main() {
	mat := [][]int{{0, 0, 0}, {0, 1, 0}, {1, 1, 1}}
	target := [][]int{{1, 1, 1}, {0, 1, 0}, {0, 0, 0}}

	fmt.Println("mat:", mat)
	fmt.Println("target:", target)
	fmt.Println("findRotation:", findRotation(mat, target))
}

func findRotation(mat [][]int, target [][]int) bool {
	for range 4 {
		if slices.EqualFunc(mat, target, slices.Equal[[]int]) {
			return true
		}
		rotate(mat)
	}
	return false
}

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

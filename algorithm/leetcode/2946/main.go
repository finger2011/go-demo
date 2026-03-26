package main

import (
	"fmt"
)

func main() {
	mat := [][]int{{1, 2, 1, 2}, {5, 5, 5, 5}, {6, 3, 6, 3}}
	k := 2
	fmt.Println("areSimilar:")
	fmt.Println("mat:", mat)
	fmt.Println("k:", k)
	fmt.Println("result:", areSimilar(mat, k))
}

// 通过左移k次，是否相等
// 对于右移，逆过程就是左移k次
// 所以只需要一致判断左移k次
func areSimilar(mat [][]int, k int) bool {
	n := len(mat[0])
	k %= n
	if k == 0 {
		return true
	}
	for _, row := range mat {
		for j, num := range row {
			if num != row[(j+k)%n] {
				return false
			}
		}
	}

	return true
}

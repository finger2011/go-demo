package main

import "fmt"

func main() {
	grid := [][]int{{1, 2}, {3, 4}}
	fmt.Println("constructProductMatrix:")
	fmt.Println("src:", grid)
	fmt.Println("after:", constructProductMatrix(grid))
}
func constructProductMatrix(grid [][]int) [][]int {
	const mod = 12345
	m, n := len(grid), len(grid[0])
	ans := make([][]int, m)
	suf := 1
	for i := m - 1; i >= 0; i-- {
		ans[i] = make([]int, n)
		for j := n - 1; j >= 0; j-- {
			ans[i][j] = suf
			suf = suf * grid[i][j] % mod
		}
	}
	pre := 1
	for i := range grid {
		for j := range grid[i] {
			ans[i][j] = ans[i][j] * pre % mod
			pre = pre * grid[i][j] % mod
		}
	}
	return ans
}

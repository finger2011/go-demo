package main

import "fmt"

func main() {
	grid := [][]int{
		{0, 0, 1},
		{1, 1, 0},
		{1, 0, 0},
	}
	fmt.Println("minSwaps(", grid, ") ===> ", minSwaps(grid))
}

// 贪心：交换最近符合条件的行，其他符合条件的，必然也符合next循环的条件
func minSwaps(grid [][]int) int {
	n := len(grid)
	pos := make([]int, n)
	for i := 0; i < n; i++ {
		pos[i] = -1
		for j := n - 1; j >= 0; j-- {
			if grid[i][j] == 1 {
				pos[i] = j
				break
			}
		}
	}
	var ans int
	for i := 0; i < n; i++ {
		rev := -1
		for j := i; j < n; j++ {
			if pos[j] <= i {
				ans += j - i
				rev = j
				break
			}
		}
		if rev == -1 {
			return -1
		} else {
			for j := rev; j > i; j-- {
				// 模拟交换
				pos[j], pos[j-1] = pos[j-1], pos[j]
			}
		}
	}
	return ans
}

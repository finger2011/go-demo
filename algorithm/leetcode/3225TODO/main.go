package main

import "fmt"

func main() {
	grid := [][]int{{10, 9, 0, 0, 15}, {7, 1, 0, 8, 0}, {5, 20, 0, 11, 0}, {0, 0, 0, 1, 2}, {8, 12, 1, 10, 3}}
	fmt.Println("score:", maximumScore(grid))
}

func maximumScore(grid [][]int) (ans int64) {
	n := len(grid)
	colSum := make([][]int64, n)
	for j := range colSum {
		colSum[j] = make([]int64, n+1)
		for i, row := range grid {
			colSum[j][i+1] = colSum[j][i] + int64(row[j])
		}
	}

	f := make([][][2]int64, n)
	for j := range f {
		f[j] = make([][2]int64, n+1)
	}
	for j := 0; j < n-1; j++ {
		// 用前缀最大值优化
		preMax := f[j][0][1] - colSum[j][0]
		for pre := 1; pre <= n; pre++ {
			f[j+1][pre][0] = max(f[j][pre][0], preMax+colSum[j][pre])
			f[j+1][pre][1] = f[j+1][pre][0]
			preMax = max(preMax, f[j][pre][1]-colSum[j][pre])
		}

		// 用后缀最大值优化
		sufMax := f[j][n][0] + colSum[j+1][n]
		for pre := n - 1; pre > 0; pre-- {
			f[j+1][pre][0] = max(f[j+1][pre][0], sufMax-colSum[j+1][pre])
			sufMax = max(sufMax, f[j][pre][0]+colSum[j+1][pre])
		}

		// 单独计算 pre=0 的状态
		f[j+1][0][0] = sufMax                      // 无需考虑 f[j][0][0]，因为不能连续三列全白
		f[j+1][0][1] = max(f[j][0][0], f[j][n][0]) // 第 j 列要么全白，要么全黑
	}

	for _, row := range f[n-1] {
		ans = max(ans, row[0])
	}
	return ans
}

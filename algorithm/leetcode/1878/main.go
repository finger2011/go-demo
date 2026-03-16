package main

import (
	"container/heap"
	"fmt"
)

func main() {
	// grid := [][]int{{3, 4, 5, 1, 3}, {3, 3, 4, 2, 3}, {20, 30, 200, 40, 10}, {1, 5, 5, 4, 1}, {4, 3, 2, 2, 5}}
	grid := [][]int{{20, 17, 9, 13, 5, 2, 9, 1, 5}, {14, 9, 9, 9, 16, 18, 3, 4, 12}, {18, 15, 10, 20, 19, 20, 15, 12, 11}, {19, 16, 19, 18, 8, 13, 15, 14, 11}, {4, 19, 5, 2, 19, 17, 7, 2, 2}}
	fmt.Println("getBiggestThree (", grid, ") ===> ", getBiggestThree(grid))
	fmt.Println("getBiggestThree2(", grid, ") ===> ", getBiggestThree2(grid))
}

func getBiggestThree2(grid [][]int) []int {
	var x, y, z int
	m, n := len(grid), len(grid[0])
	sum1, sum2 := make([][]int, m+1), make([][]int, m+1)
	for i := range sum1 {
		sum1[i] = make([]int, n+1)
		sum2[i] = make([]int, n+1)
	}
	for i, row := range grid {
		for j, v := range row {
			sum1[i+1][j+1] = sum1[i][j] + v
			sum2[i+1][j] = sum2[i][j+1] + v
		}
	}
	querySum1 := func(x, y, k int) int {
		return sum1[x+k][y+k] - sum1[x][y]
	}
	querySum2 := func(x, y, k int) int {
		return sum2[x+k][y+1-k] - sum2[x][y+1]
	}
	update := func(k int) {
		if k > x {
			x, y, z = k, x, y
		} else if k < x && k > y {
			y, z = k, y
		} else if k < y && k > z {
			z = k
		}
	}
	for i, row := range grid {
		for j, v := range row {
			update(v)
			maxLen := min(i, j, m-i-1, n-j-1)
			for k := 1; k <= maxLen; k++ {
				update(querySum1(i-k, j, k) + querySum1(i, j-k, k) + querySum2(i-k+1, j-1, k-1) + querySum2(i, j+k, k+1))
			}
		}
	}
	ans := []int{x, y, z}
	for ans[len(ans)-1] == 0 {
		ans = ans[:len(ans)-1]
	}
	return ans
}

// 以i，j为中心构建的菱形个数,暴力求解
func getBiggestThree(grid [][]int) []int {
	var ans []int
	if len(grid) == 0 {
		return ans
	}
	h := &MaxHeap{}
	heap.Init(h)
	m, n := len(grid), len(grid[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			fmt.Println("center(", i, ", ", j, ") start:")
			// 面积为0
			heap.Push(h, grid[i][j])
			maxLen := min(i, m-i-1, j, n-j-1)
			fmt.Println("maxLen:", maxLen)
			for k := 1; k <= maxLen; k++ {
				var area int
				for l := k; l >= 0; l-- {
					area += grid[i-l][j-(k-l)]
					area += grid[i+l][j+(k-l)]
					if l != k && l != 0 {
						area += grid[i-l][j+(k-l)]
						area += grid[i+l][j-(k-l)]
					}
				}
				fmt.Println("len:", k, ", area:", area)
				heap.Push(h, area)
			}
			fmt.Println("center(", i, ", ", j, ") end")
		}
	}
	fmt.Println("heap:", h)
	for {
		if h.Len() == 0 {
			break
		}
		area := heap.Pop(h).(int)
		if len(ans) == 0 || area != ans[len(ans)-1] {
			ans = append(ans, area)
		}
		if len(ans) == 3 {
			break
		}
	}
	return ans
}

// MaxHeap 是一个基于整数切片的大根堆类型
type MaxHeap []int

// 实现 sort.Interface 的三个方法
func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i] > h[j] } // 关键点：大于号，确保大根堆
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

// 实现 heap.Interface 的 Push 方法
func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

// 实现 heap.Interface 的 Pop 方法
func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

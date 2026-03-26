package main

import (
	"fmt"
	"slices"
)

type MyTest struct {
	grid   [][]int
	except bool
}

func main() {
	tests := []MyTest{
		{
			grid:   [][]int{{1, 2, 4}, {2, 3, 5}},
			except: false,
		},
		{
			grid:   [][]int{{10, 5, 4, 5}},
			except: false,
		},
		{
			grid:   [][]int{{10}, {5}, {4}, {5}},
			except: false,
		},
		{
			grid:   [][]int{{5, 5, 6, 2, 2, 2}},
			except: true,
		},
		{
			grid:   [][]int{{1}, {2}, {1}},
			except: true,
		},
		{
			grid:   [][]int{{1, 2, 1}},
			except: true,
		},
		{
			grid:   [][]int{{1, 2, 4}, {1, 6, 6}, {5, 6, 7}},
			except: true,
		},
		{
			grid:   [][]int{{100000}, {100000}, {100000}, {100000}, {1}},
			except: true,
		},
	}
	for i, test := range tests {
		fmt.Println("test[", i, "]:")
		fmt.Println("grid:", test.grid)
		if canPartitionGrid2(test.grid) == test.except {
			fmt.Println("============================success")
		} else {
			fmt.Println("============================failed")
		}
	}
}

func canPartitionGrid2(grid [][]int) bool {
	var total int
	for _, row := range grid {
		for _, n := range row {
			total += n
		}
	}
	check := func(a [][]int) bool {
		m, n := len(a), len(a[0])
		var delta int
		checkPre := func() bool {
			hasNum := map[int]bool{0: true}
			sum := 0
			for i, row := range a[:m-1] {
				for j, num := range row {
					sum += num
					if i > 0 || j == 0 || j == n-1 {
						hasNum[num] = true
					}
				}
				delta = 2*sum - total
				if n == 1 {
					if delta == 0 || delta == a[0][0] || delta == row[0] {
						return true
					}
					continue
				}
				if hasNum[delta] {
					return true
				}
				if i == 0 {
					for _, num := range row {
						hasNum[num] = true
					}
				}
			}
			return false
		}
		if checkPre() {
			return true
		}
		slices.Reverse(a)
		return checkPre()
	}
	return check(grid) || check(rotate(grid))
}

// 顺时针旋转矩阵 90°
func rotate(a [][]int) [][]int {
	m, n := len(a), len(a[0])
	b := make([][]int, n)
	for i := range b {
		b[i] = make([]int, m)
	}
	for i, row := range a {
		for j, x := range row {
			b[j][m-1-i] = x
		}
	}
	return b
}

func canPartitionGrid(grid [][]int) bool {
	m, n := len(grid), len(grid[0])
	pre := make([][]int, m)
	numMap := map[int][][2]int{}
	for i := range m {
		pre[i] = make([]int, n)
		sum := 0
		for j := range n {
			numMap[grid[i][j]] = append(numMap[grid[i][j]], [2]int{i, j})
			sum += grid[i][j]
			if i == 0 {
				pre[i][j] = sum
			} else {
				pre[i][j] = sum + pre[i-1][j]
			}
		}
	}
	var isPre bool
	var delta int
	for i := range m {
		if pre[i][n-1]*2 == pre[m-1][n-1] {
			return true
		}
		if pre[i][n-1]*2 > pre[m-1][n-1] {
			isPre = true
			delta = pre[i][n-1]*2 - pre[m-1][n-1]
		} else {
			isPre = false
			delta = pre[m-1][n-1] - pre[i][n-1]*2
		}
		if nums, ok := numMap[delta]; ok {
			for _, ops := range nums {
				if isPre && ops[0] <= i {
					if n == 1 && ops[0] != 0 && ops[0] != i {
						continue
					}
					if i == 0 && ops[1] != 0 && ops[1] != n-1 {
						continue
					}
					return true
				}
				if (!isPre) && ops[0] > i {
					if n == 1 && ops[0] != i+1 && ops[0] != m-1 {
						continue
					}
					if i == m-2 && ops[1] != 0 && ops[1] != n-1 {
						continue
					}
					return true
				}
			}
		}
	}
	for j := range n {
		if pre[m-1][j]*2 == pre[m-1][n-1] {
			return true
		}
		if pre[m-1][j]*2 > pre[m-1][n-1] {
			isPre = true
			delta = pre[m-1][j]*2 - pre[m-1][n-1]
		} else {
			isPre = true
			delta = pre[m-1][n-1] - pre[m-1][j]*2
		}
		if nums, ok := numMap[delta]; ok {
			for _, ops := range nums {
				if isPre && ops[1] <= j {
					if m == 1 && ops[1] != 0 && ops[1] != j {
						continue
					}
					if j == 0 && ops[0] != 0 && ops[0] != m-1 {
						continue
					}
					return true
				}
				if (!isPre) && ops[1] > j {
					if m == 1 && ops[1] != j+1 && ops[1] != n-1 {
						continue
					}
					if j == n-2 && ops[0] != 0 && ops[0] != m-1 {
						continue
					}
					return true
				}
			}
		}
	}
	return false
}

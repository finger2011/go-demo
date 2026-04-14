package main

import (
	"fmt"
	"math"
	"slices"
)

type Input struct {
	robot   []int
	factory [][]int
	except  int64
}

func main() {
	tests := []*Input{
		{
			robot:   []int{9, 11, 99, 101},
			factory: [][]int{{10, 1}, {7, 1}, {14, 1}, {100, 1}, {96, 1}, {103, 1}},
			except:  6,
		},
		{
			robot:   []int{1, -1},
			factory: [][]int{{-2, 1}, {2, 1}},
			except:  2,
		},
		{
			robot:   []int{0, 4, 6},
			factory: [][]int{{2, 2}, {6, 2}},
			except:  4,
		},
	}
	for _, test := range tests {
		fmt.Println("======================start")
		fmt.Println("robot:", test.robot)
		fmt.Println("factory:", test.factory)
		ans := minimumTotalDistance3(test.robot, test.factory)
		fmt.Println("return:", ans)
		fmt.Println("result:", ans == test.except)
		fmt.Println("======================end")
	}
}

// 超时，大量重复计算
func minimumTotalDistance(robot []int, factory [][]int) int64 {
	n, m := len(robot), len(factory)+1
	slices.Sort(robot)
	fac := append([][]int{{math.MinInt32, 0}}, factory...)
	slices.SortFunc(fac, func(x, y []int) int {
		return x[0] - y[0]
	})
	var dfs func(i, j int) int64
	memo := make([][]int64, n)
	for i := range n {
		memo[i] = []int64{math.MaxInt64, math.MaxInt64}
	}
	dfs = func(i, j int) int64 {
		for ; j >= 0; j-- {
			if fac[j][0] < robot[i] {
				break
			}
		}
		for k := j; k < m; k++ {
			if fac[k][1] > 0 && fac[k][0] >= robot[i] {
				fac[k][1]--
				right := int64(fac[k][0] - robot[i])
				if i == 0 {
					memo[i][1] = right
				} else {
					memo[i][1] = dfs(i-1, j) + right
				}
				fac[k][1]++
				break
			}
		}
		for k := j; k >= 0; k-- {
			if fac[k][1] > 0 && fac[k][0] <= robot[i] {
				fac[k][1]--
				left := int64(robot[i] - fac[k][0])
				if i == 0 {
					memo[i][0] = left
				} else {
					memo[i][0] = dfs(i-1, j) + left
				}
				fac[k][1]++
				break
			}
		}
		return min(memo[i][0], memo[i][1])
	}
	ans := dfs(n-1, m-1)
	return ans
}

// n*m*m + nlogn + mlogm
func minimumTotalDistance2(robot []int, factory [][]int) int64 {
	n, m := len(robot), len(factory)
	slices.Sort(robot)
	slices.SortFunc(factory, func(x, y []int) int {
		return x[0] - y[0]
	})
	var dfs func(i, j int) int
	memo := make([][]int, m)
	for i := range m {
		memo[i] = make([]int, n)
		for j := range n {
			memo[i][j] = -1
		}
	}
	dfs = func(i, j int) (ans int) {
		if j < 0 {
			return 0
		}
		if i < 0 {
			return math.MaxInt / 2
		}
		p := &memo[i][j]
		if *p != -1 {
			return *p
		}
		defer func() {
			*p = ans
		}()
		ans = dfs(i-1, j)
		pos, num := factory[i][0], factory[i][1]
		sum := 0
		for k := 1; k <= min(j+1, num); k++ {
			sum += abs(robot[j-k+1] - pos)
			ans = min(ans, dfs(i-1, j-k)+sum)
		}
		return
	}
	return int64(dfs(m-1, n-1))
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func minimumTotalDistance3(robot []int, factory [][]int) int64 {
	n, m := len(robot), len(factory)
	slices.Sort(robot)
	slices.SortFunc(factory, func(x, y []int) int {
		return x[0] - y[0]
	})
	memo := make([][]int, m+1)
	for i := range memo {
		memo[i] = make([]int, n+1)
	}
	for j := range n {
		memo[0][j+1] = math.MaxInt / 2
	}
	for i, fac := range factory {
		pos, num := fac[0], fac[1]
		for j := range n {
			tmp := memo[i][j+1]
			sum := 0
			for k := j; k >= max(j-num+1, 0); k-- {
				sum += abs(robot[k] - pos)
				tmp = min(tmp, memo[i][k]+sum)
			}
			memo[i+1][j+1] = tmp
		}
	}

	return int64(memo[m][n])
}

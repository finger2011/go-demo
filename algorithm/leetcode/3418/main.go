package main

import (
	"fmt"
	"math"
)

type Input struct {
	coins  [][]int
	except int
}

func main() {
	tests := []*Input{
		{
			coins:  [][]int{{-3, -10, 11, -16}, {-13, 19, -2, 2}, {-16, -11, 5, 13}, {-5, 13, -20, -6}},
			except: 32,
		},
		{
			coins:  [][]int{{-7, 12, 12, 13}, {-6, 19, 19, -6}, {9, -2, -10, 16}, {-4, 14, -10, -9}},
			except: 60,
		},
		{
			coins:  [][]int{{-16, 8, -7, -19}, {6, 3, -10, 13}, {13, 15, 4, -3}, {-16, 4, 19, -12}},
			except: 57,
		},
		{
			coins:  [][]int{{10, 10, 10}, {10, 10, 10}},
			except: 40,
		},
		{
			coins:  [][]int{{0, 1, -1}, {1, -2, 3}, {2, -3, 4}},
			except: 8,
		},
	}
	for _, test := range tests {
		fmt.Println("start")
		fmt.Println("coins:", test.coins)
		fmt.Println("except:", test.except)
		ans := maximumAmount(test.coins)
		fmt.Println("return:", ans)
		fmt.Println("result =====> ", test.except == ans)
		fmt.Println("end")
	}
}

func maximumAmount(coins [][]int) int {
	n, m := len(coins), len(coins[0])
	mems := make([][][3]int, n)
	for i := range mems {
		mems[i] = make([][3]int, m)
		for j := range mems[i] {
			for k := range mems[i][j] {
				mems[i][j][k] = math.MinInt
			}
		}

	}
	var dfs func(i, j, k int) int
	dfs = func(i, j, k int) int {
		if i < 0 || j < 0 {
			return math.MinInt
		}
		coin := coins[i][j]
		if i == 0 && j == 0 {
			if k == 0 {
				return coin
			}
			return max(coin, 0)
		}
		mem := &mems[i][j][k]
		if *mem != math.MinInt {
			return *mem
		}
		res := max(dfs(i-1, j, k), dfs(i, j-1, k)) + coin
		if coin < 0 && k > 0 {
			res = max(res, dfs(i-1, j, k-1), dfs(i, j-1, k-1))
		}
		*mem = res
		return res
	}
	return dfs(n-1, m-1, 2)
}

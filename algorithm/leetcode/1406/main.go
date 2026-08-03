package main

import "math"

func main() {

}

func stoneGameIII(stoneValue []int) string {
	n := len(stoneValue)
	memo := make([]int, n)
	for i := range memo {
		memo[i] = math.MinInt // MinInt 表示该状态没有计算过
	}

	var dfs func(int) int
	dfs = func(i int) int {
		if i == n {
			return 0
		}

		p := &memo[i]
		if *p != math.MinInt { // 之前计算过
			// 之前计算过
			return *p
		}

		res := math.MinInt
		sum := 0
		for j := i; j < min(i+3, n); j++ {
			sum += stoneValue[j]
			res = max(res, sum-dfs(j+1))
		}
		*p = res // 记忆化
		return res
	}

	diff := dfs(0)
	if diff == 0 {
		return "Tie"
	}
	if diff > 0 {
		return "Alice"
	}
	return "Bob"
}

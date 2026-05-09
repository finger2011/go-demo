package main

import "math"

func main() {

}

func maximumJumps(nums []int, target int) int {
	n := len(nums)
	memo := make([]int, n)
	var dfs func(t int) int
	dfs = func(t int) int {
		if t == 0 {
			return 0
		}
		p := &memo[t]
		if *p != 0 {
			return *p
		}
		ans := math.MinInt
		for i, num := range nums[:t] {
			if abs(num-nums[t]) <= target {
				ans = max(ans, dfs(i)+1)
			}
		}
		*p = ans
		return ans
	}
	res := dfs(n - 1)
	if res < 0 {
		return -1
	}
	return res
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

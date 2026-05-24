package main

import "fmt"

func main() {
	arr := []int{6, 4, 14, 6, 8, 13, 9, 7, 10, 6, 12}
	d := 2
	fmt.Println("max:", maxJumps(arr, d))
}

func maxJumps(arr []int, d int) int {
	n := len(arr)
	if n == 1 {
		return 1
	}
	memo := make([]int, n)
	var ans int
	var dfs func(op int) int
	dfs = func(op int) int {
		if memo[op] > 0 {
			return memo[op]
		}
		var ant int
		for i := op - 1; i >= op-d && i >= 0 && arr[i] < arr[op]; i-- {
			ant = max(ant, dfs(i))
		}
		for i := op + 1; i <= op+d && i < n && arr[i] < arr[op]; i++ {
			ant = max(ant, dfs(i))
		}
		ant++
		memo[op] = ant
		return ant
	}
	for i := range memo {
		ans = max(ans, dfs(i))
	}
	return ans
}

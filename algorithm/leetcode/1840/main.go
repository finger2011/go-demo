package main

import (
	"fmt"
	"slices"
)

func main() {
	n := 10
	restrictions := [][]int{{5, 3}, {2, 5}, {7, 4}, {10, 3}}
	fmt.Println("max :", maxBuilding(n, restrictions))
}

func maxBuilding(n int, restrictions [][]int) int {
	if len(restrictions) == 0 {
		return n - 1
	}
	var ans int
	memo := make([]int, n-1)
	for i := range n - 1 {
		memo[i] = 1
	}
	slices.SortFunc(restrictions, func(x, y []int) int {
		return x[0] - y[0]
	})
	var preId, preHeight int
	for _, res := range restrictions {
		height := preHeight + (res[0] - 1 - preId) // max height
		k := height - res[1]
		if k <= 0 {
			continue
		}
		preId, preHeight = res[0]-1, res[1]
		i := res[0] - 2
		for ; i >= 0 && k > 0; i-- {
			if memo[i] == 0 || (memo[i] == 1 && k == 1) {
				memo[i]--
				k--
			} else if memo[i] == 1 {
				memo[i] = -1
				k -= 2
			}
		}

		if i < 0 {
			return -1
		}
	}
	var height int
	for i := range n - 1 {
		height += memo[i]
		if height < 0 {
			return -1
		}
		ans = max(ans, height)
	}

	return ans
}

package main

import (
	"fmt"
	"slices"
)

func main() {
	cost := []int{6, 5, 7, 9, 2, 2}
	fmt.Println("min:", minimumCost(cost))
}

func minimumCost(cost []int) int {
	slices.SortFunc(cost, func(x, y int) int {
		return y - x
	})
	var ans int
	length := len(cost)
	for i := 0; i < length; i += 3 {
		ans += cost[i]
		if i+1 < length {
			ans += cost[i+1]
		}
	}
	return ans
}

package main

import (
	"fmt"
	"math"
	"slices"
)

func main() {
	pairs := [][]int{{1, 2}, {7, 8}, {4, 5}}
	fmt.Println("ans:", findLongestChain(pairs))
}

// pairs 按照第二位进行排序
// f[i] =
func findLongestChain(pairs [][]int) int {
	var ans int
	slices.SortFunc(pairs, func(i, j []int) int { return i[1] - j[1] })
	pre := math.MinInt
	for _, pair := range pairs {
		if pair[0] > pre {
			ans++
			pre = pair[1]
		}
	}
	return ans
}

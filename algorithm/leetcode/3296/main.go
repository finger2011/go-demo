package main

import (
	"fmt"
	"math"
	"slices"
	"sort"
)

func main() {
	h := 10
	workerTimes := []int{3, 2, 2, 4}
	fmt.Println("minNumberOfSeconds(", h, ",", workerTimes, ") ===> ", minNumberOfSeconds(h, workerTimes))
}

// 推算公式花费m秒，可以降低的山高度
// t + t * 2 + t * 3 + ... + t * h <= m
// t * (h * (h + 1)/2) <= m
// h*(h+1)/2 <= m/t = k
// h <= (-1 + sqrt(8k+1)) / 2
func minNumberOfSeconds(mountainHeight int, workerTimes []int) int64 {
	maxTime := slices.Max(workerTimes)
	maxT := (mountainHeight-1)/len(workerTimes) + 1
	ans := 1 + sort.Search(maxTime*maxT*(maxT+1)/2-1, func(m int) bool {
		m++
		h := mountainHeight
		for _, t := range workerTimes {
			h -= (int(math.Sqrt(float64(m/t*8+1))) - 1) / 2
			if h <= 0 {
				return true
			}
		}
		return false
	})
	return int64(ans)
}

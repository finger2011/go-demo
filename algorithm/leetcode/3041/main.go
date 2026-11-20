package main

import (
	"fmt"
	"slices"
)

func main() {
	// nums := []int{2, 1, 5, 1, 1, 6, 7, 8}
	// nums := []int{8, 10, 6, 12, 9, 12, 2, 3, 13, 19, 11, 18, 10, 16} // 8
	nums := []int{12, 11, 8, 7, 2, 10, 18, 12} // 6
	fmt.Println("maxSelectedElements:", maxSelectedElements(nums))
}

func maxSelectedElements(nums []int) int {
	var ans int
	slices.Sort(nums)
	f := map[int]int{}
	for _, x := range nums {
		f[x+1] = f[x] + 1
		f[x] = f[x-1] + 1
		ans = max(ans, f[x], f[x+1])
	}
	return ans
}

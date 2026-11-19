package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{1, 4, 3}
	fmt.Println("minimizeSum:", minimizeSum(nums))
}

func minimizeSum(nums []int) int {
	n := len(nums)
	sort.Ints(nums)
	a := nums[n-1] - nums[2]
	b := nums[n-2] - nums[1]
	c := nums[n-3] - nums[0]
	return min(a, min(b, c))
}

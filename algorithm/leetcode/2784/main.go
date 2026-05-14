package main

import (
	"fmt"
	"slices"
)

func main() {
	nums := []int{1, 3, 3, 2}
	fmt.Println("isGood:", isGood(nums))
}

func isGood(nums []int) bool {
	n := len(nums)
	if n <= 1 {
		return false
	}
	slices.Sort(nums)
	if nums[n-1] != n-1 {
		return false
	}
	for i := range n - 1 {
		if nums[i] != i+1 {
			return false
		}
	}
	return true
}

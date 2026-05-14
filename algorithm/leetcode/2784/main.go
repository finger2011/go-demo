package main

import (
	"fmt"
	"slices"
)

func main() {
	nums := []int{1, 3, 3, 2}
	fmt.Println("isGood:", isGood(nums))
}

// 时间复杂度:O(nlogn)
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

// 时间复杂度O(n)
func isGood2(nums []int) bool {
	n := len(nums)
	numMap := make(map[int]int, n)
	for _, num := range nums {
		if num > n-1 || (num == n-1 && numMap[num] > 1) || (num < n-1 && numMap[num] > 0) {
			return false
		}
		numMap[num]++
	}
	return true
}

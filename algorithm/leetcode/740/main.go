package main

import "slices"

func main() {

}

func deleteAndEarn(nums []int) int {
	memo := make([]int, slices.Max(nums)+1)
	for _, num := range nums {
		memo[num] += num
	}
	return rob3(memo)
}

func rob3(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return nums[0]
	}
	f0 := nums[0]
	f1 := max(nums[0], nums[1])
	for i := 2; i < n; i++ {
		f0, f1 = f1, max(f1, f0+nums[i])
	}
	return f1
}

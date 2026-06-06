package main

import "fmt"

func main() {
	nums := []int{1}
	fmt.Println("ans:", leftRightDifference(nums))
}

func leftRightDifference(nums []int) []int {
	n := len(nums)
	pre := make([]int, n)
	suff := make([]int, n)
	pre[0] = 0
	for i := range n - 1 {
		pre[i+1] = pre[i] + nums[i]
	}
	suff[n-1] = 0
	for i := n - 1; i > 0; i-- {
		suff[i-1] = suff[i] + nums[i]
	}
	ans := make([]int, n)
	for i := range n {
		ans[i] = abs(pre[i] - suff[i])
	}
	return ans
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

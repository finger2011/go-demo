package main

import (
	"fmt"
	"math"
	"slices"
)

func main() {
	nums := []int{2, 2, 3}
	fmt.Println("min:", minOperations(nums))
}

// prefix[i] 将前i个数字变为严格升序需要次数
// suffix[i] 将后i个数字变为严格降低序需要次数
// f[i] 从i变为山脉数组需要的次数 = prefix[i-1] + suffix[i+1]
func minOperations(nums []int) int {
	n := len(nums)
	if n < 3 {
		return -1
	}
	maxNum := slices.Max(nums) + n - 1
	var dp func() [][]int
	dp = func() [][]int {
		f := make([][]int, n)
		for i := range f {
			f[i] = make([]int, maxNum)
		}
		for i := 1; i < maxNum; i++ {
			f[0][i] = abs(nums[0] - i)
		}
		for i := 1; i < n; i++ {
			preMin := math.MaxInt
			for j := i + 1; j < maxNum; j++ {
				preMin = min(preMin, f[i-1][j-1])
				f[i][j] = preMin + abs(nums[i]-j)
			}
		}
		return f
	}
	prefix := dp()
	slices.Reverse(nums)
	suffix := dp()
	slices.Reverse(suffix)
	slices.Reverse(nums)
	ans := math.MaxInt
	for i := 1; i < n-1; i++ {
		for j := max(i+1, n-i); j < maxNum; j++ {
			ans = min(ans, prefix[i][j]+suffix[i][j]-abs(nums[i]-j))
		}
	}
	return ans
}

func abs(x int) int {
	if x > 0 {
		return x
	}
	return -x
}

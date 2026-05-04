package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 1}
	fmt.Println("rob:", rob(nums))
}

// memo[i][0] 第i-1不偷
// memo[i][1] 第i-1偷
func rob(nums []int) int {
	n := len(nums)
	memo := make([][2]int, n+1)
	for i, num := range nums {
		if i == 0 {
			memo[i+1] = [2]int{0, num}
			continue
		}
		memo[i+1] = [2]int{0, 0}
		memo[i+1][0] = max(memo[i][0], memo[i][1])
		memo[i+1][1] = max(memo[i][0], memo[i-1][0], memo[i-1][1]) + num
	}
	return max(memo[n][0], memo[n][1])
}

// memo退化为一维数组
func rob2(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return nums[0]
	}
	memo := make([]int, n)
	memo[0] = nums[0]
	memo[1] = max(nums[0], nums[1])
	for i := 2; i < n; i++ {
		memo[i] = max(nums[i]+memo[i-2], memo[i-1])
	}
	return memo[n-1]
}

// memo退化为常量

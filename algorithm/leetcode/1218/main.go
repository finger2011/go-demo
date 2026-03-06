package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4}
	d := 1
	fmt.Println("longestSubsequence(", arr, ", ", d, ") ===> ", longestSubsequence(arr, d))
}

// 动态规划，dp[i] = dp[j] + 1 => dp[v] = dp[v - d] + 1
func longestSubsequence(arr []int, d int) int {
	var ans int
	dp := make(map[int]int, len(arr))
	for _, v := range arr {
		dp[v] = dp[v-d] + 1
		if dp[v] > ans {
			ans = dp[v]
		}
	}
	return ans
}

package main

import (
	"fmt"
)
// https://leetcode-cn.com/problems/delete-operation-for-two-strings/
// leetcode 583

func testMinDistance() {
	// var a, b = "aef", "aef"
	var a, b = "ab", "a"
	fmt.Printf("step:%d\n", minDistance(a, b))
}

func minDistance(word1 string, word2 string) int {
	if len(word1) == 0 || len(word2) == 0 {
		return len(word1) + len(word2)
	}
	var dp = make([][]int, len(word1))
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, len(word2))
		if i == 0 {
			for j := 0; j < len(dp[i]); j++ {
				if j == 0 {
					if word1[i] == word2[j] {
						dp[i][j] = 0
					} else {
						dp[i][j] = 2
					}
				} else {
					if word1[i] == word2[j] {
						dp[i][j] = j
					} else {
						dp[i][j] = min(j + 1, dp[i][j - 1]) + 1
					}
				}
			}
		} else {
			if word1[i] == word2[0] {
				dp[i][0] = i
			} else {
				dp[i][0] = min(i + 1, dp[i - 1][0]) + 1
			}
		}
	}
	for i := 1; i < len(word1); i++ {
		for j := 1; j < len(word2); j++ {
			if word1[i] == word2[j] {
				dp[i][j] = dp[i - 1][j - 1]
			} else {
				dp[i][j] = min(dp[i - 1][j], dp[i][j - 1]) + 1
			}
		}
	}
	fmt.Printf("dp:%v\n", dp)
	return dp[len(word1) - 1][len(word2) - 1]
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

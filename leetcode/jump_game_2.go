package main

import (
	"fmt"
)

//https://leetcode-cn.com/problems/jump-game-ii/
// leetcode 45

func testJump() int {
	nums := []int{2, 1, 1, 1, 4}
	return jump(nums)
}

func jump(nums []int) int {
	if len(nums) <= 1 {
		return 0
	}
	var step, maxLen, i int
	for {
		var tmpMax int
		step++
		for ; i <= maxLen; i++ {
			if tmpMax < i+nums[i] {
				tmpMax = i + nums[i]
			}
		}
		if tmpMax > maxLen {
			maxLen = tmpMax
		}
		fmt.Printf("step %d, max jump:%d\n", step, maxLen)
		if maxLen >= len(nums)-1 {
			break
		}
	}

	return step
}

package main

import (
	"fmt"
	"math"
)

type Input struct {
	nums                  []int
	target, start, except int
}

func main() {
	tests := []*Input{
		{
			nums:   []int{1, 2, 3, 4, 5},
			target: 5,
			start:  3,
			except: 1,
		},
	}
	for _, test := range tests {
		fmt.Println("======================start")
		fmt.Println("nums:", test.nums)
		ans := getMinDistance(test.nums, test.target, test.start)
		fmt.Println("return:", ans)
		fmt.Println("result:", ans == test.except)
		fmt.Println("======================end")
	}
}

func getMinDistance(nums []int, target int, start int) int {
	ans := math.MaxInt
	for i := start; i >= 0; i-- {
		if nums[i] == target {
			ans = start - i
			break
		}
	}
	if ans == 0 {
		return ans
	}
	for i := start; i < len(nums); i++ {
		if nums[i] == target {
			ans = min(ans, i-start)
			break
		}
	}
	return ans
}

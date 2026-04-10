package main

import (
	"fmt"
	"math"
)

type Input struct {
	nums   []int
	except int
}

func main() {
	tests := []*Input{
		{
			nums:   []int{1, 2, 1, 1, 3},
			except: 6,
		},
	}
	for _, test := range tests {
		fmt.Println("======================start")
		fmt.Println("nums:", test.nums)
		ans := minimumDistance(test.nums)
		fmt.Println("return:", ans)
		fmt.Println("result:", ans == test.except)
		fmt.Println("======================end")
	}
}

// (x, y, z) y-x+z-x+z- y = 2(z-x)
func minimumDistance(nums []int) int {
	if len(nums) < 3 {
		return -1
	}
	numMap := map[int][]int{}
	for i, num := range nums {
		numMap[num] = append(numMap[num], i)
	}
	ans := math.MaxInt
	for _, dis := range numMap {
		length := len(dis)
		if length < 3 {
			continue
		}
		for i := 0; i < length-2; i++ {
			ans = min(ans, 2*(dis[i+2]-dis[i]))
		}
	}
	if ans == math.MaxInt {
		ans = -1
	}
	return ans
}

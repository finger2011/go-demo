package main

import "math"

func main() {

}

func findMissingElements(nums []int) []int {
	flag := make([]bool, 101)
	var minNum, maxNum int
	minNum = math.MaxInt
	for _, num := range nums {
		flag[num] = true
		if maxNum < num {
			maxNum = num
		}
		if minNum > num {
			minNum = num
		}
	}
	var ans []int
	for i := minNum; i <= maxNum; i++ {
		if !flag[i] {
			ans = append(ans, i)
		}
	}
	return ans
}

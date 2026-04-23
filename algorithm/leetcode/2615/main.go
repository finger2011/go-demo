package main

import "fmt"

func main() {
	// 5,0,3,4,0
	nums := []int{1, 3, 1, 1, 2}
	fmt.Println("dis:", distance(nums))
}

func distance(nums []int) []int64 {
	n := len(nums)
	numM := make(map[int][]int, n)
	ans := make([]int64, n)
	for i, num := range nums {
		numM[num] = append(numM[num], i)
	}
	for _, g := range numM {
		m := len(g)
		pre := make([]int, m+1)
		for i, p := range g {
			pre[i+1] = pre[i] + p
		}
		for i, p := range g {
			ans[p] = int64((i*p - pre[i]) + (pre[m] - pre[i] - p*(m-i)))
		}
	}
	return ans
}

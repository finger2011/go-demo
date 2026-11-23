package main

import (
	"fmt"
	"math/bits"
	"slices"
)

func main() {
	nums := []int{3, 5, 12}
	k := 1
	fmt.Println("min:", minSetBits(nums, k))
}

func minSetBits(nums []int, k int) int {
	n := len(nums)
	w := bits.Len(uint(slices.Max(nums)))

	res := make([]int, n-k+1)
	for i := 0; i+k <= n; i++ {
		v := nums[i]
		if v == 0 {
			return -1
		}
		for j := i + 1; j < i+k; j++ {
			v &= nums[j]
		}
		res[i] = v
	}
	fmt.Println("res:", res)
	bs := make([][]int, n-k+1)
	for i := range n - k + 1 {
		bs[i] = make([]int, w)
		for j := range w {
			// 第 j 个元素对应 bit (w-1-i)
			bs[i][j] = (res[i] >> (w - 1 - j)) & 1
		}
	}
	fmt.Println("bs:", bs)
	var ans int
	for i := range k {
		isNotZero := false
		for j := range n - k {
			if bs[j][i] == k {
				isNotZero = true
				break
			}
		}
		if isNotZero {
			ans++
		}
	}
	return ans
}

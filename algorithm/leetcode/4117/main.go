package main

import (
	"fmt"
	"slices"
)

func main() {
	a := 9
	b := 2
	c := 5
	fmt.Println("min:", minRange(a, b, c))
}

// 假设a b c从大到小排序
// 不做操作 极差 = a-c
// 1次
// c = a+b 极差 = a+b - b = a 序列 a+b a b
// a = b+c 极差 = b+c - c = b 序列 b+c b c
// b = a+c 极差 = a+c - c = a 序列 a+c a c
func minRange(a int, b int, c int) int {
	nums := []int{a, b, c}
	slices.Sort(nums)
	return min(nums[2]-nums[0], nums[1])
}

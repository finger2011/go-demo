package main

import "slices"

func main() {

}

func minimumEffort(tasks [][]int) (ans int) {
	slices.SortFunc(tasks, func(a, b []int) int {
		return (b[1] - b[0]) - (a[1] - a[0]) // 按照 minimum - actual 从大到小排序
	})

	s := 0 // 累计耗费的能量
	for _, t := range tasks {
		actual, minimum := t[0], t[1]
		// 题目要求 E0 - s >= minimum，即 E0 >= s + minimum
		// 由此可以得到 n 个关于 E0 的下界，所有下界的最大值即为答案
		ans = max(ans, s+minimum)
		s += actual
	}
	return
}

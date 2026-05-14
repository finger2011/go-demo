package main

import "math"

func main() {

}

func minMoves(nums []int, limit int) int {
	n := len(nums)
	diff := make([]int, limit*2+2)
	for i, x := range nums[:n/2] {
		y := nums[n-1-i]
		l := min(x, y) + 1
		r := max(x, y) + limit

		// [2, l-1] += 2
		diff[2] += 2
		diff[l] -= 2

		// [l, r] += 1
		diff[l]++
		diff[r+1]--

		// x+y 实际操作 0 次，从 [l, r] 中去掉
		// [x+y, x+y] -= 1
		diff[x+y]--
		diff[x+y+1]++

		// [r+1, limit*2] += 2
		diff[r+1] += 2
		// limit*2+1 不在 [2, limit*2] 中，可以省略
	}

	ans := math.MaxInt
	sum := 0
	for _, d := range diff[2 : limit*2+1] {
		sum += d
		ans = min(ans, sum)
	}
	return ans
}

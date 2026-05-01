package main

import "fmt"

func main() {
	// nums := []int{4, 3, 2, 6} // 26
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10} // 330
	fmt.Println("maxRotateFunction:", maxRotateFunction(nums))
}

func maxRotateFunction(nums []int) int {
	var ans, f, sum int
	for i, num := range nums {
		sum += num
		f += i * num
	}
	ans = f
	n := len(nums)
	for i := n - 1; i > 0; i-- {
		f += sum - n*nums[i]
		ans = max(ans, f)
	}
	return ans
}

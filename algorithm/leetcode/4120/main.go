package main

import "fmt"

func main() {
	nums := []int{96, 75, 55}
	fmt.Println("ans:", countDominantElements(nums))
}
func countDominantElements(nums []int) int {
	var ans int
	n := len(nums)
	if n <= 1 {
		return n
	}
	for i, num := range nums {
		if i == 0 && num >= nums[i+1] {
			ans++
		} else if i == n-1 && num >= nums[i-1] {
			ans++
		} else if i > 0 && i < n-1 && num >= (nums[i-1]+nums[i+1]+1)/2 {
			ans++
		}
	}
	return ans
}

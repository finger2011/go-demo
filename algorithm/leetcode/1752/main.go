package main

import "fmt"

func main() {
	nums := []int{2, 1, 3, 4} //false
	fmt.Println("check:", check(nums))
}

func check(nums []int) bool {
	sorted := true
	n := len(nums)
	if n == 1 {
		return true
	}
	for i := 1; i < n; i++ {
		if nums[i] < nums[i-1] {
			if !sorted {
				return false
			}
			sorted = false
		}
	}
	return sorted || nums[0] >= nums[n-1]
}

package main

import "fmt"

func main() {
	nums := []int{3, 4, 5, 1, 2}
	fmt.Println("min:", findMin(nums))
}

func findMin(nums []int) int {
	n := len(nums)
	if n == 1 || nums[0] < nums[n-1] {
		return nums[0]
	}
	l, r := 0, n-1
	for l < r {
		mid := (l + r + 1) / 2
		if mid > 0 && nums[mid-1] > nums[mid] {
			return nums[mid]
		}
		if mid < n-1 && nums[mid] > nums[mid+1] {
			return nums[mid+1]
		}
		if nums[mid] > nums[r] {
			l = mid + 1
		} else if nums[mid] < nums[l] {
			r = mid - 1
		}
	}
	return -1
}

package main

import (
	"fmt"
)

func main() {
	nums1 := []int{30, 29, 19, 5}
	nums2 := []int{25, 25, 25, 25, 25}
	fmt.Println("dis:", maxDistance(nums1, nums2))
}

func maxDistance(nums1 []int, nums2 []int) int {
	var ans int
	n := len(nums2)
	for i := range nums1 {
		start, end := i, n-1
		for start <= end {
			mid := (start + end + 1) / 2
			if nums2[mid] < nums1[i] {
				if mid == i {
					break
				}
				end = mid - 1
				if nums2[end] >= nums1[i] {
					ans = max(ans, end-i)
					break
				}
			} else {
				if mid+1 >= n {
					ans = max(ans, mid-i)
					break
				}
				start = mid + 1
			}
		}
	}
	return ans
}

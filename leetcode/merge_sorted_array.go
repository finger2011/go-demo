package main

import (
	"fmt"
)

//https://leetcode-cn.com/problems/merge-sorted-array/
//leetcode 88

func testMerge() {
	nums := merge([]int{2, 0}, 1, []int{1}, 1)
	fmt.Printf("nums:%v", nums)
}

func merge(nums1 []int, m int, nums2 []int, n int) []int {
	if n == 0 {
		return nums1
	}
	if m == 0 {
		nums1 = nums2
		return nums1
	}
	i := m - 1
	j := n - 1
	for ; i >= 0; i-- {
		for ; j >= 0; j-- {
			if nums2[j] >= nums1[i] {
				nums1[i+j+1] = nums2[j]
			} else {
				break
			}
		}
		nums1[i+j+1] = nums1[i]
	}
	if j >= 0 {
		for ; j >= 0; j-- {
			nums1[j] = nums2[j]
		}
	}
	return nums1
}

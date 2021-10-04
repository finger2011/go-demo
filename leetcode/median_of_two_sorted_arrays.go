package main

import (
	"fmt"
)

// https://leetcode-cn.com/problems/median-of-two-sorted-arrays/
// leetcode 4 hard

func testFindMedianSortedArrays() {
	// nums1 := []int{1, 3}
	// nums2 := []int{2, 4}
	// nums1 := []int{0, 0}
	// nums2 := []int{0, 0}
	// nums1 := []int{1, 3}
	// nums2 := []int{2, 4, 5, 6}
	nums1 := []int{4}
	nums2 := []int{1, 2,3,5,6}
	result := findMedianSortedArrays(nums1, nums2)
	fmt.Printf("result:%v\n", result)
}

func findMedianSortedArrays1(nums1 []int, nums2 []int) float64 {
	var result float64
	if len(nums1) == 0 && len(nums2) == 0 {
		return result
	}
	if len(nums1)+len(nums2) == 1 {
		if len(nums1) == 0 {
			return float64(nums2[0])
		}
		return float64(nums1[0])
	}
	if len(nums1) == 0 {
		nums1 = append(nums1, nums2[0])
		nums2 = nums2[1:]
	}
	if len(nums2) == 0 {
		nums2 = append(nums2, nums1[0])
		nums1 = nums1[1:]
	}
	if nums1[0] > nums2[0] {
		var tmp = nums1
		nums1 = nums2
		nums2 = tmp
	}
	var tar int
	var double bool
	if (len(nums1)+len(nums2))%2 == 0 {
		double = true
		tar = (len(nums1)+len(nums2))/2 - 1
	} else {
		double = false
		tar = (len(nums1) + len(nums2)) / 2
	}
	var cur1, cur2 int
	var second bool
	for cur1+cur2 < tar {
		if cur2 >= len(nums2) || (cur1 < len(nums1)-1 && nums1[cur1+1] <= nums2[cur2]) {
			cur1++
			second = false
			continue
		}
		cur2++
		second = true
	}
	if second {
		result = float64(nums2[cur2-1])
	} else {
		result = float64(nums1[cur1])
	}
	if double {
		if cur1 >= len(nums1)-1 {
			result = (result + float64(nums2[cur2])) / 2
		} else if cur2 > len(nums2)-1 {
			result = (result + float64(nums1[cur1+1])) / 2
		} else {
			if nums1[cur1+1] <= nums2[cur2] {
				result = (result + float64(nums1[cur1+1])) / 2
			} else {
				result = (result + float64(nums2[cur2])) / 2
			}
		}
	}

	return result
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if len(nums1) == 0 || len(nums1) < len(nums2) {
		var tmp = nums1
		nums1 = nums2
		nums2 = tmp
	}
	if len(nums2) == 0 {
		if len(nums1) == 0 {
			return 0
		} else if len(nums1)%2 == 1 {
			return float64(nums1[len(nums1)/2])
		} else {
			return (float64(nums1[len(nums1)/2] + nums1[len(nums1)/2 - 1]))/2
		}
	}
	imax := len(nums1)
	imin := 0
	for imin <= imax {
		i := (imin + imax) / 2
		j := (len(nums1) + len(nums2) + 1)/2 - i
		if j > 0 && i >= 0 && i < len(nums1) && nums2[j-1] > nums1[i] {
			imin = i + 1
		} else if i > 0 && j >=0 && j < len(nums2) && nums1[i-1] > nums2[j] {
			imax = i - 1
		} else {
			var maxLeft int
			if i == 0 {
				maxLeft = nums2[j-1]
			} else if j == 0 {
				maxLeft = nums1[i-1]
			} else {
				maxLeft = nums1[i-1]
				if nums1[i-1] < nums2[j-1] {
					maxLeft = nums2[j-1]
				}
			}
			if (len(nums1)+len(nums2))%2 == 1 {
				return float64(maxLeft)
			}
			var minRight int
			if i == len(nums1) {
				minRight = nums2[j]
			} else if j == len(nums2) {
				minRight = nums1[i]
			} else {
				minRight = nums1[i]
				if nums1[i] > nums2[j] {
					minRight = nums2[j]
				}
			}
			return (float64(minRight + maxLeft)) / 2

		}
	}
	return 0
}

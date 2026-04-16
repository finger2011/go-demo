package main

import (
	"fmt"
)

type Input struct {
	nums1, nums2 []int
	except       float64
}

func main() {
	tests := []*Input{
		{
			nums1:  []int{1, 3},
			nums2:  []int{2, 4},
			except: 2.5,
		},
	}
	for _, test := range tests {
		fmt.Println("============start")
		fmt.Println("nums1:", test.nums1)
		fmt.Println("nums2:", test.nums2)
		ans := findMedianSortedArrays(test.nums1, test.nums2)
		fmt.Println("return:", ans)
		fmt.Println("result:", test.except == ans)
		fmt.Println("============end")
	}
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	// m, n := len(nums1), len(nums2)
	// if m < n {
	// 	m, n = n, m
	// 	nums1, nums2 = nums2, nums1
	// }
	// var isOdd bool
	// if (m+n)%2 != 0 {
	// 	isOdd = true
	// 	if nums1[m-1] > nums2[n-1] {
	// 		nums1 = append(nums1, math.MaxInt)
	// 		m++
	// 	} else {
	// 		nums2 = append(nums2, math.MaxInt)
	// 		n++
	// 	}
	// }
	// length := (m + n) / 2
	// start, end := 0, m-1
	// for {
	// 	mid := (start + end) / 2
	// 	mid2 := length - mid
	// }
}

// 时间复杂度为O(m+n), 不符合题意
func findMedianSortedArrays2(nums1 []int, nums2 []int) float64 {
	m, n := len(nums1), len(nums2)
	pos1 := (m + n) / 2
	pos2 := pos1
	if (m+n)%2 == 0 {
		pos2--
	}
	var i, j, k, cur, n1, n2 int
	for k <= pos1 {
		if i >= m {
			cur = nums2[j]
			j++
		} else if j >= n {
			cur = nums1[i]
			i++
		} else if nums1[i] < nums2[j] {
			cur = nums1[i]
			i++
		} else {
			cur = nums2[j]
			j++
		}
		if k == pos1 {
			n1 = cur
		}
		if k == pos2 {
			n2 = cur
		}
		k++
	}

	return float64(n1+n2) / 2
}

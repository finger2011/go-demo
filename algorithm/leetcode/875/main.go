package main

import (
	"fmt"
	"sort"
)

func main() {
	piles := []int{3, 6, 7, 11}
	h := 8
	fmt.Println("minEatingSpeed (", piles, ",", h, ") ===> ", minEatingSpeed(piles, h))
	fmt.Println("minEatingSpeed2(", piles, ",", h, ") ===> ", minEatingSpeed2(piles, h))
}

// 二分查找
func minEatingSpeed(piles []int, h int) int {
	var left, right, ans int
	left = 1
	for _, pile := range piles {
		if pile > right {
			right = pile
		}
	}
	ans = right
	for left <= right {
		mid := (left + right) / 2
		time := 0
		for _, pile := range piles {
			time += (pile + mid - 1) / mid
		}
		if time <= h {
			ans = mid
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return ans
}

// sort.Search自动使用二分查找
func minEatingSpeed2(piles []int, h int) int {
	max := 0
	for _, pile := range piles {
		if pile > max {
			max = pile
		}
	}
	return 1 + sort.Search(max-1, func(speed int) bool {
		speed++
		time := 0
		for _, pile := range piles {
			time += (pile + speed - 1) / speed
		}
		return time <= h
	})
}

package main

import (
	"fmt"
)

func main() {
	// queries := [][]int{{1, 2}, {2, 3, 3}, {2, 3, 1}, {2, 2, 2}} // false true true
	// queries := [][]int{{2, 1, 1}} // false
	// queries := [][]int{{1, 1}, {1, 11}, {1, 4}, {1, 8}, {2, 13, 7}} // false
	queries := [][]int{{1, 2}, {1, 13}, {2, 10, 3}} // true
	fmt.Println("res:", getResults(queries))
}

func getResults(queries [][]int) []bool {
	ans := []bool{}
	blocks := []int{}
	maxBlocks := []int{}
	var blockLen, blockNum int
	check := func(x, length int) bool {
		if length > x {
			return false
		}
		if len(blocks) == 0 || x < blocks[0] {
			return true
		}
		idx := searchInsert(blocks, x)
		if idx == blockNum {
			if maxBlocks[blockNum-1] >= length || x-blocks[blockNum-1] >= length {
				return true
			}
		} else {
			if maxBlocks[idx] >= length || x-blocks[idx] >= length {
				return true
			}
		}
		return false
	}
	for _, query := range queries {
		blockLen++
		if query[0] == 1 {
			if len(blocks) == 0 {
				blocks = []int{query[1]}
				maxBlocks = []int{query[1]}
				blockNum++
			} else if query[1] <= blocks[0] {
				blocks = append([]int{query[1]}, blocks...)
				maxBlocks = append([]int{query[1]}, maxBlocks...)
				for i := range blockNum {
					maxBlocks[i+1] = max(maxBlocks[i], blocks[i+1]-blocks[i])
				}
				blockNum++
			} else if query[1] >= blocks[blockNum-1] {
				blocks = append(blocks, query[1])
				maxBlocks = append(maxBlocks, max(maxBlocks[blockNum-1], query[1]-blocks[blockNum-1]))
				blockNum++
			} else {
				idx := searchInsert(blocks, query[1])
				blocks = append(blocks[:idx], append([]int{query[1]}, blocks[idx:]...)...)
				maxBlocks = append(maxBlocks[:idx], append([]int{query[1]}, maxBlocks[idx:]...)...)
				blockNum++
				for i := idx; i < blockNum; i++ {
					maxBlocks[i] = max(maxBlocks[i-1], blocks[i]-blocks[i-1])
				}
			}
		} else {
			ans = append(ans, check(query[1], query[2]))
		}
	}
	return ans
}

func searchInsert(nums []int, target int) int {
	left, right := 0, len(nums) // 注意 right 初始为 len(nums)，表示可能插入末尾
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] >= target {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}

package main

import (
	"fmt"
	"slices"
)

func main() {
	n := 4
	seats := [][]int{{1, 7}, {4, 3}, {4, 6}, {1, 4}}
	fmt.Println("max:", maxNumberOfFamilies(n, seats))
}

// 位运算
func maxNumberOfFamilies(n int, reservedSeats [][]int) int {
	seats := map[int]int{} // 2~9 有预定座位的行 -> 这一行具体哪些座位被预定
	for _, r := range reservedSeats {
		seat := r[1]
		if 2 <= seat && seat <= 9 {
			seats[r[0]] |= 1 << (seat - 2) // 把二进制数的 seat-2 这一位变成 1
		}
	}

	ans := (n - len(seats)) * 2
	for _, x := range seats {
		// 在哈希表中的行，由于 2~9 至少一个座位被预定，所以至多容纳 1 个四人小组，ans 至多增加 1
		if x&0b1111 == 0 || x&0b111100 == 0 || x&0b11110000 == 0 {
			ans++
		}
	}
	return ans
}

// 超时
func maxNumberOfFamilies3(n int, reservedSeats [][]int) int {
	var ans, ops int
	slices.SortFunc(reservedSeats, func(a, b []int) int {
		return a[0] - b[0]
	})
	length := len(reservedSeats)
	for i := 1; i <= n; i++ {
		if ops >= length || reservedSeats[ops][0] > i {
			ans += 2
			continue
		}
		left, right, mid := true, true, true
		j := ops
		for ; j < length && reservedSeats[j][0] <= i; j++ {
			if reservedSeats[j][1] > 1 && reservedSeats[j][1] < 6 {
				left = false
				if reservedSeats[j][1] > 3 {
					mid = false
				}
			} else if reservedSeats[j][1] > 5 && reservedSeats[j][1] < 10 {
				right = false
				if reservedSeats[j][1] < 8 {
					mid = false
				}
			}
		}
		ops = j
		if left || right {
			if left {
				ans++
			}
			if right {
				ans++
			}
		} else if mid {
			ans++
		}
	}
	return ans
}

// n = 1000000000 out of memory
func maxNumberOfFamilies2(n int, reservedSeats [][]int) int {
	grid := make([][]bool, n)
	for i := range n {
		grid[i] = make([]bool, 10)
	}
	for _, seat := range reservedSeats {
		grid[seat[0]-1][seat[1]-1] = true
	}
	var ans int
	for i := range n {
		var left, right, mid bool
		for j := 1; j < 9; j++ {
			if grid[i][j] {
				if j < 5 {
					left = true
				} else {
					right = true
				}
				if j > 2 && j < 7 {
					mid = true
				}
			}
		}
		if (!left) || (!right) {
			if !left {
				ans++
			}
			if !right {
				ans++
			}
		} else if !mid {
			ans++
		}
	}
	return ans
}

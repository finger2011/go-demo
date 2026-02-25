package main

import (
	"fmt"
	"math/bits"
	"sort"
)

func main() {
	arr := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println("sortByBits(", arr, ") ====> ", sortByBits(arr))
}

func sortByBits(arr []int) []int {
	sort.Slice(arr, func(i, j int) bool {
		x, y := arr[i], arr[j]
		countX, countY := bits.OnesCount(uint(x)), bits.OnesCount(uint(y))
		return countX < countY || countX == countY && x < y
	})
	return arr
}

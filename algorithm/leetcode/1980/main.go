package main

import (
	"fmt"
	"strconv"
)

func main() {
	nums := []string{"00", "01"}
	fmt.Println("findDifferentBinaryString(", nums, ") ===>", findDifferentBinaryString(nums))
}

func findDifferentBinaryString(nums []string) string {
	n := len(nums)
	nMap := make(map[int]bool, n)
	for i, _ := range nums {
		num, _ := strconv.ParseInt(nums[i], 2, 32)
		nMap[int(num)] = true
	}
	var num int
	for nMap[num] {
		num++
	}
	return fmt.Sprintf("%0"+strconv.Itoa(n)+"b", num)
}

package main

import (
	"fmt"
	"strconv"
)

func main() {
	nums := []string{"00", "10"}
	fmt.Println("findDifferentBinaryString (", nums, ") ===>", findDifferentBinaryString(nums))
	fmt.Println("findDifferentBinaryString2(", nums, ") ===>", findDifferentBinaryString2(nums))
}

// 字符串先转二进制数组存入map，
// 找到第一个不在map中的数字，转为二进制字符串输出
func findDifferentBinaryString(nums []string) string {
	n := len(nums)
	nMap := make(map[int]bool, n)
	for i := range nums {
		num, _ := strconv.ParseInt(nums[i], 2, 32)
		nMap[int(num)] = true
	}
	var num int
	for nMap[num] {
		num++
	}
	return fmt.Sprintf("%0"+strconv.Itoa(n)+"b", num)
}

// 每一位都与字符串中的某一位字符不同
// 关键在于长度n，只有n个字符串
func findDifferentBinaryString2(nums []string) string {
	ans := make([]byte, len(nums))
	for i, s := range nums {
		ans[i] = s[i] ^ 1
	}
	return string(ans)
}

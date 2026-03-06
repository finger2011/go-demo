package main

import "fmt"

func main() {
	s := "1001"
	fmt.Println("checkOnesSegment (", s, ") ===>", checkOnesSegment(s))
	fmt.Println("checkOnesSegment2(", s, ") ===>", checkOnesSegment2(s))
}

// 判断字符串是否含01来判断
// 0个连续1，不可能，因为题目明确描述不含前导零
// 1个连续1，即字符串应该为11...111000...000，判断包含01即可
func checkOnesSegment(s string) bool {
	length := len(s)
	for i := 0; i < length-1; i++ {
		if s[i] == '0' && s[i+1] == '1' {
			return false
		}
	}
	return true
}

// 通过计算连续1的数量，来判断是否符合
func checkOnesSegment2(s string) bool {
	length := len(s)
	if length <= 2 {
		return true
	}
	var ans int
	var conti bool
	for i := 1; i < length; i++ {
		if (!conti) && s[i] == '0' {
			conti = true
			ans++
		} else if conti && s[i] == '1' {
			conti = false
		}
	}
	if !conti {
		ans++
	}
	return ans <= 1
}

package main

import "fmt"

func main() {
	s := "10010100"
	fmt.Println("minOperations(", s, ") ===> ", minOperations(s))
}

// 最终结果为2种，开头为1，或者开头为0,2者的结果相加等于字符串的长度 ans0 + ans1 = len(s)
func minOperations(s string) int {
	var ans int
	length := len(s)
	if length <= 1 {
		return ans
	}
	for i := 0; i < length; i++ {
		if i%2 != int(s[i]-'0') {
			ans++
		}
	}
	if length-ans < ans {
		ans = length - ans
	}
	return ans
}

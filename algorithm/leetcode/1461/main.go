package main

import "fmt"

func main() {
	// true
	s, k := "00110110", 2
	// false
	// s, k := "0110", 2
	fmt.Println("hasAllCodes :s(", s, "), k(", k, ") ==> ", hasAllCodes(s, k))
	fmt.Println("hasAllCodes2:s(", s, "), k(", k, ") ==> ", hasAllCodes2(s, k))
}

func hasAllCodes(s string, k int) bool {
	length := len(s)
	if length <= k {
		return false
	}
	maxCount := (1 << k) + k - 1
	if length < maxCount {
		return false
	}
	strMap := make(map[string]bool)
	for i := 0; i+k <= length; i++ {
		ops := s[i : i+k]
		strMap[ops] = true
	}
	return len(strMap) == (1 << k)
}

func hasAllCodes2(s string, k int) bool {
	length := len(s)
	if length <= k {
		return false
	}
	maxCount := (1 << k) + k - 1
	if length < maxCount {
		return false
	}
	ops := 0
	for i := 0; i < k; i++ {
		ops = ops << 1
		if s[i] == '1' {
			ops |= 1
		}
	}
	intMap := make(map[int]bool)
	intMap[ops] = true
	for i := 1; i+k <= length; i++ {
		ops = (ops-(int(s[i-1]-'0')<<(k-1)))<<1 + int(s[i+k-1]-'0')
		intMap[ops] = true
	}
	return len(intMap) == (1 << k)
}

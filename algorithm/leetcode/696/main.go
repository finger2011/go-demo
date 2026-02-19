package main

import "fmt"

func main() {
	s := "00110011"
	fmt.Println(s, " countBinarySubstrings:", countBinarySubstrings(s))
	n, ss := countBinarySubstrings2(s)
	fmt.Println(s, " countBinarySubstrings2:", n)
	fmt.Println(ss)
}

func countBinarySubstrings(s string) int {
	var cur, last, ans int
	length := len(s)
	for cur < length {
		ch := s[cur]
		count := 0
		for cur < length && s[cur] == ch {
			cur++
			count++
		}
		ans += min(count, last)
		last = count
	}
	return ans
}

// 同时输出子串
func countBinarySubstrings2(s string) (int, []string) {
	var cur, last, ans int
	var res []string
	length := len(s)
	for cur < length {
		ch := s[cur]
		count := 0
		for cur < length && s[cur] == ch {
			cur++
			count++
		}
		l := min(count, last)
		start, end := cur-2*l, cur
		for i := 0; i < l; i++ {
			// "00110011"  l = 2 cur = 4 start = cur - 2 * l
			// i = 0, start = start , end = cur
			// i = 1, start = start + 1, end =
			res = append(res, s[start+i:end-i])
		}
		ans += l
		last = count
	}
	return ans, res
}

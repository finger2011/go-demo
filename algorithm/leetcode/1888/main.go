package main

import "fmt"

func main() {
	s := "01001001101" // 2
	fmt.Println("minFlips (", s, ") ===>", minFlips(s))
}

// 双倍字符串+滑动窗口，{0, 0, 1} => {0, 0, 1, 0, 0, 1}
// {[0, 0, 1], 0, 0, 1}
// {0, [0, 1, 0], 0, 1}
// {0, 0, [1, 0, 0], 1}
// {0, 0, 1, [0, 0, 1]}
// 不需要拼接双倍，用 (i + length) % 2
func minFlips(s string) int {
	length := len(s)
	var ans, cnt int
	for i := 0; i < length; i++ {
		if i%2 == 0 && s[i] == '0' {
			cnt++
		} else if i%2 != 0 && s[i] == '1' {
			cnt++
		}
	}
	ans = min(cnt, length-cnt)
	for i := 0; i < length; i++ {
		if i%2 == 0 && s[i] == '0' {
			cnt--
		} else if i%2 != 0 && s[i] == '1' {
			cnt--
		}
		if (i+length)%2 == 0 && s[i] == '0' {
			cnt++
		} else if (i+length)%2 != 0 && s[i] == '1' {
			cnt++
		}
		ans = min(ans, cnt, length-cnt)
	}
	return ans
}

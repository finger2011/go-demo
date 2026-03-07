package main

import "fmt"

func main() {
	s := "01001001101" // 2
	fmt.Println("minFlips (", s, ") ===>", minFlips(s))
	fmt.Println("minFlips2(", s, ") ===>", minFlips(s))
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

// 前缀和+后缀和
// 偶数长度的字符串不需要后缀和
// 奇数长度的字符串，用`|`隔开，为需要类型1的操作顺序
// 可能为010101....010|0101...010 或者 1010...101|1010....101
func minFlips2(s string) int {
	var ans int
	length := len(s)
	pre := make([][2]int, length)
	for i := 0; i < length; i++ {
		if i == 0 {
			pre[i][0] = isChange(s[i], 1)
			pre[i][1] = isChange(s[i], 0)
		} else {
			pre[i][0] = isChange(s[i], 1) + pre[i-1][1]
			pre[i][1] = isChange(s[i], 0) + pre[i-1][0]
		}
	}
	ans = min(pre[length-1][0], pre[length-1][1])
	if length%2 == 1 {
		suf := make([][2]int, length)
		for i := length - 1; i >= 0; i-- {
			if i == length-1 {
				suf[i][0] = isChange(s[i], 1)
				suf[i][1] = isChange(s[i], 0)
			} else {
				suf[i][0] = isChange(s[i], 1) + suf[i+1][1]
				suf[i][1] = isChange(s[i], 0) + suf[i+1][0]
			}
		}
		for i := 0; i < length-1; i++ {
			ans = min(ans, pre[i][0]+suf[i+1][0], pre[i][1]+suf[i+1][1])
		}
	}
	return ans
}

func isChange(ch byte, n int) int {
	if int(ch-'0') == n {
		return 1
	}
	return 0
}

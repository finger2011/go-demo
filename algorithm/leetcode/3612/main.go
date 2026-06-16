package main

import "fmt"

func main() {
	s := "a#b%*"
	fmt.Println("str:", processStr(s))
}

func processStr(s string) string {
	var ans string
	for _, ch := range s {
		switch ch {
		case '*':
			if len(ans) > 0 {
				ans = ans[:len(ans)-1]
			}
		case '#':
			ans += ans
		case '%':
			ans = revertString(ans)
		default:
			ans += string(ch)
		}
	}
	return ans
}

func revertString(s string) string {
	length := len(s)
	if length <= 0 {
		return s
	}
	ans := make([]byte, length)
	i, j := 0, length-1
	for i <= j {
		ans[i], ans[j] = s[j], s[i]
		i++
		j--
	}
	return string(ans)
}

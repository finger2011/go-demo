package main

import "fmt"

func main() {
	s := "(()))())("
	// s := ")()()"
	fmt.Println("longestValidParentheses[", s, "] ===> ", longestValidParentheses(s))
}

func longestValidParentheses(s string) int {
	stack := []byte{}
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case '(':
			stack = append(stack, s[i])
		case ')':
			if len(stack) == 0 {
				stack = append(stack, '|')
			} else {
				ops := len(stack) - 1
				for ops >= 0 {
					if stack[ops] == '(' {
						stack = append(stack[:ops], stack[ops+1:]...)
						stack = append(stack, '2')
						break
					} else if byteIsInt(stack[ops]) {
						ops--
					} else {
						stack = append(stack, '|')
						break
					}
				}
				if ops < 0 {
					stack = append(stack, '|')
				}
			}
		}
	}
	fmt.Println("stack:")
	for _, b := range stack {
		fmt.Printf("%s", string(b))
	}
	var ops, num, ans int
	var isInt bool
	for ops < len(stack) {
		if byteIsInt(stack[ops]) {
			if isInt {
				num += int(stack[ops] - '0')
			} else {
				num = int(stack[ops] - '0')
				isInt = true
			}
		} else {
			if isInt && num > ans {
				ans = num
			}
			isInt = false
			num = 0
		}
		ops++
	}
	if num > ans {
		ans = num
	}
	return ans
}

func byteIsInt(b byte) bool {
	if int(b-'0') >= 0 && int(b-'0') <= 9 {
		return true
	}
	return false
}

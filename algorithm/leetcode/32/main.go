package main

import "fmt"

func main() {
	// s := "(()))())("
	s := ")()()"
	fmt.Println("longestValidParentheses [", s, "] ===> ", longestValidParentheses(s))
	fmt.Println("longestValidParentheses2[", s, "] ===> ", longestValidParentheses2(s))
}

// 一次遍历字符串，一次遍历stack
// 遍历字符串，碰到匹配的左右括号则设置stack为2，有未匹配的左右括号则将stack设置为`|`表示终止，如"(()))())("的结果为22|2|(
// 遍历stack 碰到连续的2累加，其他则中断累加，取最大的长度
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
	// fmt.Println("stack:")
	// for _, b := range stack {
	// 	fmt.Printf("%s", string(b))
	// }
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

// 保持栈底元素为最后一个未被匹配的`)`的下标，为了防止第一个字符为`(`时不匹配，需要开始写入一个-1
// 遇到`(`则将下标压入栈
// 遇到`)` 先弹出栈底，表示找到了匹配的`(`，如果栈空，则直接入栈，否则用当前下标减去栈底下标，是该连续合法串的长度
// 找到其中最大长度
func longestValidParentheses2(s string) int {
	var ans int
	stack := []int{}
	stack = append(stack, -1)
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			stack = append(stack, i)
		} else {
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				stack = append(stack, i)
			} else {
				if i-stack[len(stack)-1] > ans {
					ans = i - stack[len(stack)-1]
				}
			}
		}
	}
	return ans
}

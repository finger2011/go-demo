package main

import "fmt"

func main() {
	s := "3 1+22*28 "
	// s := "3/2 "
	fmt.Println("calculate(", s, ") ===> ", calculate(s))
}

func calculate(s string) int {
	var ans int
	length := len(s)
	if length == 0 {
		return ans
	}
	intStacks := []int{}
	optStacks := []byte{}
	var ops int
	for ops < length {
		if byteIsInt(s[ops]) {
			i := ops + 1
			num := int(s[ops] - '0')
			for i < length && (byteIsInt(s[i]) || s[i] == ' ') {
				if s[i] != ' ' {
					num = num*10 + int(s[i]-'0')
				}
				i++
			}
			if len(optStacks) > 0 {
				if optStacks[len(optStacks)-1] == '*' {
					num = intStacks[len(intStacks)-1] * num
					intStacks[len(intStacks)-1] = num
					optStacks = optStacks[:len(optStacks)-1]
				} else if optStacks[len(optStacks)-1] == '/' {
					num = intStacks[len(intStacks)-1] / num
					intStacks[len(intStacks)-1] = num
					optStacks = optStacks[:len(optStacks)-1]
				} else {
					intStacks = append(intStacks, num)
				}
			} else {
				intStacks = append(intStacks, num)
			}

			ops = i
		} else if s[ops] == '+' || s[ops] == '-' || s[ops] == '*' || s[ops] == '/' {
			optStacks = append(optStacks, s[ops])
			ops++
		} else {
			ops++
		}
	}
	fmt.Println("intStacks:", intStacks, ";optStacks", optStacks)
	ans = intStacks[0]
	for i := 0; i < len(optStacks); i++ {
		switch optStacks[i] {
		case '+':
			ans += intStacks[i+1]
		case '-':
			ans -= intStacks[i+1]
		default:
		}
	}
	return ans
}

func byteIsInt(b byte) bool {
	if int(b-'0') >= 0 && int(b-'0') <= 9 {
		return true
	}
	return false
}

package main

import "fmt"

func main() {
	s := "1101"
	fmt.Println("numSteps(", s, ") ==> ", numSteps(s))
}

func numSteps(s string) int {
	var ans int
	var carry bool
	if len(s) <= 1 {
		return ans
	}
	for i := len(s) - 1; i >= 0; i-- {
		if carry {
			if s[i] == '0' {
				ans += 2
			} else {
				ans++
				carry = true
			}
		} else {
			if s[i] == '0' {
				ans++
			} else {
				if i != 0 {
					ans += 2
				}
				carry = true
			}
		}
	}
	return ans
}

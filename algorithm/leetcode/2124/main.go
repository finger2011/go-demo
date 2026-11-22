package main

import "fmt"

func main() {
	// s := "aaaa"
	// s := "bbb"
	s := "abab" // false
	fmt.Println("check:", checkString(s))
}

func checkString(s string) bool {
	var hasB bool
	for _, str := range s {
		if str == 'a' && hasB {
			return false
		} else if str == 'b' {
			hasB = true
		}
	}
	return true
}

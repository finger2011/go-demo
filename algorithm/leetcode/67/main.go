package main

import (
	"fmt"
	"strconv"
)

func main() {
	a, b := "11", "1"
	fmt.Println("addBinary (", a, ",", b, ") ==>", addBinary(a, b))
	fmt.Println("addBinary2(", a, ",", b, ") ==>", addBinary(a, b))
}

func addBinary2(a string, b string) string {
	var ans string
	var carry int
	aLen, bLen := len(a), len(b)
	length := max(aLen, bLen)
	for i := 0; i < length; i++ {
		if i < aLen {
			carry += int(a[aLen-1-i] - '0')
		}
		if i < bLen {
			carry += int(b[bLen-1-i] - '0')
		}
		ans = strconv.Itoa(carry%2) + ans
		carry /= 2
	}
	if carry > 0 {
		ans = "1" + ans
	}
	return ans
}

func addBinary(a string, b string) string {
	var ans string
	var carry bool
	if len(a) < len(b) {
		a, b = b, a
	}
	aPos, bPos := len(a)-1, len(b)-1
	for aPos >= 0 || bPos >= 0 {
		var tmp string
		if a[aPos] == '0' {
			if bPos >= 0 {
				if b[bPos] == '0' {
					if carry {
						tmp = "1"
						carry = false
					} else {
						tmp = "0"
					}
				} else {
					if carry {
						tmp = "0"
					} else {
						tmp = "1"
						carry = false
					}
				}
			} else {
				if carry {
					carry = false
					tmp = "1"
				} else {
					tmp = "0"
				}
			}
		} else {
			if bPos >= 0 {
				if b[bPos] == '0' {
					if carry {
						tmp = "0"
					} else {
						tmp = "1"
						carry = false
					}
				} else {
					if carry {
						tmp = "1"
					} else {
						tmp = "0"
					}
					carry = true
				}
			} else {
				if carry {
					tmp = "0"
				} else {
					tmp = "1"
				}
			}
		}
		ans = tmp + ans
		// fmt.Println(ans, "===>", carry)
		aPos--
		bPos--
	}
	if carry {
		ans = "1" + ans
	}

	return ans
}

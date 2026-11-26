package main

import "fmt"

func main() {
	n := 19
	fmt.Println("can win:", canAliceWin(n))
}

func canAliceWin(n int) bool {
	num := 10
	var canWin bool
	for n > 0 && num > 0 {
		if n < num {
			return canWin
		}
		n -= num
		num--
		canWin = !canWin
	}
	return canWin
}

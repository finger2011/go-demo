package main

import (
	"fmt"
	"math"
)

func main() {
	var num int
	num = 3 / 2
	fmt.Println("num:", num)
}

// balloon
func maxNumberOfBalloons(text string) int {
	textMap := make(map[rune]int, 5)
	textMap['b'] = 0
	textMap['a'] = 0
	textMap['n'] = 0
	textMap['l'] = 0
	textMap['o'] = 0
	for _, ch := range text {
		switch ch {
		case 'b', 'a', 'l', 'o', 'n':
			textMap[ch]++
		default:
		}
	}
	ans := math.MaxInt
	for ch, num := range textMap {

		switch ch {
		case 'b', 'a', 'n':
			if num == 0 {
				return 0
			}
			ans = min(ans, num)
		case 'l', 'o':
			if num < 2 {
				return 0
			}
			ans = min(ans, num/2)
		}
	}
	return ans
}

package main

import "fmt"

func main() {
	piles := []int{3, 7, 2, 3}
	fmt.Println("game:", stoneGame(piles))
}

// 数学方式证明有必胜把握
func stoneGame(piles []int) bool {
	return true
}

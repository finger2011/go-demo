package main

import "fmt"

func main() {
	num := 150
	fmt.Println("num:", mirrorDistance(num))
}

func mirrorDistance(n int) int {
	var revertN int
	i := n
	for i >= 10 {
		revertN = revertN*10 + i%10
		i /= 10
	}
	revertN = revertN*10 + i
	if revertN > n {
		return revertN - n
	}
	return n - revertN
}

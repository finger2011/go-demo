package main

import "fmt"

func main() {
	n := 5
	fmt.Println(n, " is ", hasAlternatingBits(n))
}

func hasAlternatingBits(n int) bool {
	ans := n ^ (n >> 1)
	return ans&(ans+1) == 0
}

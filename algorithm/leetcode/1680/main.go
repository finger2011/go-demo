package main

import (
	"fmt"
	"math/bits"
)

func main() {
	n := 12
	fmt.Println("concatenatedBinary(", n, ") ==> ", concatenatedBinary(n))
}

func concatenatedBinary(n int) int {
	var ans int
	mod := 1000000007
	for i := 1; i <= n; i++ {
		ans = (ans<<(bits.Len(uint(i))) | i) % mod
	}
	return ans
}

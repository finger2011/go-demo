package main

import (
	"fmt"
	"math/bits"
)

func main() {
	n := 0
	fmt.Println("bitwiseComplement :(", n, ") ===> ", bitwiseComplement(n))
	fmt.Println("bitwiseComplement2:(", n, ") ===> ", bitwiseComplement2(n))
}

// 5 : 101 => 010 => 2
func bitwiseComplement(n int) int {
	var ans, arr int
	if n == 0 {
		return 1
	}
	for n > 0 {
		if n%2 == 0 {
			ans += 1 << arr
		}
		n >>= 1
		arr++
	}
	return ans
}

// math.bits库
func bitwiseComplement2(n int) int {
	if n == 0 {
		return 1
	}
	return n ^ (1<<bits.Len(uint(n)) - 1)
}

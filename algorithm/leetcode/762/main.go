package main

import (
	"fmt"
	"math/bits"
)

func main() {
	left, right := 6, 10
	fmt.Println("countPrimeSetBits from ", left, "to ", right, ":", countPrimeSetBits(left, right))
}

func countPrimeSetBits(left int, right int) int {
	primeMap := map[int]bool{
		2:  true,
		3:  true,
		5:  true,
		7:  true,
		11: true,
		13: true,
		17: true,
		19: true,
		23: true,
	}
	var ans int
	for n := left; n <= right; n++ {
		count := bits.OnesCount(uint(n))
		if primeMap[count] {
			ans++
		}
	}
	return ans
}

// 判断是否为质数
func isPrime(x int) bool {
	if x < 2 {
		return false
	}
	for i := 2; i*i <= x; i++ {
		if x%i == 0 {
			return false
		}
	}
	return true
}

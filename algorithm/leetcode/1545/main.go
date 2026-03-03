package main

import "fmt"

func main() {
	n, k := 4, 11
	fmt.Println("findKthBit(n:", n, ", k:", k, ") ====> ", string(findKthBit(n, k)))
}

// len(n) = 2 ** n - 1
// 第2 ** i 次位肯定是1
// 前半段(2**(n - 1) - 1)是n - 1
// 后半段(2**(n - 1) - 1)是n - 1的反转再翻转，翻转则是找（2 ** n - k），反转是异或(XOR)操作
func findKthBit(n int, k int) byte {
	if n == 1 {
		return '0'
	}
	mid := 1 << (n - 1)
	if k == mid {
		return '1'
	} else if k < mid {
		return findKthBit(n-1, k)
	} else {
		return findKthBit(n-1, 1<<n-k) ^ 1
	}
}

package main

import "fmt"

func main() {
	n := 22
	fmt.Println(n, "binaryGap:", binaryGap(n))
}

func binaryGap(n int) int {
	var ans int
	for i, last := 0, -1; n > 0; i++ {
		if n&1 == 1 {
			if last != -1 && (i-last) > ans {
				ans = i - last
			}
			last = i
		}
		n >>= 1
	}
	return ans
}

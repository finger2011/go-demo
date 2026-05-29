package main

import (
	"fmt"
	"math"
)

func main() {
	nums := []int{999, 19, 199}
	fmt.Println("min:", minElement(nums))
}

func minElement(nums []int) int {
	ans := math.MaxInt
	var tmp int
	for _, num := range nums {
		tmp = 0
		for num > 0 {
			tmp += num % 10
			num /= 10
			if tmp > ans {
				break
			}
		}
		if ans > tmp {
			ans = tmp
		}
	}
	return ans
}

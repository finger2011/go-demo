package main

import (
	"fmt"
	"math"
)

func main() {
	n := 99999
	fmt.Println("bulbSwitch(", n, ") ===> ", bulbSwitch(n))
}

// 只有完全平方数才会亮着
// 判断第i灯泡是否亮着，取决于i的约数的个数，只有奇数个数的约数才会亮着
// 有奇数个数约数的i只有完全平方数：因为如果存在x是i的约数，那么i/x也是i的约数，这时只有x = i/x时，才有奇数个约数
func bulbSwitch(n int) int {
	return int(math.Sqrt(float64(n)))
}

package main

import (
	"fmt"
	"math"
)

func main() {
	n := 13
	fmt.Println("num:", numSquares2(n))
}

var memo [101][10001]int

func init() {
	for i := range memo {
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}
}
func numSquares(n int) int {
	var dfs func(i, j int) int
	dfs = func(i, j int) int {
		if j == 0 {
			return 0
		}
		if i == 0 {
			return math.MaxInt
		}
		p := &memo[i][j]
		if *p != -1 {
			return *p
		}
		if j < i*i {
			*p = dfs(i-1, j)
		} else {
			*p = min(dfs(i-1, j), dfs(i, j-i*i)+1)
		}
		return *p
	}
	return dfs(int(math.Sqrt(float64(n))), n)
}

const mx = 10001

var f [mx]int

func init() {
	for i := 1; i < mx; i++ {
		f[i] = math.MaxInt
	}
	for i := 1; i*i < mx; i++ {
		for j := i * i; j < mx; j++ {
			f[j] = min(f[j], f[j-i*i]+1) // 不选 vs 选
		}
	}
}

func numSquares2(n int) int {
	return f[n]
}

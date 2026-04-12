package main

import (
	"fmt"
	"math"
	"slices"
)

type Input struct {
	word   string
	except int
}

func main() {
	tests := []*Input{
		{
			word:   "CAKE",
			except: 3,
		},
	}
	for _, test := range tests {
		fmt.Println("======================start")
		fmt.Println("word:", test.word)
		ans := minimumDistance2(test.word)
		fmt.Println("return:", ans)
		fmt.Println("result:", ans == test.except)
		fmt.Println("======================end")
	}
}

var dis [26][26]int

const column = 6

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func init() {
	for i := range 26 {
		for j := range 26 {
			dis[i][j] = abs(i/column-j/column) + abs(i%column-j%column)
		}
	}
}

func minimumDistance(word string) int {
	n := len(word)
	memo := make([][26][26]int, n)
	var dfs func(int, byte, byte) int
	dfs = func(i int, f1, f2 byte) (ans int) {
		if i < 0 {
			return 0
		}
		p := &memo[i][f1][f2]
		if *p != 0 {
			return *p - 1
		}
		defer func() {
			*p = ans + 1
		}()
		ch := word[i] - 'A'
		ans1 := dfs(i-1, ch, f2) + dis[f1][ch]
		ans2 := dfs(i-1, f1, ch) + dis[f2][ch]
		return min(ans1, ans2)
	}
	ans := math.MaxInt
	for f2 := range byte(26) {
		ans = min(ans, dfs(n-2, word[n-1]-'A', f2))
	}
	return ans
}

func minimumDistance2(word string) int {
	var f, nf [26]int
	for i := range len(word) - 1 {
		x, y := word[i]-'A', word[i+1]-'A'
		for f2 := range 26 {
			nf[f2] = min(f[f2]+dis[x][y], f[y]+dis[f2][x])
		}
		f, nf = nf, f
	}
	return slices.Min(f[:])
}

package main

import "fmt"

func main() {
	parent := []int{-1, 0, 0, 1, 1, 1}
	s := "abaabc"

	fmt.Println("sub:", findSubtreeSizes(parent, s))
}

func findSubtreeSizes(parent []int, s string) []int {
	n := len(parent)
	g := make([][]int, n)
	for i := 1; i < n; i++ {
		p := parent[i]
		g[p] = append(g[p], i)
	}

	ancestor := [26]int{}
	for i := range ancestor {
		ancestor[i] = -1
	}
	var rebuild func(int)
	rebuild = func(x int) {
		sx := s[x] - 'a'
		old := ancestor[sx]
		ancestor[sx] = x
		for i, y := range g[x] {
			if anc := ancestor[s[y]-'a']; anc != -1 {
				g[anc] = append(g[anc], y)
				g[x][i] = -1 // -1 表示删除 y
			}
			rebuild(y)
		}
		ancestor[sx] = old // 恢复现场
	}
	rebuild(0)

	size := make([]int, n)
	var dfs func(int)
	dfs = func(x int) {
		size[x] = 1
		for _, y := range g[x] {
			if y != -1 { // y 没被删除
				dfs(y)
				size[x] += size[y]
			}
		}
	}
	dfs(0)
	return size
}

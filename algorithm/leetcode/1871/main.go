package main

import "fmt"

func main() {
	s := "01101110"
	minJump := 2
	maxJump := 3
	fmt.Println("can reach:", canReach(s, minJump, maxJump))
}

func canReach(s string, minJump, maxJump int) bool {
	n := len(s)
	f := make([]int, n)
	// f 的前缀和
	sum := make([]int, n+1)
	f[0] = 1
	sum[1] = 1
	for j := 1; j < n; j++ {
		if j >= minJump && s[j] == '0' && sum[j-minJump+1] > sum[max(j-maxJump, 0)] {
			f[j] = 1
		}
		sum[j+1] = sum[j] + f[j]
	}
	return f[n-1] == 1
}

// 在s全为0， minJump到maxJump较大时，会超时
func canReach2(s string, minJump int, maxJump int) bool {
	n := len(s)
	memo := make([]int, n)
	reach := []int{0}
	memo[0] = 1
	for len(reach) > 0 {
		tmp := []int{}
		for _, idx := range reach {
			for i := idx + minJump; i <= min(idx+maxJump, n-1); i++ {
				if s[i] == '1' || memo[i] == 1 {
					continue
				}
				if i == n-1 {
					return true
				}
				tmp = append(tmp, i)
				memo[i] = 1
			}
		}
		reach = tmp
	}
	return false
}

func canReach3(s string, minJump int, maxJump int) bool {
	n := len(s)
	if s[n-1] == '1' {
		return false
	}
	if maxJump == n-1 {
		return true
	}
	memo := make([]bool, n)
	for i, ch := range s {
		if ch == '1' {
			memo[i] = true
		}
	}

	var dfs func(i int) bool
	dfs = func(i int) bool {
		if i == n-1 {
			return true
		}
		if memo[i] {
			return false
		}
		for j := i + minJump; j <= min(n-1, i+maxJump); j++ {
			if dfs(j) {
				return true
			}
			memo[j] = true
		}
		return false
	}

	return dfs(0)
}

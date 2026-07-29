package main

import (
	"bytes"
	"slices"
)

func main() {

}

func smallestPalindrome(s string, k int) string {
	n := len(s)
	m := n / 2

	total := [26]int{}
	for _, b := range s[:m] {
		total[b-'a']++
	}

	cnt := make([]int, 26)
	perm := 1
	i, j := m-1, 25
	// 倒排
	for ; i >= 0 && perm < k; i-- {
		for cnt[j] == total[j] {
			j--
		}
		cnt[j]++
		perm = perm * (m - i) / cnt[j]
	}

	if perm < k {
		return ""
	}

	ans := make([]byte, 0, n) // 预分配空间
	// 已经有足够的排列数了，<= i 的位置直接填字典序最小的排列
	for ch, c := range cnt[:j+1] {
		ans = append(ans, bytes.Repeat([]byte{'a' + byte(ch)}, total[ch]-c)...)
	}

	// 试填法
	j0 := j
	for i++; i < m; i++ {
		for j := j0; j < 26; j++ {
			if cnt[j] == 0 {
				continue
			}
			// 假设填字母 j，根据 perm = p * (m - i) / cnt[j] 倒推 p
			p := perm * cnt[j] / (m - i)
			if p >= k {
				ans = append(ans, 'a'+byte(j))
				cnt[j]--
				perm = p
				break
			}
			k -= p
		}
	}

	rev := slices.Clone(ans)
	if n%2 > 0 {
		ans = append(ans, s[n/2])
	}
	slices.Reverse(rev)
	ans = append(ans, rev...)
	return string(ans)
}

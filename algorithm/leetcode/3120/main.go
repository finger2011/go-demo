package main

import (
	"fmt"
	"math/bits"
)

func main() {
	word := "aaAbcBC"
	fmt.Println("num :", numberOfSpecialChars(word))
	fmt.Println("num2:", numberOfSpecialChars2(word))
}

// 时间复杂度 O(n + |Σ|) Σ = 26， n是字符串长度
// 空间复杂度 O(n)
func numberOfSpecialChars(word string) int {
	s := make(map[rune]bool, 52)
	for _, ch := range word {
		s[ch] = true
	}
	var ans int
	for ch := 'a'; ch <= 'z'; ch++ {
		if s[ch] && s[ch-'a'+'A'] {
			ans++
		}
	}
	return ans
}

// 对于大写英文字母：其二进制从右往左第 6 个比特值一定是 0。

// 对于小写英文字母：其二进制从右往左第 6 个比特值一定是 1。

// 对于任何英文字母：其小写字母二进制低 5 位，一定和其大写字母二进制低 5 位相等。

func numberOfSpecialChars2(word string) int {
	mask := [2]int{}
	for _, c := range word {
		mask[c>>5&1] |= 1 << (c & 31)
	}
	return bits.OnesCount(uint(mask[0] & mask[1]))
}

package main

import "slices"

func main() {

}

func minimumPushes(word string) int {
	var ans int
	for i := range word {
		ans += (i + 8) / 8
	}
	return ans

}

// 进阶：如果有相同字母
func minimumPushes2(word string) int {
	chars := make([]int, 26)
	for _, ch := range word {
		chars[ch-'a']++
	}
	slices.SortFunc(chars, func(a, b int) int {
		return b - a
	})
	var ans int
	for i, num := range chars {
		k := (i + 8) / 8
		ans += k * num
	}
	return ans
}

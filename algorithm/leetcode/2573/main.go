package main

import "fmt"

type ANS struct {
	grid   [][]int
	expect string
}

func main() {
	tests := []*ANS{
		{
			grid:   [][]int{{4, 0, 2, 0}, {0, 3, 0, 1}, {2, 0, 2, 0}, {0, 1, 0, 1}},
			expect: "abab",
		},
	}
	for i, test := range tests {
		fmt.Println("test:", i)
		fmt.Println("lcp:", test.grid)
		fmt.Println("except:", test.expect)
		fmt.Println("result ====================> ", findTheString(test.grid) == test.expect)
	}
}

func findTheString(lcp [][]int) string {
	length := len(lcp)
	words := make([]byte, length)
	ch := byte('a')
	for i := range length {
		if words[i] == 0 {
			if ch > 'z' {
				return ""
			}
			words[i] = ch
			for j := i + 1; j < length; j++ {
				if lcp[i][j] > 0 {
					words[j] = words[i]
				}
			}
			ch++
		}
	}
	for i := length - 1; i >= 0; i-- {
		for j := length - 1; j >= 0; j-- {
			var cal int
			if words[i] == words[j] {
				if i == length-1 || j == length-1 {
					cal = 1
				} else {
					cal = lcp[i+1][j+1] + 1
				}
			}
			if lcp[i][j] != cal {
				return ""
			}
		}
	}
	return string(words)
}

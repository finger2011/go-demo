package main

import "fmt"

type Input struct {
	words      []string
	target     string
	startIndex int
	except     int
}

func main() {
	tests := []*Input{
		{
			words:      []string{"hello", "i", "am", "leetcode", "hello"},
			target:     "hello",
			startIndex: 1,
			except:     1,
		},
	}
	for _, test := range tests {
		fmt.Println("======================start")
		fmt.Println("words:", test.words)
		fmt.Println("target:", test.target)
		fmt.Println("startIndex:", test.startIndex)
		ans := closestTarget(test.words, test.target, test.startIndex)
		fmt.Println("return:", ans)
		fmt.Println("result:", ans == test.except)
		fmt.Println("======================end")
	}
}

func closestTarget(words []string, target string, startIndex int) int {
	n, ans := len(words), -1
	for i := range n/2 + 1 {
		if words[(startIndex+i)%n] == target || words[(startIndex-i+n)%n] == target {
			ans = i
			break
		}
	}
	return ans
}

func closestTarget2(words []string, target string, startIndex int) int {
	n := len(words)
	ans := n
	for i, word := range words {
		if word == target {
			d := i - startIndex
			if d < 0 {
				d = -d
			}
			ans = min(ans, d, n-d)
		}
	}
	if ans == n {
		return -1
	}
	return ans
}

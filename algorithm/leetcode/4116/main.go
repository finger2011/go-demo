package main

import "fmt"

func main() {
	// rain := []int{1, 1, 0, 1, 1}
	rain := []int{1, 1, 1}
	fmt.Println("ans:", assignUmbrella(rain))
}

func assignUmbrella(rain []int) []string {
	n := len(rain)
	ans := make([]string, n)
	numA, numB := 0, 0
	for i, r := range rain {
		switch r {
		case 0:
			ans[i] = "None"
		case 1:
			if i < n-1 && rain[i+1] == 1 {
				if numA <= numB {
					ans[i] = "Alice"
					numA++
				} else {
					ans[i] = "Bob"
					numB++
				}
			} else {
				ans[i] = "Bob"
				numB++
			}
		}
	}
	return ans
}

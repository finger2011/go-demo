package main

import "fmt"

func main() {
	n := "27346209830709182346"
	fmt.Println("minPartitions(", n, ") ===> ", minPartitions(n))
}

// 针对每一位i，最多int(n[i])次，只需要寻找最大的n[i]即可
func minPartitions(n string) int {
	var ans int
	for _, v := range n {
		t := int(v - '0')
		if t > ans {
			ans = t
		}
	}
	return ans
}

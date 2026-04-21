package main

import "fmt"

func main() {
	// "word","note","wood"
	queries := []string{"word", "note", "ants", "wood"}
	dictionary := []string{"wood", "joke", "moat"}
	fmt.Println("twoEditWords:", twoEditWords(queries, dictionary))
}

// 暴力
func twoEditWords(queries []string, dictionary []string) []string {
	var ans []string
	for _, q := range queries {
		for _, d := range dictionary {
			var invalid bool
			var ant int
			for i := range q {
				if q[i] != d[i] {
					ant++
					if ant > 2 {
						invalid = true
						break
					}
				}
			}
			if !invalid {
				ans = append(ans, q)
				break
			}
		}
	}
	return ans
}

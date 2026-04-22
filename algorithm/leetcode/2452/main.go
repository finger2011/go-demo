package main

import "fmt"

func main() {
	// "word","note","wood"
	// queries := []string{"word", "note", "ants", "wood"}
	// dictionary := []string{"wood", "joke", "moat"}

	// "afcb"
	queries := []string{"afcb", "covd", "npfq"}
	dictionary := []string{"knfl", "ahcb", "kxvs"}
	fmt.Println("twoEditWords:", twoEditWords(queries, dictionary))
	fmt.Println("twoEditWords2:", twoEditWords2(queries, dictionary))
}

// 暴力 时间: O(qkn), q=queries长度 k=dictionary长度 n=单词长度； 空间:O(1)
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

type dic struct {
	child [26]*dic
	isEnd bool
}

func (d *dic) insert(word string) {
	node := d
	for _, c := range word {
		idx := c - 'a'
		if node.child[idx] == nil {
			node.child[idx] = &dic{}
		}
		node = node.child[idx]
	}
	node.isEnd = true
}

// 字典树
// 时间:  k= dictionary长度 n=单词长度 q = queries长度
// 1. 建字典树需要 O(kn)
// 2. 查询对每一个字母有不修改1种，修改最多26种, 最多修改两次，q*n*n*25*25
// 所以总的时间复杂度为O(kn+q*n*n*25*25)
// 空间：O(kn) 字典树空间为kn，
func twoEditWords2(queries []string, dictionary []string) []string {
	var ans []string
	var root = &dic{}
	var dfs func(word string, i, cnt int, node *dic) bool
	dfs = func(word string, i, cnt int, node *dic) bool {
		if cnt > 2 || node == nil {
			return false
		}
		if i >= len(word) {
			return node.isEnd
		}
		idx := word[i] - 'a'
		for j, n := range node.child {
			if idx == byte(j) && dfs(word, i+1, cnt, n) {
				return true
			}
			if cnt < 2 && dfs(word, i+1, cnt+1, n) {
				return true
			}
		}
		return false
	}
	for _, d := range dictionary {
		root.insert(d)
	}
	for _, q := range queries {
		if dfs(q, 0, 0, root) {
			ans = append(ans, q)
		}
	}
	return ans
}

package main

import (
	"fmt"
)

type Input struct {
	source, target []int
	allowedSwaps   [][]int
	except         int
}

func main() {
	tests := []*Input{
		{
			source:       []int{1, 2, 3, 4},
			target:       []int{2, 1, 4, 5},
			allowedSwaps: [][]int{},
			except:       4,
		},
	}
	for _, test := range tests {
		fmt.Println("============start")
		fmt.Println("source:", test.source)
		fmt.Println("target:", test.target)
		fmt.Println("allowedSwaps:", test.allowedSwaps)
		ans := minimumHammingDistance(test.source, test.target, test.allowedSwaps)
		fmt.Println("return:", ans)
		fmt.Println("result:", test.except == ans)
		fmt.Println("============end")
	}
}

//  1. 按照allowedSwaps 分组，有下标相同的分在同一组，如[1,2], [2,3] => [1,2,3]
//  2. 对于每个分组，比较source 和 target 中对应下标组成的数组,不同元素的个数
//     如[1,2] => [source[1], source[2]], [targe[1], target[2]]
//  3. 每个分组不同元素的和即为最终答案
// 需要用到并查集
// TODO
// func minimumHammingDistance2(source []int, target []int, allowedSwaps [][]int) int {
// 	n := len(source)
// }

// 使用无向图
func minimumHammingDistance(source []int, target []int, allowedSwaps [][]int) int {
	n := len(source)
	group := make([][]int, n)
	for _, swap := range allowedSwaps {
		group[swap[0]] = append(group[swap[0]], swap[1])
		group[swap[1]] = append(group[swap[1]], swap[0])
	}
	nums := map[int]int{}
	pos := make([]bool, n)
	var dfs func(int)
	dfs = func(i int) {
		pos[i] = true
		nums[source[i]]++
		nums[target[i]]--
		for _, y := range group[i] {
			if !pos[y] {
				dfs(y)
			}
		}
	}
	var ans int
	for i, p := range pos {
		if !p {
			clear(nums)
			dfs(i)
			for _, n := range nums {
				if n > 0 {
					ans += n
				} else if n < 0 {
					ans -= n
				}
			}
		}
	}
	return ans / 2
}

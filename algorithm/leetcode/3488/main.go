package main

import (
	"fmt"
	"slices"
)

type Input struct {
	nums, queries, except []int
}

func main() {
	tests := []*Input{
		{
			nums:    []int{14, 14, 4, 2, 19, 19, 14, 19, 14},
			queries: []int{2, 4, 8, 6, 3},
			except:  []int{-1, 1, 1, 2, -1},
		},

		// {
		// 	nums:    []int{15, 1, 10, 1, 20, 4, 6, 14, 4, 9, 4, 18},
		// 	queries: []int{0, 2, 10, 6, 11, 8},
		// 	except:  []int{-1, -1, 2, -1, -1, 2},
		// },
		// {
		// 	nums:    []int{1, 3, 1, 4, 1, 3, 2},
		// 	queries: []int{0, 3, 5},
		// 	except:  []int{2, -1, 3},
		// },
	}
	for _, test := range tests {
		fmt.Println("=========start")
		fmt.Println("nums:", test.nums)
		fmt.Println("queries:", test.queries)
		ans := solveQueries(test.nums, test.queries)
		fmt.Println("return:", ans)
		fmt.Println("result:", slices.Equal(ans, test.except))
		fmt.Println("=========end")
	}
}

func solveQueries(nums []int, queries []int) []int {
	n := len(nums)
	ans := make([]int, n)
	for i := range ans {
		ans[i] = -1
	}
	numMap := make(map[int][2]int, n)
	for i, num := range nums {
		pos, exist := numMap[num]
		if !exist {
			numMap[num] = [2]int{i, i}
			continue
		}
		dis1, dis2 := i-pos[1], pos[0]+n-i
		if ans[pos[1]] == -1 || ans[pos[1]] > dis1 {
			ans[pos[1]] = dis1
		}
		if ans[pos[0]] == -1 || ans[pos[0]] > dis2 {
			ans[pos[0]] = dis2
		}
		ans[i] = min(dis1, dis2)
		numMap[num] = [2]int{pos[0], i}
	}
	qAns := make([]int, len(queries))
	for i, pos := range queries {
		qAns[i] = ans[pos]
	}
	return qAns
}

package main

import (
	"fmt"
	"slices"
)

type Input struct {
	startPos, homePos, rowCosts, colCosts []int
	expect                                int
}

func main() {
	tests := []*Input{
		{
			startPos: []int{1, 0},
			homePos:  []int{2, 3},
			rowCosts: []int{5, 4, 3},
			colCosts: []int{8, 2, 6, 7},
			expect:   18,
		},
	}
	for _, test := range tests {
		fmt.Println("================start")
		fmt.Println("startPos:", test.startPos)
		fmt.Println("homePos:", test.homePos)
		fmt.Println("rowCosts:", test.rowCosts)
		fmt.Println("colCosts:", test.colCosts)
		fmt.Println("expect:", test.expect)
		ans := minCost(test.startPos, test.homePos, test.rowCosts, test.colCosts)
		fmt.Println("return :", ans)
		fmt.Println("result:", test.expect == ans)
		fmt.Println("================end")
	}
}

func minCost(startPos []int, homePos []int, rowCosts []int, colCosts []int) int {
	if slices.Equal(startPos, homePos) {
		return 0
	}
	var ans, rowPath, colPath int
	if startPos[0] > homePos[0] {
		rowPath = -1
	} else if startPos[0] < homePos[0] {
		rowPath = 1
	}
	if startPos[1] > homePos[1] {
		colPath = -1
	} else if startPos[1] < homePos[1] {
		colPath = 1
	}
	if rowPath != 0 {
		for i := startPos[0] + rowPath; ; i = i + rowPath {
			ans += rowCosts[i]
			if i == homePos[0] {
				break
			}
		}
	}
	if colPath != 0 {
		for i := startPos[1] + colPath; ; i = i + colPath {
			ans += colCosts[i]
			if i == homePos[1] {
				break
			}
		}
	}

	return ans
}

package main

import (
	"fmt"
	"math"
	"slices"
)

type Input struct {
	robots, distance, walls []int
	except                  int
}

func main() {
	tests := []*Input{
		{
			robots:   []int{17, 59, 32, 11, 72, 18},
			distance: []int{5, 7, 6, 5, 2, 10},
			walls:    []int{17, 25, 33, 29, 54, 53, 18, 35, 39, 37, 20, 14, 34, 13, 16, 58, 22, 51, 56, 27, 10, 15, 12, 23, 45, 43, 21, 2, 42, 7, 32, 40, 8, 9, 1, 5, 55, 30, 38, 4, 3, 31, 36, 41, 57, 28, 11, 49, 26, 19, 50, 52, 6, 47, 46, 44, 24, 48},
			except:   37,
		},
		{
			robots:   []int{10, 2},
			distance: []int{5, 1},
			walls:    []int{5, 2, 7},
			except:   3,
		},
	}
	for _, test := range tests {
		fmt.Println("start")
		fmt.Println("robots:", test.robots)
		fmt.Println("distance:", test.distance)
		fmt.Println("walls:", test.walls)
		fmt.Println("except:", test.except)
		ans := maxWalls(test.robots, test.distance, test.walls)
		fmt.Println("return:", ans)
		fmt.Println("result =======> ", test.except == ans)
		fmt.Println("start")
	}
}

func maxWalls(robots []int, distance []int, walls []int) int {
	n, m := len(robots), len(walls)
	type pair struct{ x, d int }
	a := make([]pair, n+2)
	for i, x := range robots {
		a[i] = pair{x, distance[i]}
	}
	a[n+1].x = math.MaxInt // 哨兵
	slices.SortFunc(a, func(a, b pair) int { return a.x - b.x })
	slices.Sort(walls)

	var f0, f1, left, cur, right0, right1 int
	for i := 1; i <= n; i++ {
		p := a[i]

		// 往左射，墙的坐标范围为 [leftX, p.x]
		leftX := max(p.x-p.d, a[i-1].x+1) // +1 表示不能射到左边那个机器人
		for left < m && walls[left] < leftX {
			left++
		}
		for cur < m && walls[cur] < p.x {
			cur++
		}
		cur1 := cur
		if cur < m && walls[cur] == p.x {
			cur++
		}
		leftRes := f0 + cur - left // 下标在 [left, cur-1] 中的墙都能摧毁

		// 往右射，右边那个机器人往左射，墙的坐标范围为 [p.x, rightX]
		q := a[i+1]
		rightX := min(p.x+p.d, q.x-q.d-1) // -1 表示不能射到右边那个机器人（或者它往左射到的墙）
		for right0 < m && walls[right0] <= rightX {
			right0++
		}
		f0 = max(leftRes, f1+right0-cur1) // 下标在 [cur1, right0-1] 中的墙都能摧毁

		// 往右射，右边那个机器人往右射，墙的坐标范围为 [p.x, rightX]
		rightX = min(p.x+p.d, q.x-1) // -1 表示不能射到右边那个机器人（或者它往左射到的墙）
		for right1 < m && walls[right1] <= rightX {
			right1++
		}
		f1 = max(leftRes, f1+right1-cur1) // 下标在 [cur1, right0-1] 中的墙都能摧毁
	}
	return f1
}

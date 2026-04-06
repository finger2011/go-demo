package main

import "fmt"

type Input struct {
	commands  []int
	obstacles [][]int
	except    int
}

func main() {
	tests := []*Input{
		{
			commands:  []int{4, -1, 3},
			obstacles: [][]int{},
			except:    25,
		},
	}
	for _, test := range tests {
		fmt.Println("======================start")
		fmt.Println("commands:", test.commands)
		fmt.Println("obstacles:", test.obstacles)
		ans := robotSim(test.commands, test.obstacles)
		fmt.Println("return:", ans)
		fmt.Println("result:", ans == test.except)
		fmt.Println("======================end")
	}
}

func robotSim(commands []int, obstacles [][]int) int {
	var ans int
	direct := [4][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	x, y, d := 0, 0, 0
	const n = 60001
	obstacleMap := make(map[int]bool, len(obstacles))
	for _, ob := range obstacles {
		obstacleMap[ob[0]*n+ob[1]] = true
	}
	for _, command := range commands {
		if command == -2 {
			d = (d + 4 - 1) % 4
			continue
		}
		if command == -1 {
			d = (d + 1) % 4
			continue
		}
		for range command {
			if obstacleMap[(x+direct[d][0])*n+y+direct[d][1]] {
				break
			}
			x += direct[d][0]
			y += direct[d][1]

		}
		ans = max(ans, x*x+y*y)
	}
	return ans
}

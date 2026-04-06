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
	x, y := 0, 0
	direct := 0
	type pair struct{ x, y int }
	obstacleMap := make(map[pair]bool, len(obstacles))
	for _, ob := range obstacles {
		obstacleMap[pair{ob[0], ob[1]}] = true
	}
	for _, command := range commands {
		if command == -2 {
			direct = (direct + 4 - 1) % 4
			continue
		}
		if command == -1 {
			direct = (direct + 1) % 4
			continue
		}
		var op int
		switch direct {
		case 0:
			for op = range command {
				if obstacleMap[pair{x, y + op}] {

				}
			}
			y += command
		case 1:
			x += command
		case 2:
			y -= command
		case 3:
			x -= command
		}
		if x*x+y*y > ans {
			ans = x*x + y*y
		}
	}
	return ans
}

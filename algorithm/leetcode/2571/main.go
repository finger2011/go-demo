package main

import (
	"fmt"
	"slices"
)

type Input struct {
	positions, healths, except []int
	directions                 string
}

func main() {
	tests := []*Input{
		{
			positions:  []int{17, 24, 18},
			healths:    []int{1, 39, 30},
			directions: "LLR",
			except:     []int{1, 38},
		},
		{
			positions:  []int{1, 40},
			healths:    []int{10, 11},
			directions: "RL",
			except:     []int{10},
		},
		{
			positions:  []int{1, 2, 5, 6},
			healths:    []int{10, 10, 11, 11},
			directions: "RLRL",
			except:     []int{},
		},
		{
			positions:  []int{3, 5, 2, 6},
			healths:    []int{10, 10, 15, 12},
			directions: "RLRL",
			except:     []int{14},
		},
		{
			positions:  []int{5, 4, 3, 2, 1},
			healths:    []int{2, 17, 9, 15, 10},
			directions: "RRRRR",
			except:     []int{2, 17, 9, 15, 10},
		},
	}
	for _, test := range tests {
		fmt.Println("start")
		fmt.Println("positions:", test.positions)
		fmt.Println("healths:", test.healths)
		fmt.Println("directions:", test.directions)
		fmt.Println("except:", test.except)
		res := survivedRobotsHealths(test.positions, test.healths, test.directions)
		fmt.Println("return:", res)
		fmt.Println("result ======> ", slices.Equal(test.except, res))
		fmt.Println("end")
	}
}

func survivedRobotsHealths(positions []int, healths []int, directions string) []int {
	n := len(positions)
	idx := make([]int, n)
	for i := range idx {
		idx[i] = i
	}
	slices.SortFunc(idx, func(a, b int) int {
		return positions[a] - positions[b]
	})
	robotPos := []int{}
	for _, i := range idx {
		if directions[i] == 'R' {
			robotPos = append(robotPos, i)
			continue
		}
		for len(robotPos) > 0 {
			prev := robotPos[len(robotPos)-1]
			if healths[i] == healths[prev] {
				healths[i] = 0
				healths[prev] = 0
				robotPos = robotPos[:len(robotPos)-1]
				break
			} else if healths[i] < healths[prev] {
				healths[i] = 0
				healths[prev]--
				break
			} else {
				healths[i]--
				healths[prev] = 0
				robotPos = robotPos[:len(robotPos)-1]
			}
		}
	}

	var ans []int
	for _, health := range healths {
		if health > 0 {
			ans = append(ans, health)
		}
	}

	return ans
}

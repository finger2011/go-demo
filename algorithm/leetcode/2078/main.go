package main

import "fmt"

type Input struct {
	colors []int
	except int
}

func main() {
	tests := []*Input{
		{
			colors: []int{6, 6, 6, 6, 6, 6, 6, 6, 6, 19, 19, 6, 6},
			except: 10,
		},
		{
			colors: []int{1, 8},
			except: 1,
		},
		{
			colors: []int{1, 8, 3, 8, 3},
			except: 4,
		},
		{
			colors: []int{1, 1, 1, 6, 1, 1, 1},
			except: 3,
		},
	}
	for _, test := range tests {
		fmt.Println("============start")
		fmt.Println("colors:", test.colors)
		ans := maxDistance(test.colors)
		fmt.Println("return:", ans)
		fmt.Println("result:", test.except == ans)
		fmt.Println("============end")
	}
}

func maxDistance(colors []int) int {
	n := len(colors)
	i, j, ans := 0, n-1, -1
	for i >= 0 && i < n && j >= 0 && i <= j {
		if colors[i] != colors[n-1] {
			ans = max(ans, n-1-i)
			break
		}
		if colors[j] != colors[0] {
			ans = max(ans, j)
			break
		}
		i++
		j--
	}
	return ans
}

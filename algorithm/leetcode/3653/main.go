package main

import "fmt"

type Input struct {
	nums    []int
	queries [][]int
	except  int
}

const mode = 1000000007

func main() {
	tests := []*Input{
		{
			nums:    []int{4, -1, 3},
			queries: [][]int{},
			except:  25,
		},
	}
	for _, test := range tests {
		fmt.Println("======================start")
		fmt.Println("nums:", test.nums)
		fmt.Println("queries:", test.queries)
		ans := xorAfterQueries(test.nums, test.queries)
		fmt.Println("return:", ans)
		fmt.Println("result:", ans == test.except)
		fmt.Println("======================end")
	}
}

func xorAfterQueries(nums []int, queries [][]int) int {

	for _, query := range queries {
		for i := query[0]; i <= query[1]; i += query[2] {
			nums[i] = (nums[i] * query[3]) % mode
		}
	}
	var ans int
	for _, num := range nums {
		ans ^= num
	}
	return ans
}

package main

import (
	"fmt"
	"math"
)

type Input struct {
	nums   []int
	except int
}

func main() {
	// revertNum(1001)
	// return
	tests := []*Input{
		{
			nums:   []int{3794, 6634, 9999, 7118, 6283, 9057, 2086, 8101, 210, 7658, 5917, 7446, 9322, 6694, 937, 7297, 9341, 1927, 9687, 5402, 7150, 2826, 2619, 5613, 916, 5356, 7111, 3184, 8783, 3064, 635, 6092, 6965, 662, 4705, 1749, 7386, 4274, 6779, 4469, 3527, 9026, 1250, 6929, 838, 9500, 7882, 7488, 186, 9788, 3248, 4182, 4736, 9616, 3977, 9094, 9371, 937, 4308, 2420, 6191, 1154, 6862, 2102, 3439, 1001, 7180, 7722, 1135, 6028, 4220, 8259, 3080, 5632, 7923, 6014, 3631, 2204, 5391, 2631, 6689, 9964, 1490, 9479, 1662, 2605, 483, 454, 106, 1851, 724, 1591, 8580, 1001, 6721, 2294, 6803, 9609, 8306},
			except: 28,
		},
		// {
		// 	nums:   []int{21, 120},
		// 	except: -1,
		// },
		// {
		// 	nums:   []int{120, 21},
		// 	except: 1,
		// },
		// {
		// 	nums:   []int{12, 21, 45, 33, 54},
		// 	except: 1,
		// },
	}
	for _, test := range tests {
		fmt.Println("============start")
		fmt.Println("nums:", test.nums)
		ans := minMirrorPairDistance(test.nums)
		fmt.Println("return:", ans)
		fmt.Println("result:", test.except == ans)
		fmt.Println("============end")
	}
}

func minMirrorPairDistance(nums []int) int {
	ans := math.MaxInt
	numMap := make(map[int]int, len(nums))
	for i := len(nums) - 1; i >= 0; i-- {
		revertNum := revertNum(nums[i])
		if pos, exist := numMap[revertNum]; exist {
			fmt.Println("accept")
			ans = min(ans, pos-i)
		}
		numMap[nums[i]] = i
	}
	if ans == math.MaxInt {
		return -1
	}
	return ans
}

func revertNum(num int) int {
	var ans int
	for num >= 10 {
		ans = ans*10 + num%10
		num /= 10
	}
	ans = ans*10 + num
	return ans
}

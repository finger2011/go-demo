package main

import "math"

func main() {

}

func earliestFinishTime2(landStartTime []int, landDuration []int,
	waterStartTime []int, waterDuration []int) int {

	ans, n, m := math.MaxInt, len(landStartTime), len(waterStartTime)
	for i := range n {
		for j := range m {
			tmp := landStartTime[i] + landDuration[i]
			if waterStartTime[j] >= tmp {
				tmp = waterStartTime[j] + waterDuration[j]
			} else {
				tmp += waterDuration[j]
			}
			ans = min(ans, tmp)
		}
	}
	for j := range m {
		for i := range n {
			tmp := waterStartTime[j] + waterDuration[j]
			if landStartTime[i] >= tmp {
				tmp = landStartTime[i] + landDuration[i]
			} else {
				tmp += landDuration[i]
			}
			ans = min(ans, tmp)
		}
	}
	return ans
}

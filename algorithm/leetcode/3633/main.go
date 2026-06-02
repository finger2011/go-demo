package main

import "math"

func main() {

}

func earliestFinishTime(landStartTime []int, landDuration []int,
	waterStartTime []int, waterDuration []int) int {

	return min(finish(landStartTime, landDuration, waterStartTime, waterDuration),
		finish(waterStartTime, waterDuration, landStartTime, landDuration))
}

func finish(landStartTime []int, landDuration []int,
	waterStartTime []int, waterDuration []int) int {
	minFinish, ans := math.MaxInt, math.MaxInt
	for i := range landStartTime {
		minFinish = min(minFinish, landStartTime[i]+landDuration[i])
	}
	for i := range waterStartTime {
		ans = min(ans, max(minFinish, waterStartTime[i])+waterDuration[i])
	}
	return ans
}

// 暴力双重遍历 n*m
func earliestFinishTime2(landStartTime []int, landDuration []int,
	waterStartTime []int, waterDuration []int) int {

	ans, n, m := math.MaxInt, len(landStartTime), len(waterStartTime)
	for i := range n {
		for j := range m {
			tmp := landStartTime[i] + landDuration[i]
			if tmp >= ans || tmp+waterDuration[j] >= ans {
				continue
			}
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
			if tmp >= ans || tmp+landDuration[i] >= ans {
				continue
			}
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

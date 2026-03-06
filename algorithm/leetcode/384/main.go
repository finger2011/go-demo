package main

import (
	"fmt"
	"math/rand"
)

func main() {
	nums := []int{1, 2, 3}
	obj := Constructor(nums)
	fmt.Println("Reset:", obj.Reset())
	fmt.Println("Shuffle:", obj.Shuffle())
}

type Solution struct {
	nums, origin []int
}

func Constructor(nums []int) Solution {
	return Solution{nums, append([]int(nil), nums...)}
}

func (this *Solution) Reset() []int {
	copy(this.nums, this.origin)
	return this.nums
}

func (this *Solution) Shuffle() []int {
	n := len(this.nums)
	for i := range this.nums {
		j := i + rand.Intn(n-i)
		this.nums[i], this.nums[j] = this.nums[j], this.nums[i]
	}
	return this.nums
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Reset();
 * param_2 := obj.Shuffle();
 */

package main

import "fmt"

func main() {
	ops := []string{"Fancy", "append", "addAll", "append", "multAll", "getIndex", "addAll", "append", "multAll", "getIndex", "getIndex", "getIndex"}
	opsNums := []int{0, 2, 3, 7, 2, 0, 3, 10, 2, 0, 1, 2}
	var obj Fancy
	for i := 0; i < len(ops); i++ {
		switch ops[i] {
		case "Fancy":
			obj = Constructor()
			fmt.Println("Fancy Constructor")
		case "append":
			obj.Append(opsNums[i])
			fmt.Println("Append", opsNums[i])
		case "addAll":
			obj.AddAll(opsNums[i])
			fmt.Println("addAll", opsNums[i])
		case "multAll":
			obj.MultAll(opsNums[i])
			fmt.Println("multAll", opsNums[i])
		case "getIndex":
			fmt.Println("GetIndex(", opsNums[i], "):", obj.GetIndex(opsNums[i]))
		}
	}
	var obj2 Fancy2
	for i := 0; i < len(ops); i++ {
		switch ops[i] {
		case "Fancy":
			obj2 = Constructor2()
			fmt.Println("Fancy2 Constructor")
		case "append":
			obj2.Append(opsNums[i])
			fmt.Println("Append", opsNums[i])
		case "addAll":
			obj2.AddAll(opsNums[i])
			fmt.Println("addAll", opsNums[i])
		case "multAll":
			obj2.MultAll(opsNums[i])
			fmt.Println("multAll", opsNums[i])
		case "getIndex":
			fmt.Println("GetIndex(", opsNums[i], "):", obj2.GetIndex(opsNums[i]))
		}
	}
}

type Fancy struct {
	count  int
	nums   [][2]int // num, startOps
	ops    []string
	opsNum []int
}

const mod = 1000000007

func Constructor() Fancy {
	return Fancy{}
}

func (this *Fancy) Append(val int) {
	this.count++
	this.nums = append(this.nums, [2]int{val, len(this.ops)})
}

func (this *Fancy) AddAll(inc int) {
	this.ops = append(this.ops, "+")
	this.opsNum = append(this.opsNum, inc)
}

func (this *Fancy) MultAll(m int) {
	this.ops = append(this.ops, "*")
	this.opsNum = append(this.opsNum, m)
}

func (this *Fancy) GetIndex(idx int) int {
	if idx >= this.count {
		return -1
	}
	opsLen := len(this.ops)
	if this.nums[idx][1] >= opsLen {
		return this.nums[idx][0]
	}
	ans := this.nums[idx][0]
	for i := this.nums[idx][1]; i < opsLen; i++ {
		switch this.ops[i] {
		case "+":
			ans = int((int64(ans) + int64(this.opsNum[i])) % mod)
		case "*":
			ans = int((int64(ans) * int64(this.opsNum[i])) % mod)
		}
	}
	this.nums[idx] = [2]int{ans, opsLen}
	return ans
}

// ---------------------------------------------------------------------//

// 通过保存add和mul，每次getindex时进行计算
// add: add+inc
// mul : add*m, mul*m
// 同时防止溢出需要取模
// append的时候先对val进行反运算，即(val - add) / mul,把除法转化为乘法（乘以mod的乘法逆元）
// getindex的时候 (val+add) * mul即可

type Fancy2 struct {
	nums []int
	add  int
	mul  int
}

func Constructor2() Fancy2 {
	return Fancy2{mul: 1}
}

func (this *Fancy2) Append(val int) {
	this.nums = append(this.nums, (val-this.add+mod)*pow(this.mul, mod-2)%mod)
}

func (this *Fancy2) AddAll(inc int) {
	this.add = (this.add + inc) % mod
}

func (this *Fancy2) MultAll(m int) {
	this.add = (this.add * m) % mod
	this.mul = this.mul * m % mod
}

func (this *Fancy2) GetIndex(idx int) int {
	if idx >= len(this.nums) {
		return -1
	}
	return (this.nums[idx]*this.mul + this.add) % mod
}

func pow(x, y int) int {
	res := 1
	for ; y > 0; y /= 2 {
		if y%2 > 0 {
			res = res * x % mod
		}
		x = x * x % mod
	}
	return res
}

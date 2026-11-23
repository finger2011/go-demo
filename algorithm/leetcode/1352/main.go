package main

import "fmt"

func main() {
	opts := []string{"ProductOfNumbers", "add", "add", "add", "add", "add", "getProduct", "getProduct", "getProduct", "add", "getProduct"}
	vars := [][]int{{}, {3}, {0}, {2}, {5}, {4}, {2}, {3}, {4}, {8}, {2}}
	var obj ProductOfNumbers
	for i, opt := range opts {
		switch opt {
		case "ProductOfNumbers":
			obj = Constructor()
		case "add":
			obj.Add(vars[i][0])
		case "getProduct":
			fmt.Println("get product:", obj.GetProduct(vars[i][0]))
		}
	}
}

type ProductOfNumbers []int

func Constructor() ProductOfNumbers {
	return []int{1}
}

func (this *ProductOfNumbers) Add(num int) {
	if num == 0 {
		*this = []int{1}
	} else {
		*this = append(*this, (*this)[len(*this)-1]*num)
	}
}

func (this *ProductOfNumbers) GetProduct(k int) int {
	n := len(*this)
	if n <= k {
		return 0
	}
	return (*this)[n-1] / (*this)[n-k-1]
}

/**
 * Your ProductOfNumbers object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(num);
 * param_2 := obj.GetProduct(k);
 */

package main

import (
	"fmt"
)

func main()  {
	fmt.Println("start:")
	var order = []int{1,3,4,2}
	fmt.Printf("result:%v", verifyPreorder(order))
}
package main

import (
	"fmt"
	"sync"
)

func main() {
	var (
		wg    = &sync.WaitGroup{}
		count int
	)
	wg.Add(100)
	for i := 1; i <= 100; i++ {

		go func() {
			defer wg.Done()
			count += i
			fmt.Println("add:", i)
		}()
	}
	wg.Wait()
	fmt.Println(count)
}

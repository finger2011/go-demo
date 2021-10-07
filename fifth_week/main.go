package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	var rate, _ = newRate(10)
	var wg sync.WaitGroup
	rate.Start()
	var start = time.Now().UnixNano()
	var errNum int
	for i := 0; i < 2100; i++ {
		wg.Add(1)
		go func() {
			var j int
			time.Sleep(time.Duration(rand.Intn(800)) * time.Millisecond)
			for j < 10 {
				err := rate.AddValue(1)
				if err != nil {
					errNum++
					fmt.Printf("err:%v\n", err)
				}
				j++
			}
			// fmt.Printf("j done:10\n")
			wg.Done()
		}()
	}
	wg.Wait()
	time.Sleep(1 * time.Millisecond)
	rate.PrintValues(start)
	fmt.Printf("err num:%v\n", errNum)
}

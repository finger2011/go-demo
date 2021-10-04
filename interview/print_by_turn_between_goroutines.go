package main

import (
	"fmt"

	"golang.org/x/sync/errgroup"
)

var result = "12AB34CD56EF78GH910IJ1112KL1314MN1516OP1718QR1920ST2122UV2324WX2526YZ2728"
var runStr string

// 使用两个 goroutine 交替打印序列，一个 goroutine 打印数字， 另外一 个 goroutine 打印字母， 最终效果如下:
// 12AB34CD56EF78GH910IJ1112KL1314MN1516OP1718QR1920ST2122UV2324WX2526YZ2728
func printByTurnBetweenGos() {
	var ch1, ch2 = make(chan bool, 1), make(chan bool, 1)
	var str = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	var g errgroup.Group
	defer func() {
		fmt.Println("===========>end")
		fmt.Printf("result:%v\n", result == runStr)
		close(ch1)
		close(ch2)
	}()
	g.Go(func() error {
		return printStr(ch1, ch2, str)
	})
	g.Go(func() error {
		return printInt(ch2, ch1)
	})
	fmt.Println("===========>start")
	ch2 <- true
	if err := g.Wait(); err != nil {
		fmt.Printf("\nwait err:%v\n", err)
	} else {
		fmt.Println("\nwait done")
	}
}

func printStr(ch, ch1 chan bool, str string) error {
	var i = 0
	for {
		select {
		case d := <-ch:
			if d {
				if i <= len(str)-1 {
					fmt.Printf(string(str[i]))
					runStr += string(str[i])
					i++
				}
				if i <= len(str)-1 {
					fmt.Printf(string(str[i]))
					runStr += string(str[i])
					i++
				}
				ch1 <- true
				if i >= len(str) {
					return nil
				}
			}
		}
	}
}

func printInt(ch, ch1 chan bool) error {
	var i = 1
	for {
		select {
		case d := <-ch:
			if d {
				fmt.Printf("%d", i)
				fmt.Printf("%d", i+1)
				runStr += fmt.Sprintf("%d", i) + fmt.Sprintf("%d", i+1)
				i = i + 2
				ch1 <- true
			}
			if i > 28 {
				return nil
			}

		}
	}
}

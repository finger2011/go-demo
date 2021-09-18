package web

import (
	"fmt"
	"syscall"
	"os/signal"
	"os"
)

//ShutDownPrepare 关闭server
func ShutDownPrepare()  {
	
}

//ShutDown shut down
func ShutDown()  {
	ch := make(chan os.Signal)
	//监听指定信号
	signal.Notify(ch, syscall.SIGKILL, syscall.SIGINT)

	go func(){
		for s:= range ch {
			switch s {
			case syscall.SIGKILL, syscall.SIGINT:
				fmt.Println("Program Exit...")
				GracefllExit()
			}
		}
	}()
}

//GracefllExit gracefll exit
func GracefllExit()  {
	fmt.Println("Start Exit")
	fmt.Println("Clean...")
	fmt.Println("End Exit")
	os.Exit(0)
}
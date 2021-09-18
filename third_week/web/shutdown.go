package web

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

//ShutDownPrepare 关闭server
//步骤
//1 下线服务
//2 拒绝请求
//3 请求执行完毕
//4 关闭进程
//5 超时直接关闭
func ShutDownPrepare() {

}

//ShutDown shut down
func ShutDown() {
	ch := make(chan os.Signal)
	//监听指定信号
	signal.Notify(ch, syscall.SIGKILL, syscall.SIGINT)

	go func() {
		for s := range ch {
			switch s {
			case syscall.SIGKILL, syscall.SIGINT:
				fmt.Printf("get %s signal, Program Exit...\n", s)
				//超时控制，超过1min未成功关闭，强制退出
				time.AfterFunc(time.Minute, func() {
					fmt.Println("shut down timeout...")
					os.Exit(-1)
				})
				GracefllExit()
			}
		}
	}()
}

//GracefllExit gracefll exit
func GracefllExit() {
	fmt.Println("Start Exit")
	_, cancel := context.WithTimeout(context.Background(), time.Second*30)
	fmt.Println("Clean...")
	cancel()
	fmt.Println("End Exit")
	os.Exit(0)
}

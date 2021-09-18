package web

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

//步骤
//1 下线服务
//2 拒绝请求
//3 请求执行完毕
//4 关闭进程
//5 超时直接关闭

// type ShutDown struct {

// }

//ShutDownPrepare 关闭server
func ShutDownPrepare() {

}

//DownService 下线服务
func DownService() error {
	fmt.Println("mock down service...")
	time.Sleep(time.Second)
	return nil
}

//RejectNewRequest 拒绝新请求
func RejectNewRequest() error {
	fmt.Println("mock reject new request...")
	time.Sleep(time.Second)
	return nil
}

//DoneAllRequest 请求执行完毕
func DoneAllRequest() error {
	fmt.Println("mock done all request...")
	time.Sleep(time.Second)
	return nil
}

//ShutDown shut down
func ShutDown() {
	ch := make(chan os.Signal, 1)
	//监听指定信号
	signal.Notify(ch, syscall.SIGKILL, syscall.SIGINT)
	select {
	case s := <-ch:
		fmt.Printf("get %s signal, Program Exit...\n", s)
		//超时控制，超过1min未成功关闭，强制退出
		time.AfterFunc(time.Minute, func() {
			fmt.Println("shut down timeout...")
			os.Exit(-1)
		})
		GracefllExit()
	}
	// go func() {
	// 	for s := range ch {
	// 		switch s {
	// 		case syscall.SIGKILL, syscall.SIGINT:
	// 			fmt.Printf("get %s signal, Program Exit...\n", s)
	// 			//超时控制，超过1min未成功关闭，强制退出
	// 			time.AfterFunc(time.Minute, func() {
	// 				fmt.Println("shut down timeout...")
	// 				os.Exit(-1)
	// 			})
	// 			GracefllExit()
	// 		}
	// 	}
	// }()
}

//GracefllExit gracefll exit
func GracefllExit() {
	fmt.Println("Start Exit")
	_, cancel := context.WithTimeout(context.Background(), time.Second*30)
	err := DownService()
	if err != nil {
		fmt.Printf("down service failed:%v", err)
	}
	err = RejectNewRequest()
	if err != nil {
		fmt.Printf("reject new request failed:%v", err)
	}
	err = DoneAllRequest()
	if err != nil {
		fmt.Printf("done all request failed:%v", err)
	}
	fmt.Println("Clean...")
	cancel()
	fmt.Println("End Exit")
	os.Exit(0)
}

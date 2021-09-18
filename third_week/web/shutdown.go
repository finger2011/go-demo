package web

import (
	"context"
	"errors"
	"fmt"
	"os"
	"os/signal"
	"time"
)

//步骤
//1 下线服务
//2 拒绝请求
//3 请求执行完毕
//4 关闭进程
//5 超时直接关闭

//GetSignalError get signal error
// var GetSignalError = errors.New("get signal stop")

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
func ShutDown(ctx context.Context, app *App, ch chan os.Signal) error {
	//监听指定信号
	signal.Notify(ch, app.GetSignal()...)
	select {
	case s := <-ch:
		fmt.Printf("get %s signal, Program Exit...\n", s)
		return errors.New("get exit signal")
	case <-ctx.Done():
		return ctx.Err()
	}
}

//GracefllExit gracefll exit
func GracefllExit(ctx context.Context) error {
	fmt.Println("Start Exit...")
	err := DownService()
	if err != nil {
		fmt.Printf("down service failed:%v", err)
		return err
	}
	err = RejectNewRequest()
	if err != nil {
		fmt.Printf("reject new request failed:%v", err)
		return err
	}
	err = DoneAllRequest()
	if err != nil {
		fmt.Printf("done all request failed:%v", err)
		return err
	}
	fmt.Println("Clean...")
	fmt.Println("End Exit")
	return nil
}

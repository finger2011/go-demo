package main

import (
	"context"
	"os"
	"syscall"
	"time"

	// "errors"
	"fmt"

	"golang.org/x/sync/errgroup"

	// "net/http"
	"finger2011/third-week/web"
)

func home(c *web.Context) {
	fmt.Fprintf(c.W, "这是主页")
}

func user(c *web.Context) {
	fmt.Fprintf(c.W, "这是用户")
}

func createUser(c *web.Context) {
	fmt.Fprintf(c.W, "这是创建用户")
}

func order(c *web.Context) {
	fmt.Fprintf(c.W, "这是订单")
}

//模拟系统调用退出
func mockSyscal(c *web.Context) {
	fmt.Fprintf(c.W, "模拟系统调用退出")
	ch <- syscall.SIGUSR1
}

//用于模拟系统调用，测试
var ch = make(chan os.Signal, 1)

//进程异常退出时，app退出
var closed = make(chan bool, 1)

func main() {
	app := web.CreateApp(context.Background(), "golang")
	group, ctx := errgroup.WithContext(app.GetContext())
	servers := []string{":8080", ":8081"}

	for index, serverName := range servers {
		serverName := serverName
		index := index
		var server = web.CreateSdkHTTPServer(ctx, serverName, closed)
		//demo route
		if index == 0 {
			server.Route("GET", "/", home)
			server.Route("GET", "/user", user)
			server.Route("GET", "/syscall", mockSyscal)
		} else {
			server.Route("PUT", "/user", createUser)
			server.Route("GET", "/order", order)
		}
		group.Go(func() error {
			if err := server.Start(ctx, serverName); err != nil {
				//启动失败，直接退出
				//Start永远返回nil，实际执行不到
				// return err
			}
			fmt.Println("start server done" + serverName)
			return nil
		})
		app.AddServer(server)
	}
	//signal 信号的注册和处理
	group.Go(func() error {
		return web.ShutDown(ctx, app, ch)
	})

	go func() {
		time.Sleep(time.Second * 5)
		app.RandomShutDownServer(ctx)
	}()
	serverClosed(app.GetContext(), app)

	if err := group.Wait(); err != nil {
		fmt.Printf("server error:%v\n", err)
		app.Stop()
	} else {
		fmt.Printf("server nil error:%v\n", err)
	}
}

//server异常退出时，app退出
func serverClosed(ctx context.Context, app *web.App) {
	go func() {
		select {
		case v, ok := <-closed:
			if ok && v {
				break
			}
		}
		fmt.Println("check server closed")
		app.Stop()
	}()
}
